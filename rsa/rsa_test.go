package rsa

import (
	"testing"
)

func TestPowerModulo(t *testing.T) {
	if PowerModulo(2, 5, 33) != 32 {
		t.Errorf("2^5 mod 33 should be 32 instead if %d", PowerModulo(2, 5, 33))
	}
	if PowerModulo(5, 5, 7) != 3 {
		t.Errorf("5^5 mod 7 should be 3 instead if %d", PowerModulo(5, 5, 7))
	}
	if PowerModulo(1, 0, 7) != 1 {
		t.Errorf("1^0 mod 7 should be 1 instead if %d", PowerModulo(1, 0, 7))
	}
}

func TestRSAPair(t *testing.T) {
	pubKey, priKey := NewRSAKeyPair(83, 97)
	message := 1234
	encodedMsg := decrypt(message, pubKey.e, pubKey.n)
	decodedMsg := encrypt(encodedMsg, priKey.d, priKey.n)
	if message != decodedMsg {
		t.Errorf("RSA Key pair is not working. Original message: %d - decrypted message: %d", message, decodedMsg)
	}

	encodedMsg = decrypt(message, priKey.d, priKey.n)
	decodedMsg = encrypt(encodedMsg, pubKey.e, pubKey.n)
	if message != decodedMsg {
		t.Errorf("RSA Key pair is not working. Original message: %d - decrypted message: %d", message, decodedMsg)
	}
}

func TestRSAEncryptTwoTimes(t *testing.T) {
	pubKeyA, priKeyA := NewRSAKeyPair(83, 97)
	pubKeyB, priKeyB := NewRSAKeyPair(79, 127)
	originalMessageFromA := 1234
	encryptedMessageFromA := decrypt(decrypt(originalMessageFromA, pubKeyB.e, pubKeyB.n), priKeyA.d, priKeyA.n)
	decryptedMessage := encrypt(encrypt(encryptedMessageFromA, pubKeyA.e, pubKeyA.n), priKeyB.d, priKeyB.n)

	if originalMessageFromA != decryptedMessage {
		t.Errorf("RSA Key pairs are not working. Original message: %d - decrypted message: %d", originalMessageFromA, decryptedMessage)
	}
}
