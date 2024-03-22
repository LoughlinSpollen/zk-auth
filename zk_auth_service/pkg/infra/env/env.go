package env

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func WithDefaultString(name, defaultValue string) string {
	val := os.Getenv(name)
	if val == "" {
		val = defaultValue
	}
	return val
}

func WithDefaultInt(name string, defaultValue int) int {
	result := defaultValue
	val := os.Getenv(name)
	if val != "" {
		var err error
		if result, err = strconv.Atoi(val); err != nil {
			log.Printf("could not load environment variable, expected integer, got type: %T\n", val)
			result = defaultValue
		}
	}
	return result
}

func WithDefaultInt64(name string, defaultValue int64) int64 {
	result := defaultValue
	val := os.Getenv(name)
	if val != "" {
		i64, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			log.Printf("could not load environment variable, expected int64, got type: %T\n", val)
			panic(err)
		}
		result = i64
	}
	return result
}
