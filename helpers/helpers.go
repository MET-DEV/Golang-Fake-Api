package helpers

import (
	"log"
)

func Met() {

}

func CheckError(err error) {
	if err != nil {
		log.Println(err.Error())
	}
}
