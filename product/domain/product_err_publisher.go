package domain

type ProductErrPubsliherRepo interface {
	ProductErrPublisherFromProductToOrder(encryptedOrder []byte, nameQueue string) error
}
