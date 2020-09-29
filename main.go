package main

import (
	"fmt"

	"github.com/jordiroig-tf/workshop/cypher"
)

func main() {
	message := "HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML"

	decryptedMessage := cypher.Decrypt(message)

	fmt.Println(decryptedMessage)

	encryptedMessage := cypher.Encrypt(decryptedMessage)

	fmt.Println(encryptedMessage)
}
