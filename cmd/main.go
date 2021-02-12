package main

import (
	"context"
	"log"
	"net/http"
	"oslo-bikes/pkg"
)

func Main() error {
	config, err := pkg.NewConfig()
	if err != nil {
		return err
	}

	ctx := context.Background()
	http := &http.Client{Transport: nil}
	client, err := pkg.NewOsloBikesClient(ctx, http, config)
	if err != nil {
		return err
	}
	server := pkg.NewServer(ctx, 8080, client)
	err = server.Start()
	if err != nil {
		return err
	}
	return nil

}
func main() {
	if err := Main(); err != nil {
		log.Fatalf("%v", err)
	}
}
