package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"strings"
)

type Options struct {
	// EncryptionKey is the encryption key used to encrypt the payloads
	// this key must be 16, 24, 32 characters in length
	// https://en.wikipedia.org/wiki/Advanced_Encryption_Standard
	EncryptionKey        []byte
	EncryptionBackupKeys [][]byte
}

type AESEncryptionServiceV1 struct {
	Cipher           cipher.AEAD
	BackupKeysCipher []cipher.AEAD
}

func newAESEncryptionServiceV1(opts Options) (*AESEncryptionServiceV1, error) {
	// must be 16, 24, 32 byte length
	// this is your encryption key
	// will fail to initialize if length requirements are not met
	cipherBlock, err := aes.NewCipher(opts.EncryptionKey)
	if err != nil {
		// likely invalid key length if errors here
		return nil, err
	}
	gcm, err := cipher.NewGCM(cipherBlock)
	if err != nil {
		return nil, err
	}

	// if opts.EncryptionBackupKeys is empty will just return empty cipher list
	var gcmList = make([]cipher.AEAD, len(opts.EncryptionBackupKeys))
	for i, k := range opts.EncryptionBackupKeys {
		cipherBlockForBackupKeys, err := aes.NewCipher(k[:])
		if err != nil {
			// likely invalid key length if errors here
			return nil, err
		}
		gcmForBackupKeys, err := cipher.NewGCM(cipherBlockForBackupKeys)
		if err != nil {
			return nil, err
		}
		gcmList[i] = gcmForBackupKeys
	}

	return &AESEncryptionServiceV1{
		Cipher:           gcm,
		BackupKeysCipher: gcmList,
	}, nil
}

// Encrypt takes a byte array and returns an encrypted byte array
// as base64 encoded
func (a AESEncryptionServiceV1) Encrypt(unencryptedBytes []byte) ([]byte, error) {
	if len(unencryptedBytes) == 0 { // prevent err on empty byte arrays - "cipher: message authentication failed"
		return []byte(""), nil
	}
	nonce := make([]byte, a.Cipher.NonceSize())
	// populate nonce with random data
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	encryptedBytes := a.Cipher.Seal(nonce, nonce, unencryptedBytes, nil)
	encryptedEncodedData := make([]byte, base64.RawURLEncoding.EncodedLen(len(encryptedBytes)))
	base64.RawURLEncoding.Encode(encryptedEncodedData, encryptedBytes)
	return encryptedEncodedData, nil
}

// Decrypt takes an encrypted base64 byte array then
// returns an unencrypted byte array if same key was used to encrypt it
func (a AESEncryptionServiceV1) Decrypt(encryptedBytes []byte) ([]byte, error) {
	if len(encryptedBytes) == 0 {
		return []byte(""), nil
	}
	decodedEncryptedBytes := make([]byte, base64.RawURLEncoding.DecodedLen(len(encryptedBytes)))
	if _, err := base64.RawURLEncoding.Decode(decodedEncryptedBytes, encryptedBytes); err != nil {
		return nil, fmt.Errorf("unable to decode: %v", err)
	}
	nonceSize := a.Cipher.NonceSize()
	if len(encryptedBytes) < nonceSize {
		return nil, fmt.Errorf("ciphertext too short: %v", len(encryptedBytes))
	}
	decrypted, err := a.Cipher.Open(nil, decodedEncryptedBytes[:nonceSize], decodedEncryptedBytes[nonceSize:], nil)
	if err != nil {
		if strings.Contains(err.Error(), "message authentication failed") {
			log.Println("authentication failed, will proceed to decrypt with the backup keys")
			return a.DecryptWithBackupKeys(encryptedBytes, decodedEncryptedBytes)
		}
	}
	return decrypted, err
}

func (a AESEncryptionServiceV1) DecryptWithBackupKeys(encryptedBytes, decodedEncryptedBytes []byte) ([]byte, error) {
	if len(a.BackupKeysCipher) > 1 {
		for _, cipher := range a.BackupKeysCipher {
			nonceSize := cipher.NonceSize()
			if len(encryptedBytes) < nonceSize {
				return nil, fmt.Errorf("ciphertext too short: %v", len(encryptedBytes))
			}
			decrypted, err := cipher.Open(nil, decodedEncryptedBytes[:nonceSize], decodedEncryptedBytes[nonceSize:], nil)
			if err == nil && len(decrypted) > 0 {
				return decrypted, err
			}
		}
	}
	return nil, fmt.Errorf("unable to decrypt payloads: none of the keys matches to the encrypted payload")
}
