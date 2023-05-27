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

func NewOrderAESRepo() domain.OrderAESRepo {
	return &orderAESRepo{}
}
