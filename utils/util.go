package utils

import "os"

func Getenv(var_name string, default_value string) string {
	value := os.Getenv(var_name)
	if value == "" {
		value = default_value
	}
	return value
}
