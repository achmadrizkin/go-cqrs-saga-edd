package domain

type ProductPublisherRepo interface {
	ProductPublisherFromProductToOrderQuery(encryptedOrder []byte, nameQueue string) error
}
