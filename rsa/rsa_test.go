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
	msg := 123
	decodedMsg := PowerModulo(msg, pubKey.e, pubKey.n)
	encodedMsg := PowerModulo(decodedMsg, priKey.d, priKey.n)
	if msg != encodedMsg {
		t.Errorf("RSA Key pair is not working. Original msg: %d - encoded msg: %d", msg, encodedMsg)
	}

	decodedMsg = PowerModulo(msg, priKey.d, priKey.n)
	encodedMsg = PowerModulo(decodedMsg, pubKey.e, pubKey.n)
	if msg != encodedMsg {
		t.Errorf("RSA Key pair is not working. Original msg: %d - encoded msg: %d", msg, encodedMsg)
	}
}
