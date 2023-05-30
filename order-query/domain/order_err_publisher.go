package domain

type OrderErrPublisherRepo interface {
	CreateErrOrderQueryPublisherToProductRepo(encryptedOrder []byte, nameQueue string) error
}
