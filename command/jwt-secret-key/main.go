package main

import (
	"github.com/joho/godotenv"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func String(length int) string {
	return StringWithCharset(length, charset)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}

	key := os.Getenv("JWT_SECRET")
	file, err := os.ReadFile(".env")
	check(err)

	err = os.WriteFile(".env", []byte(strings.Replace(string(file), "\nJWT_SECRET="+key, "", -1)+"\nJWT_SECRET="+String(64)+"\n"), 0644)
	check(err)
}
