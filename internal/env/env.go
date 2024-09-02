package env

import (
	"os"
	"strconv"
)

// Buscando variaveis no ambiente
// Retorna o fallback caso nao encontre

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	return val
}

// O sistema retorna o valor da key como
// string, sendo necessario converter para int
func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}
