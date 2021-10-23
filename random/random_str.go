package random

import (
	"crypto/rand"
	"math/big"
)

func RandomString(words []string) string {
	max := len(words)
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(max)))
	return get(r.Int64(), words)
}

func get(index int64, words []string) string {
	return words[index]
}

func Names_db() string {
	names := []string{"Igor", "Alex", "Tom", "Andrey", "Ilya", "Ivan", "Oleg", "Pavel", "Gleb", "Denis"}
	return RandomString(names)
}

func LastNames_db() string {
	family := []string{"Ivanov", "Petrov", "Sviridov", "Avdeev", "Bauman", "Davidof", "Demidov", "Solodkevich", "Kostin", "Romanov"}
	return RandomString(family)
}

func Cards_db() string {
	cards := []string{"VISA", "MASTERCARD"}
	return RandomString(cards)
}
