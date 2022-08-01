package kubectl

import (
	"fmt"
	"os"
)

func getenv(env string) (string, error) {
	value := os.Getenv(env)
	if value == "" {
		return "", fmt.Errorf("%s is empty", env)
	}
	return value, nil
}
