package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var count int

	if len(os.Args) > 1 {
		parsedint, err := strconv.Atoi(os.Args[1])
		if err == nil {
			count = parsedint
		}
	}

	if count == 0 {
		count = 16
	}

	b := make([]byte, count)

	get := func() byte {
		sym := []byte("123456789abcdefghjkmnpqrstwxyzABCDEFGHJKMNPQRSTWXYZ")
		return sym[rand.Intn(len(sym))]
	}

	for i := 0; i < count; i++ {
		b = append(b, get())
	}

	output := string(b)

	fmt.Println(output)
}
