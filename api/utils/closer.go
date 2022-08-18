package utils

import (
	"io"
	"log"
)

func HandleClose(closer io.Closer) {
	if closer != nil {
		err := closer.Close()
		if err != nil {
			log.Printf("unable to close %T: %s", closer, err.Error())
		}
	}
}
