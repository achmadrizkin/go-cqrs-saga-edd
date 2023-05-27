package domain

type ProductAESRepo interface {
	DecryptProductAES(message []byte) (string, error)
}
