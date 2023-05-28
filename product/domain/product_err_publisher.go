package domain

type ProductErrPubsliher interface {
	ProductErrPublisherFromProductToOrder(encryptedOrder []byte, nameQueue string) error
}
