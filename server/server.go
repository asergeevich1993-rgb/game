package server

import (
	"context"
	"go-from-zero/concurrency3/handlers"
	"net/http"
)

type HTTPServer struct {
	handler *handlers.HttpHandlers
	svr     *http.Server
}

func NewServer(hh *handlers.HttpHandlers) *HTTPServer {
	return &HTTPServer{
		handler: hh,
	}
}
func (hs *HTTPServer) StartServer() {
	router := http.NewServeMux()

	router.HandleFunc("POST /miners", hs.handler.HandleCreateMiner)
	router.HandleFunc("POST /equipment/{name}", hs.handler.HandlerBuyEquipment)
	router.HandleFunc("GET /equipment", hs.handler.HandleShowEquipment)
	router.HandleFunc("GET /miners", hs.handler.HandleShowMiners)
	router.HandleFunc("GET /enterprise", hs.handler.HandleShowAllInfo)
	router.HandleFunc("POST /finish", hs.handler.HandleFinish)

	hs.svr = &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	hs.svr.ListenAndServe()
}

func (hs *HTTPServer) Shutdown(ctx context.Context) error {
	return hs.svr.Shutdown(ctx)
}
