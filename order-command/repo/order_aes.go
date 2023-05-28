package repo

import (
	"encoding/json"
	"errors"
	"go-cqrs-saga-edd/order-command/config"
	"go-cqrs-saga-edd/order-command/domain"
	"go-cqrs-saga-edd/order-command/model"
	"go-cqrs-saga-edd/order-command/utils"
	"log"
)

type orderAESRepo struct{}

// EncryptOrderAES implements domain.OrderEncryptRepo
func (*orderAESRepo) EncryptOrderAES(order model.Order) ([]byte, error) {
	// privateToken: 1423456789012345
	privateToken := config.Config("AES_PRIVATE_TOKEN")
	key := []byte(privateToken)

	orderJSON, errMarshal := json.Marshal(order)
	if errMarshal != nil {
		return nil, errors.New("errMarshalOrder: " + errMarshal.Error())
	}

	// Encrypt
	encrypted, errEncryptedAES := utils.EncryptAES(key, string(orderJSON))
	if errEncryptedAES != nil {
		return nil, errors.New("errEncryptedAES: " + errEncryptedAES.Error())
	}

	log.Println("Encrypted Order:", encrypted)
	log.Println([]byte(encrypted))
	return []byte(encrypted), nil
}

// DecryptOrderAES implements domain.OrderAES
func (*orderAESRepo) DecryptOrderAES(message []byte) (model.Order, error) {
	privateToken := config.Config("AES_PRIVATE_TOKEN")
	key := []byte(privateToken)

	// Decrypt message
	decrypted, err := utils.DecryptAES(key, string(message))
	if err != nil {
		return model.Order{}, errors.New("errorDecryptAES: " + err.Error())
	}

	log.Println("Decrypted Order Message:", decrypted)

	var order model.Order
	if err := json.Unmarshal([]byte(decrypted), &order); err != nil {
		return model.Order{}, errors.New("errorUnmarshalingJSON: " + err.Error())
	}

	log.Println("Order Id: " + order.Id)
	log.Println("Order ProductId: " + order.ProductId)
	log.Println("Order Quantity: ", order.Quantity)
	log.Println("Order Address: " + order.Address)
	log.Println("Order ShipMethod : " + order.ShipMethod)
	log.Println("Order Date : " + order.Date.String())

	return order, nil
}

func NewOrderAESRepo() domain.OrderAESRepo {
	return &orderAESRepo{}
}
