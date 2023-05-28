package usecase

import (
	"fmt"
	"go-cqrs-saga-edd/product/domain"
	"log"
)

type productConsumerUseCase struct {
	productConsumerRepo domain.ProductConsumerRepo
	productAESRepo      domain.ProductAESRepo
	productRepo         domain.ProductRepo
	productErrPublisher domain.ProductErrPubsliher
}

// ConsumerProductFromOrderUseCase implements domain.ProductConsumerUseCase
func (p *productConsumerUseCase) ConsumerProductFromOrderUseCase(nameQueueConsumer string, nameQueueErrPublisherToOrder string) error {
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
				log.Println(errDecryptedAES)
				continue
			}

			tx, errUpdateStockProduct := p.productRepo.UpdateStockProductRepo(getMessageOrder.ProductId, int64(getMessageOrder.Quantity), 1)
			if errUpdateStockProduct != nil {
				tx.Rollback() // rollback

				// test rollback -> success
				// already tested
				if errErrPublisher := p.productErrPublisher.ProductErrPublisherFromProductToOrder(d.Body, nameQueueErrPublisherToOrder); errErrPublisher != nil {
					log.Println("errPublisher: ", errErrPublisher)
					continue
				}

				log.Println("errUpdateStockProduct", errUpdateStockProduct, "And rollback to event", nameQueueErrPublisherToOrder)
				continue
			}

			tx.Commit()
		}
	}()

	fmt.Println("[*] Waiting for messages...")

	<-forever // wait until the channel is closed

	return nil
}

func NewProductConsumerUseCase(productConsumerRepo domain.ProductConsumerRepo, productAESRepo domain.ProductAESRepo, productRepo domain.ProductRepo, productErrPublisher domain.ProductErrPubsliher) domain.ProductConsumerUseCase {
	return &productConsumerUseCase{productConsumerRepo, productAESRepo, productRepo, productErrPublisher}
}
