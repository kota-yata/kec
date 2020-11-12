package convert

import (
	"log"
	"os"
)

func ErrorHandling(errMessage error) {
	if errMessage != nil {
		log.Fatal(errMessage)
		os.Exit(1)
	}
}