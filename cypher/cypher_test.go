package cypher_test

import (
	"testing"

	"github.com/jordiroig-tf/workshop/cypher"
)

func TestEncrypt(t *testing.T) {
	message := "HELLO THERE"
	expected := "IENVS GKEIG"

	encrypted := cypher.Encrypt(message)

	if encrypted != expected {
		t.Errorf("Error with decrypt: expected %s, got %s", expected, encrypted)
	}
}

func TestDecrypt(t *testing.T) {
	message := "UHKC MF D TVUA."
	expected := "THIS IS A TEST."
	decrypted := cypher.Decrypt(message)

	if decrypted != expected {
		t.Errorf("Error with decrypt: expected %s, got %s", expected, decrypted)
	}
}

func TestEncryptAndDecrypt(t *testing.T) {
	message := "LET'S ENCRYPT AND DECRYPT!"

	result := cypher.Decrypt(cypher.Encrypt(message))

	if result != message {
		t.Errorf("Error with decrypt: expected %s, got %s", message, result)
	}
}
