package repo

import (
	"encoding/json"
	"errors"
	"go-cqrs-saga-edd/order-query/config"
	"go-cqrs-saga-edd/order-query/domain"
	"go-cqrs-saga-edd/order-query/model"
	"go-cqrs-saga-edd/order-query/utils"
	"log"
)

type orderAESRepo struct{}

// DecryptOrderProductAES implements domain.OrderAESRepo
func (o *orderAESRepo) DecryptOrderProductAES(message []byte) (model.OrderProduct, error) {
	privateToken := config.Config("AES_PRIVATE_TOKEN")
	key := []byte(privateToken)

	// Decrypt message
	decrypted, err := utils.DecryptAES(key, string(message))
	if err != nil {
		return model.OrderProduct{}, errors.New("errorDecryptAES: " + err.Error())
	}

	log.Println("Decrypted OrderProduct Message:", decrypted)

	//
	var orderProduct model.OrderProduct
	if err := json.Unmarshal([]byte(decrypted), &orderProduct); err != nil {
		return model.OrderProduct{}, errors.New("errorUnmarshalingJSON: " + err.Error())
	}

	log.Printf("Unmarshaling OrderProduct Message: %+v\n", orderProduct)

	return orderProduct, nil
}

func NewOrderAESRepo() domain.OrderAESRepo {
	return &orderAESRepo{}
}
