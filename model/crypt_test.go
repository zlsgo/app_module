package model

import (
	"testing"

	"github.com/sohaha/zlsgo"
)

func TestHashIDCrypter(t *testing.T) {
	tt := zlsgo.NewTest(t)

	crypter := NewHashIDCrypter("test-salt", 8)

	testIDs := []int64{1, 100, 999999, 1234567890}
	for _, id := range testIDs {
		encrypted, err := crypter.Encrypt(id)
		tt.NoError(err)
		tt.Equal(true, len(encrypted) >= 8)

		decrypted, err := crypter.Decrypt(encrypted)
		tt.NoError(err)
		tt.Equal(id, decrypted)
	}
}

func TestAESCrypter(t *testing.T) {
	tt := zlsgo.NewTest(t)

	key := []byte("0123456789abcdef")
	crypter, err := NewAESCrypter(key)
	tt.NoError(err)

	testIDs := []int64{1, 100, 999999, 1234567890}
	for _, id := range testIDs {
		encrypted, err := crypter.Encrypt(id)
		tt.NoError(err)

		decrypted, err := crypter.Decrypt(encrypted)
		tt.NoError(err)
		tt.Equal(id, decrypted)
	}
}

func TestAESCrypterInvalidKey(t *testing.T) {
	tt := zlsgo.NewTest(t)

	_, err := NewAESCrypter([]byte("short"))
	tt.Equal(ErrInvalidKeyLength, err)

	_, err = NewAESCrypter([]byte("0123456789abcdef"))
	tt.NoError(err)

	_, err = NewAESCrypter([]byte("0123456789abcdef01234567"))
	tt.NoError(err)

	_, err = NewAESCrypter([]byte("0123456789abcdef0123456789abcdef"))
	tt.NoError(err)
}

func TestAESCrypterDifferentOutputs(t *testing.T) {
	tt := zlsgo.NewTest(t)

	key := []byte("0123456789abcdef")
	crypter, _ := NewAESCrypter(key)

	id := int64(12345)
	encrypted1, _ := crypter.Encrypt(id)
	encrypted2, _ := crypter.Encrypt(id)

	tt.Equal(true, encrypted1 != encrypted2)

	decrypted1, _ := crypter.Decrypt(encrypted1)
	decrypted2, _ := crypter.Decrypt(encrypted2)
	tt.Equal(id, decrypted1)
	tt.Equal(id, decrypted2)
}

func TestAESCrypterInvalidDecrypt(t *testing.T) {
	tt := zlsgo.NewTest(t)

	key := []byte("0123456789abcdef")
	crypter, _ := NewAESCrypter(key)

	_, err := crypter.Decrypt("invalid-base64!@#")
	tt.Equal(ErrInvalidCryptedID, err)

	_, err = crypter.Decrypt("c2hvcnQ")
	tt.Equal(ErrInvalidCryptedID, err)
}

func BenchmarkHashIDCrypter(b *testing.B) {
	crypter := NewHashIDCrypter("benchmark-salt", 8)
	id := int64(1234567890)

	b.Run("Encrypt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			crypter.Encrypt(id)
		}
	})

	encrypted, _ := crypter.Encrypt(id)
	b.Run("Decrypt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			crypter.Decrypt(encrypted)
		}
	})
}

func BenchmarkAESCrypter(b *testing.B) {
	key := []byte("0123456789abcdef")
	crypter, _ := NewAESCrypter(key)
	id := int64(1234567890)

	b.Run("Encrypt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			crypter.Encrypt(id)
		}
	})

	encrypted, _ := crypter.Encrypt(id)
	b.Run("Decrypt", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			crypter.Decrypt(encrypted)
		}
	})
}
