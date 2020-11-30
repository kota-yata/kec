package convert

import (
	"log"
	"os"
)

// ErrorHandling stdout error message and exit with status code 1
func ErrorHandling(errMessage error) {
	if errMessage != nil {
		log.Fatal(errMessage)
		os.Exit(1)
	}
}
