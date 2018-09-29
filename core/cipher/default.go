package cipher

type defaultCipher struct {
	crypt *Crypt
	skip  int
	cur   int64
	// Original iv
	iv [4]byte
}

func NewDefaultCipher(version uint16, iv [4]byte, skip int) ICipher {
	crypt := NewCrypt(iv, version)
	return &defaultCipher{
		crypt: &crypt,
		skip:  skip,
		iv:    iv,
	}
}

func (c *defaultCipher) Encrypt(data []byte) {
	if c.cur < int64(c.skip) {
		return
	}
	c.crypt.Encrypt(data)
	c.crypt.Shuffle()
	c.cur++
}

func (c *defaultCipher) DecryptHeader(data []byte) int {
	return GetPacketLength(data)
}

func (c *defaultCipher) DecryptBody(data []byte) {
	if c.cur < int64(c.skip) {
		return
	}
	c.crypt.Decrypt(data)
	c.crypt.Shuffle()
	c.cur++
}

func (c *defaultCipher) GetKey() []byte {
	return c.iv[:]
}
