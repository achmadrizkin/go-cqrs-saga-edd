package usecase

import (
	"fmt"
	"go-cqrs-saga-edd/order-command/domain"
	"log"
)

type orderErrConsumerUseCase struct {
	orderErrConsumerRepo domain.OrderErrConsumerRepo
	orderRepo            domain.OrderRepo
	orderAESRepo         domain.OrderAESRepo
}

// ConsumerErrFromOrderToProduct implements domain.OrderErrConsumerUseCase
func (o *orderErrConsumerUseCase) ConsumerErrFromOrderToProduct(nameQueue string) error {
	msgs, err := o.orderErrConsumerRepo.ConsumerErrFromOrderToProduct(nameQueue)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)

			// this already decrypt and already unmarshal to object
			getMessageOrder, errDecryptedAES := o.orderAESRepo.DecryptOrderAES(d.Body)
			if errDecryptedAES != nil {
				log.Println(errDecryptedAES)
				continue
			}

			tx, errDelete := o.orderRepo.DeleteOrderRepo(getMessageOrder)
			if errDelete != nil {
				log.Println("errDelete: ", errDelete)
			}

			tx.Commit()
		}
	}()

	fmt.Println("[*] Waiting for messages...")

	<-forever // wait until the channel is closed

	return nil
}

func NewOrderErrConsumerUseCase(orderErrConsumerRepo domain.OrderErrConsumerRepo, orderRepo domain.OrderRepo, orderAESRepo domain.OrderAESRepo) domain.OrderErrConsumerUseCase {
	return &orderErrConsumerUseCase{orderErrConsumerRepo, orderRepo, orderAESRepo}
}
