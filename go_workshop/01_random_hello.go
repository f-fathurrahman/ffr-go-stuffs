package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

var helloList = []string{
	"Hello world",
	"Halo dunia",
	"こんにちは世界",
	"Ave, munde",
	"سلام دنیا‎ ",
	"Καλημέρα κόσμε",
	"Привет, мир",
	"Ciao mondo",
	"مرحبا بالعالم",
	"Bonjour le monde",
}

func main() {
	// Random seed
	rand.Seed(time.Now().UnixNano())
	// Random index
	index := rand.Intn(len(helloList))
	// Call a function
	msg, err := hello(index)
	if err != nil {
		log.Fatal(err)
	}
	// Print the message
	fmt.Println(msg)
}

func hello(index int) (string, error) {
	if index < 0 || index > len(helloList)-1 {
		return "", errors.New("out of range index: " + strconv.Itoa(index))
	}
	return helloList[index], nil
}
