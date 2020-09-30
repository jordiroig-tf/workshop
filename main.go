package main

import (
	"fmt"

	"github.com/jordiroig-tf/workshop/blackjack"
)

func main() {
	player1 := blackjack.Player("Jordi")
	players := []blackjack.Player{player1}

	game, err := blackjack.New(players)

	fmt.Println(game, err)

	//message := "HO KC EJHSFOL, IGH TXUCPZ FWX XB SLRA DQML"
	//
	//decryptedMessage := cypher.Decrypt(message)
	//
	//fmt.Println(decryptedMessage)
	//
	//encryptedMessage := cypher.Encrypt(decryptedMessage)
	//
	//fmt.Println(encryptedMessage)
}
