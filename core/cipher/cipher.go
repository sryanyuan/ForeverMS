package cipher

type ICipher interface {
	Encrypt([]byte)
	// DecryptHeader should thread-safe
	DecryptHeader([]byte) int
	DecryptBody([]byte)
	GetKey() []byte
}
