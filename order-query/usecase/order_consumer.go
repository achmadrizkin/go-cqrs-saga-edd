package usecase

import (
	"context"
	"errors"
	"fmt"
	"go-cqrs-saga-edd/order-query/domain"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

type orderQueryConsumerUseCase struct {
	orderQueryConsumerRepo domain.OrderQueryConsumerRepo
	orderAESRepo           domain.OrderAESRepo
	orderCommandRepo       domain.OrderCommandRepo
	orderErrPublisherRepo  domain.OrderErrPublisherRepo
	client                 *mongo.Client
}

// ConsumerOrderQueryConsumerRepo implements domain.OrderQueryConsumerUseCase
func (o *orderQueryConsumerUseCase) ConsumerOrderQueryConsumerRepo(ctx context.Context, nameQueueConsumer string, sendErrorPublisherToProduct string) error {
	msgs, err := o.orderQueryConsumerRepo.ConsumerOrderQuerConsumerRepo(nameQueueConsumer)
	if err != nil {
		return err
	}

	// Set the read and write concerns for the transaction
	txOptions := options.Transaction().
		SetReadConcern(readconcern.Majority()).
		SetWriteConcern(writeconcern.New())

	// Start a session
	session, err := o.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	// Start a transaction
	err = session.StartTransaction(txOptions)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received message: %s\n", d.Body)

			// Start Session
			errSession := mongo.WithSession(ctx, session, func(sc mongo.SessionContext) error {
				// This already decrypts and unmarshals to an object
				getMessageOrderProduct, errDecryptedAES := o.orderAESRepo.DecryptOrderProductAES(d.Body)
				if errDecryptedAES != nil {
					o.orderErrPublisherRepo.CreateErrOrderQueryPublisherToProductRepo(d.Body, sendErrorPublisherToProduct)
					return errors.New("errDecryptedAES: " + errDecryptedAES.Error())
				}

				errCreatedProduct := o.orderCommandRepo.CreateOrderProduct(sc, getMessageOrderProduct)
				if errCreatedProduct != nil {
					o.orderErrPublisherRepo.CreateErrOrderQueryPublisherToProductRepo(d.Body, sendErrorPublisherToProduct)
					return errors.New("errCreatedProduct: " + errCreatedProduct.Error())
				}

				// check rollback
				// for rollback purpose
				// i := 1
				// ix := 1
				// if i == ix {
				// 	o.orderErrPublisherRepo.CreateErrOrderQueryPublisherToProductRepo(d.Body, sendErrorPublisherToProduct)
				// 	return errors.New("testRollback: ")
				// }

				if errCommitTransaction := session.CommitTransaction(sc); errCommitTransaction != nil {
					o.orderErrPublisherRepo.CreateErrOrderQueryPublisherToProductRepo(d.Body, sendErrorPublisherToProduct)
					return errors.New("errCommitTransaction: " + errCommitTransaction.Error())
				}
				return nil
			})

			if errSession != nil {
				continue
			}
		}
	}()

	fmt.Println("[*] Waiting for messages...")

	<-forever // Wait until the channel is closed

	return nil
}

func NewOrderQueryConsumerUseCase(orderQueryConsumerRepo domain.OrderQueryConsumerRepo, orderAESRepo domain.OrderAESRepo, orderCommandRepo domain.OrderCommandRepo, client *mongo.Client, orderErrPublisherRepo domain.OrderErrPublisherRepo) domain.OrderQueryConsumerUseCase {
	return &orderQueryConsumerUseCase{
		orderQueryConsumerRepo,
		orderAESRepo,
		orderCommandRepo,
		orderErrPublisherRepo,
		client,
	}
}
