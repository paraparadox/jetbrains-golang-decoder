package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(strings.Split(s, "\n")[0], " ")

	g, err := strconv.ParseInt(input[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	p, err := strconv.ParseInt(input[6], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	b := rand.Int63n(p-1) + 1
	var B, i int64
	B = 1
	for i = 0; i < b; i++ {
		B = (B * g) % p
	}

	fmt.Println("OK")

	s, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	input = strings.Split(strings.Split(s, "\n")[0], " ")
	A, err := strconv.ParseInt(input[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	var S int64
	S = 1
	for i = 0; i < b; i++ {
		S = (S * A) % p
	}
	S %= 26
	fmt.Println("B is", B)

	// [65; 90] A-Z
	// [97; 122] a-z

	message := "Will you marry me?"
	encryptedMessage := make([]byte, len(message))
	for i = 0; i < int64(len(message)); i++ {
		if int(message[i]) > 64 && int(message[i]) < 91 {
			encryptedMessage[i] = byte((int64(message[i]) + S) % 91 + 65)
		} else if int(message[i]) > 96 && int(message[i]) < 123 {
			encryptedMessage[i] = byte((int64(message[i]) + S) % 123 + 97)
		} else {
			encryptedMessage[i] = message[i]
		}
	}
	fmt.Println(encryptedMessage, string(encryptedMessage[:]))
	message, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	message = strings.Split(message, "\n")[0]
	decryptedMessage := make([]byte, len(message))
	for i = 0; i < int64(len(message)); i++ {
		if int(message[i]) > 64 && int(message[i]) < 91 {
			encryptedMessage[i] = byte(int64(message[i]) - 65 + 91)
		} else if int(message[i]) > 96 && int(message[i]) < 123  {
			encryptedMessage[i] = byte(int64(message[i]) - 97 + 123)
		} else {
			encryptedMessage[i] = message[i]
		}
	}
	fmt.Println(string(decryptedMessage[:]))
}
