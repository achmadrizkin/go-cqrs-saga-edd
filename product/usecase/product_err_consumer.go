package usecase

import (
	"fmt"
	"go-cqrs-saga-edd/product/domain"
	"log"
)

type productErrConsumerUseCase struct {
	productErrConsumerRepo  domain.ProductErrConsumerRepo
	productAESRepo          domain.ProductAESRepo
	productRepo             domain.ProductRepo
	productErrPublisherRepo domain.ProductErrPubsliherRepo
}

// ErrConsumerProductFromOrderUseCase implements domain.ProductErrConsumerUseCase
func (p *productErrConsumerUseCase) ErrConsumerProductFromOrderUseCase(nameQueueConsumer string, nameQueueErrPublisher string) error {
	msgs, err := p.productErrConsumerRepo.ErrConsumerProductFromOrderRepo(nameQueueConsumer)
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

			tx, errUpdateStockProduct := p.productRepo.UpdateStockProductRepo(getMessageOrder.ProductId, int64(getMessageOrder.Quantity), 0)
			if errUpdateStockProduct != nil {
				log.Println(errUpdateStockProduct)
				continue
			}

			if err := p.productErrPublisherRepo.ProductErrPublisherFromProductToOrder(d.Body, nameQueueErrPublisher); err != nil {
				log.Println(err)
				continue
			}

			tx.Commit()
		}
	}()

	fmt.Println("[*] Waiting for messages...")

	<-forever // wait until the channel is closed

	return nil
}

func NewProductErrConsumerUseCase(
	productErrConsumerRepo domain.ProductErrConsumerRepo,
	productAESRepo domain.ProductAESRepo,
	productRepo domain.ProductRepo,
	productErrPublisherRepo domain.ProductErrPubsliherRepo) domain.ProductErrConsumerUseCase {
	return &productErrConsumerUseCase{
		productErrConsumerRepo,
		productAESRepo,
		productRepo,
		productErrPublisherRepo,
	}
}
