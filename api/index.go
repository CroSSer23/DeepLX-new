package handler

import (
	"fmt"
	"net/http"

	"github.com/OwO-Network/DeepLX/service"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	cfg := service.InitConfig()
	app := service.Router(cfg)
	
	// Route the request through the Gin router
	app.ServeHTTP(w, r)
	
	fmt.Printf("Request processed: %s %s\n", r.Method, r.URL.Path)
} 