package httpext

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

type Port int

func (p Port) Addr() string {
	return fmt.Sprintf(":%d", p)
}

type JsonError struct {
	Error string `json:"error"`
}

func Dump(request *http.Request, response *http.Response) {
	if request != nil {
		b, err := httputil.DumpRequest(request, true)
		if err != nil {
			log.Println("error on dump request:", err.Error())
		} else {
			fmt.Println(string(b))
		}
	}
	if response != nil {
		b, err := httputil.DumpResponse(response, true)
		if err != nil {
			log.Println("error on dump request:", err.Error())
		} else {
			fmt.Println(string(b))
		}
	}
}
