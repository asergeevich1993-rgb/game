package main

import (
	"context"
	"go-from-zero/concurrency3/enterprise"
	"go-from-zero/concurrency3/handlers"
	"go-from-zero/concurrency3/server"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ent := enterprise.NewEnterprice()
	ent.Start(ctx)
	handler := handlers.NewHandler(ent, cancel, ctx)
	sserver := server.NewServer(handler)
	go func() {
		<-ctx.Done()
		sserver.Shutdown(context.Background())
	}()
	sserver.StartServer()

}
