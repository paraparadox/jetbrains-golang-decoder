package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

	b := rand.Int63n(p)

	B := int64(math.Pow(float64(g), float64(b))) % p
	fmt.Println(B)
}
