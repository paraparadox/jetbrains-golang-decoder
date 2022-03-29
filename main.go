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

	// Reading g and p
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

	// Generating b and computing B
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

	// Reading A
	input = strings.Split(strings.Split(s, "\n")[0], " ")
	A, err := strconv.ParseInt(input[2], 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Computing the secret
	var S int64
	S = 1
	for i = 0; i < b; i++ {
		S = (S * A) % p
	}
	S %= 26
	fmt.Println("B is", B)

	// [65; 90] A-Z
	// [97; 122] a-z

	// Encoding and printing the question
	message := "Will you marry me?"
	encryptedMessage := make([]byte, len(message))
	for i = 0; i < int64(len(message)); i++ {
		if int(message[i]) >= 65 && int(message[i]) <= 90 {
			encryptedMessage[i] = byte(int64(message[i]) + S)
			if encryptedMessage[i] > 90 {
				encryptedMessage[i] -= 26
			}
		} else if int(message[i]) >= 97 && int(message[i]) <= 122 {
			encryptedMessage[i] = byte(int64(message[i]) + S)
			if encryptedMessage[i] > 122 {
				encryptedMessage[i] -= 26
			}
		} else {
			encryptedMessage[i] = message[i]
		}
	}
	fmt.Println(string(encryptedMessage[:]))

	// Reading and decoding the answer
	message, err = reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	message = strings.Split(message, "\n")[0]
	decryptedMessage := make([]byte, len(message))
	for i = 0; i < int64(len(message)); i++ {
		if int(message[i]) >= 65 && int(message[i]) <= 90 {
			decryptedMessage[i] = byte(int64(message[i]) - S)
			if decryptedMessage[i] < 65 {
				decryptedMessage[i] += 26
			}
		} else if int(message[i]) >= 97 && int(message[i]) <= 122 {
			decryptedMessage[i] = byte(int64(message[i]) - S)
			if decryptedMessage[i] < 97 {
				decryptedMessage[i] += 26
			}
		} else {
			decryptedMessage[i] = message[i]
		}
	}
	answer := string(decryptedMessage[:])

	// Checking the decoded answer
	if answer == "Yeah, okay!" {
		message = "Great!"
	} else if answer == "Let's be friends." {
		message = "What a pity!"
	} else {
		message = ""
	}

	if message != "" {
		// Encoding and printing response
		encryptedMessage = nil
		encryptedMessage = make([]byte, len(message))
		for i = 0; i < int64(len(message)); i++ {
			if int(message[i]) >= 65 && int(message[i]) <= 90 {
				encryptedMessage[i] = byte(int64(message[i]) + S)
				if encryptedMessage[i] > 90 {
					encryptedMessage[i] -= 26
				}
			} else if int(message[i]) >= 97 && int(message[i]) <= 122 {
				encryptedMessage[i] = byte(int64(message[i]) + S)
				if encryptedMessage[i] > 122 {
					encryptedMessage[i] -= 26
				}
			} else {
				encryptedMessage[i] = message[i]
			}
		}
		fmt.Println(string(encryptedMessage[:]))
	}
}
