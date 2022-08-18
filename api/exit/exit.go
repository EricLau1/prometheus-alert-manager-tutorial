package exit

import (
	"log"
	"os"
	"os/signal"
)

const (
	SUCCESS = 0
	FAILURE = 1
)

func Graceful(onStop ...func()) {
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		for {
			select {
			case _ = <-interrupt:
				for _, stop := range onStop {
					stop()
				}
				log.Println("Stopped.")
				os.Exit(SUCCESS)
			}
		}
	}()
}
