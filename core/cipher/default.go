package cipher

type defaultCipher struct {
	crypt  *Crypt
	skip   int
	cur    int64
	shanda bool
	// Original iv
	iv [4]byte
}

func NewDefaultCipher(version uint16, iv [4]byte, shanda bool, skip int) ICipher {
	crypt := NewCrypt(iv, version)
	return &defaultCipher{
		crypt:  &crypt,
		skip:   skip,
		iv:     iv,
		shanda: shanda,
	}
}

func (c *defaultCipher) Encrypt(data []byte) {
	if c.cur < int64(c.skip) {
		return
	}
	if c.shanda {
		c.crypt.Encrypt(data)
	} else {
		c.crypt.EncryptNoShanda(data)
	}

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
	if c.shanda {
		c.crypt.Decrypt(data)
	} else {
		c.crypt.DecryptNoShanda(data)
	}

	c.crypt.Shuffle()
	c.cur++
}

func (c *defaultCipher) GetKey() []byte {
	return c.iv[:]
}
