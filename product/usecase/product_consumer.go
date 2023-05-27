package usecase

import (
	"fmt"
	"go-cqrs-saga-edd/product/domain"
	"log"
)

type productConsumerUseCase struct {
	productConsumerRepo domain.ProductConsumerRepo
	productAESRepo      domain.ProductAESRepo
}

// ConsumerProductFromOrderUseCase implements domain.ProductConsumerUseCase
func (p *productConsumerUseCase) ConsumerProductFromOrderUseCase(nameQueue string) error {
	msgs, err := p.productConsumerRepo.ConsumerProductFromOrderRepo(nameQueue)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)

			_, errDecryptedAES := p.productAESRepo.DecryptProductAES(d.Body)
			if errDecryptedAES != nil {
				log.Println(errDecryptedAES)
				continue
			}

		}
	}()

	fmt.Println("[*] Waiting for messages...")

	<-forever // wait until the channel is closed

	return nil
}

func NewProductConsumerUseCase(productConsumerRepo domain.ProductConsumerRepo, productAESRepo domain.ProductAESRepo) domain.ProductConsumerUseCase {
	return &productConsumerUseCase{productConsumerRepo, productAESRepo}
}
