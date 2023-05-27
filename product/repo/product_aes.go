package repo

import (
	"errors"
	"go-cqrs-saga-edd/order-command/utils"
	"go-cqrs-saga-edd/product/config"
	"go-cqrs-saga-edd/product/domain"
	"log"
)

type productAESRepo struct{}

// DecryptProductAES implements domain.ProductAESRepo
func (*productAESRepo) DecryptProductAES(message []byte) (string, error) {
	privateToken := config.Config("AES_PRIVATE_TOKEN")
	key := []byte(privateToken)

	// Decrypt message
	decrypted, err := utils.DecryptAES(key, string(message))
	if err != nil {
		return "", errors.New("errorDecryptAES: " + err.Error())
	}

	log.Println("Decrypted Product Message: " + decrypted)
	return decrypted, nil
}

func NewProductAESRepo() domain.ProductAESRepo {
	return &productAESRepo{}
}
