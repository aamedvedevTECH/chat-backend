package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"github.com/aamedvedevTECH/chat-backend/config"
	"github.com/aamedvedevTECH/chat-backend/internal/app"
)

func main() {
	var (
		configFile   = flag.String("config", "deployments/local.yaml", "specify config file to use")
		globalConfig = config.NewConfig()
		err          error
	)

	flag.Parse()

	log.Printf("config file: %s", *configFile)
	viper.SetConfigFile(*configFile)

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("conf [%s] read err: %s", *configFile, err.Error())
	}

	if err := viper.Unmarshal(globalConfig); err != nil {
		log.Fatalf("unable to unmarshall the config %v", err)
	}

	//TODO: config validation with tests

	var ctx, cancel = context.WithCancel(context.Background())

	appInstance := app.NewApp(globalConfig)

	if err = appInstance.Start(ctx); err != nil {
		log.Fatal(err)
	}

	shutdownApp := make(chan struct{})

	go func(app *app.App) {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
		sig := <-interrupt

		log.Printf("got signal %v\n", sig)

		app.Close(ctx)

		cancel()

		close(shutdownApp)
	}(appInstance)

	<-shutdownApp

	log.Println("app stopped")
}
