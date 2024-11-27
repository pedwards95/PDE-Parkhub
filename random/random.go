package random

import (
	"math/rand"
	"time"

	"github.com/pedwards95/PDE-Parkhub/blockchain"
)

// ...
const (
	MIN_LEN = 12
	MAX_LEN = 16
)

// GenerateRandomBlock makes a new random block object
func GenerateRandomBlock() *blockchain.Block {
	block := &blockchain.Block{
		Str:     GenerateRandomString(),
		Time:    time.Now(),
		Company: GenerateRandomName(),
	}
	return block
}

// GenerateRandomString create a random alphanumeric string between min and max length
func GenerateRandomString() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	min := MIN_LEN
	max := MAX_LEN
	length := rand.Intn(max-min) + min

	var result []byte
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(charset))
		result = append(result, charset[randomIndex])
	}

	return string(result)
}

// GenerateRandomName picks a name randomly from the name array
func GenerateRandomName() string {
	names := []string{"Alice", "Bob", "Charlie", "David", "Emily", "ParkHub", "Wayleadr", "Parkable", "Tidaro", "Whatspot"}
	randomIndex := rand.Intn(len(names))
	return names[randomIndex]
}
