// Package env reads environment variables.
package env

import (
	"log"
	"os"
	"strconv"
)

// String extracts string from env var.
// It returns the provided defaultValue if the env var is empty.
// The string returned is also recorded in logs.
func String(name string, defaultValue string) string {
	str := os.Getenv(name)
	if str != "" {
		log.Printf("%s=[%s] using %s=%s default=%s", name, str, name, str, defaultValue)
		return str
	}
	log.Printf("%s=[%s] using %s=%s default=%s", name, str, name, defaultValue, defaultValue)
	return defaultValue
}

// Bool extracts bool from env var.
// It returns the provided defaultValue if the env var is empty.
// The bool returned is also recorded in logs.
func Bool(name string, defaultValue bool) bool {
	str := os.Getenv(name)
	if str != "" {
		log.Printf("%s=[%s] using %s=%s default=%t", name, str, name, str, defaultValue)
		v, errConv := strconv.ParseBool(str)
		if errConv == nil {
			return v
		}
		log.Printf("%s=[%s] error: %v", name, str, errConv)
	}
	log.Printf("%s=[%s] using %s=%t default=%t", name, str, name, defaultValue, defaultValue)
	return defaultValue
}
