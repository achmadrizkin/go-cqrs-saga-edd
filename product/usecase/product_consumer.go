package usecase

import (
	"fmt"
	"go-cqrs-saga-edd/product/domain"
	"go-cqrs-saga-edd/product/utils"
	"log"

	"gorm.io/gorm"
)

type productConsumerUseCase struct {
	productConsumerRepo     domain.ProductConsumerRepo
	productAESRepo          domain.ProductAESRepo
	productRepo             domain.ProductRepo
	productErrPublisherRepo domain.ProductErrPubsliherRepo
	productPublisherRepo    domain.ProductPublisherRepo
}

// ConsumerProductFromOrderUseCase implements domain.ProductConsumerUseCase
func (p *productConsumerUseCase) ConsumerProductFromOrderUseCase(nameQueueConsumer string, nameQueueErrPublisherToOrder string, nameQueuePublisherToOrderQuery string) error {
	msgs, err := p.productConsumerRepo.ConsumerProductFromOrderRepo(nameQueueConsumer)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)

			// this already decrypt and already unmarshal to object
			getMessageOrder, errDecryptedAES := p.productAESRepo.DecryptProductAES(d.Body)
			if errDecryptedAES != nil {
				p.HandleRollbackAndErrorInvalidEncryption(d.Body, nameQueueErrPublisherToOrder, errDecryptedAES)
				continue
			}

			tx, errUpdateStockProduct := p.productRepo.UpdateStockProductRepo(getMessageOrder.ProductId, int64(getMessageOrder.Quantity), 1)
			if errUpdateStockProduct != nil {
				p.HandleRollbackAndError(tx, d.Body, nameQueueErrPublisherToOrder, errUpdateStockProduct)
				continue
			}

			// get product
			product, errGetProduct := p.productRepo.GetProductByIdRepo(getMessageOrder.ProductId)
			if errGetProduct != nil {
				p.HandleRollbackAndError(tx, d.Body, nameQueueErrPublisherToOrder, errGetProduct)
				continue
			}

			// convert into orderProduct for saving in NoSQL
			orderProduct := utils.ConverterOrderAndProductToOrderProduct(getMessageOrder, product)
			log.Printf("ORDER PRODUCT: %+v\n", orderProduct)

			// encrypt data first
			encryptedOrderProductMessage, errEncryptOrderProduct := p.productAESRepo.EncryptOrderProductAES(orderProduct)
			if errEncryptOrderProduct != nil {
				p.HandleRollbackAndError(tx, d.Body, nameQueueErrPublisherToOrder, errEncryptOrderProduct)
				continue
			}

			// publish event from product to order query (no sql)
			if errPublisherProductToOrderQuery := p.productPublisherRepo.ProductPublisherFromProductToOrderQuery(encryptedOrderProductMessage, nameQueuePublisherToOrderQuery); errPublisherProductToOrderQuery != nil {
				p.HandleRollbackAndError(tx, d.Body, nameQueueErrPublisherToOrder, errPublisherProductToOrderQuery)
				continue
			}

			// commit all transaction
			tx.Commit()
		}
	}()

	fmt.Println("[*] Waiting for messages...")

	<-forever // wait until the channel is closed

	return nil
}

func (p *productConsumerUseCase) HandleRollbackAndError(tx *gorm.DB, messageBody []byte, nameQueueErrPublisherToOrder string, err error) {
	tx.Rollback()

	if errErrPublisher := p.productErrPublisherRepo.ProductErrPublisherFromProductToOrder(messageBody, nameQueueErrPublisherToOrder); errErrPublisher != nil {
		log.Println("errPublisher: ", errErrPublisher)
	}
	log.Printf("Error: %v. Rolled back to event: %s\n", err, nameQueueErrPublisherToOrder)
}

func (p *productConsumerUseCase) HandleRollbackAndErrorInvalidEncryption(messageBody []byte, nameQueueErrPublisherToOrder string, err error) {
	if errErrPublisher := p.productErrPublisherRepo.ProductErrPublisherFromProductToOrder(messageBody, nameQueueErrPublisherToOrder); errErrPublisher != nil {
		log.Println("errPublisher: ", errErrPublisher)
	}
	log.Printf("Error: %v. Rolled back to event: %s\n", err, nameQueueErrPublisherToOrder)
}

func NewProductConsumerUseCase(productConsumerRepo domain.ProductConsumerRepo,
	productAESRepo domain.ProductAESRepo,
	productRepo domain.ProductRepo,
	productErrPublisherRepo domain.ProductErrPubsliherRepo,
	productPublisherRepo domain.ProductPublisherRepo) domain.ProductConsumerUseCase {
	return &productConsumerUseCase{productConsumerRepo,
		productAESRepo,
		productRepo,
		productErrPublisherRepo,
		productPublisherRepo,
	}
}
