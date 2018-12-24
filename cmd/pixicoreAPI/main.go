package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jalilbengoufa/pixicoreAPI/pkg/api"
	"github.com/jalilbengoufa/pixicoreAPI/pkg/config"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

func main() {

	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	myConfigFile, err := config.InitConfig()
	if err != nil {
		log.Fatalln(err)
	}
	controller := api.InitController(myConfigFile)

	go func() {
		sig := <-sigs
		log.Println(sig)
		done <- true
	}()

	go func() {

		gin.SetMode(gin.ReleaseMode)

		r := api.GetRouter(controller)
		r.Run(":3000")
	}()

	// The program will wait here until it gets the
	// expected signal (as indicated by the goroutine
	// above sending a value on `done`) and then exit.
	<-done
	myConfigFile.WriteYamlConfig()
	log.Println("exiting")
}
