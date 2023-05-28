package repo

import (
	"encoding/json"
	"errors"
	"go-cqrs-saga-edd/order-command/utils"
	"go-cqrs-saga-edd/product/config"
	"go-cqrs-saga-edd/product/domain"
	"go-cqrs-saga-edd/product/model"
	"log"
)

type productAESRepo struct{}

// EncryptOrderAES implements domain.OrderEncryptRepo
func (*productAESRepo) EncryptOrderProductAES(orderProduct model.OrderProduct) ([]byte, error) {
	// privateToken: 1423456789012345
	privateToken := config.Config("AES_PRIVATE_TOKEN")
	key := []byte(privateToken)

	orderProductJSON, errMarshal := json.Marshal(orderProduct)
	if errMarshal != nil {
		return nil, errors.New("errMarshalOrder: " + errMarshal.Error())
	}

	// Encrypt
	encrypted, errEncryptedAES := utils.EncryptAES(key, string(orderProductJSON))
	if errEncryptedAES != nil {
		return nil, errors.New("errEncryptedAES: " + errEncryptedAES.Error())
	}

	log.Println("Encrypted OrderProduct:", encrypted)
	log.Println([]byte(encrypted))
	return []byte(encrypted), nil
}

// DecryptProductAES implements domain.ProductAESRepo
func (*productAESRepo) DecryptProductAES(message []byte) (model.Order, error) {
	privateToken := config.Config("AES_PRIVATE_TOKEN")
	key := []byte(privateToken)

	// Decrypt message
	decrypted, err := utils.DecryptAES(key, string(message))
	if err != nil {
		return model.Order{}, errors.New("errorDecryptAES: " + err.Error())
	}

	log.Println("Decrypted Product Message:", decrypted)

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

func NewProductAESRepo() domain.ProductAESRepo {
	return &productAESRepo{}
}
