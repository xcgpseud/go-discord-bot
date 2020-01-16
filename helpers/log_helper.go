package helpers

import "fmt"

func LogError(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
