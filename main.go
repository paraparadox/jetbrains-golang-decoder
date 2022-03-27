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
	fmt.Println("B is", B)
}
