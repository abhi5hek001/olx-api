package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/abhi5hek001/olx-api/internal/config"
	"github.com/abhi5hek001/olx-api/internal/db"
	"github.com/abhi5hek001/olx-api/internal/handler"
)

func main() {
	cfg := config.MustLoad()
	db, err := db.Connect(cfg.DATABASE_URL)
	if err != nil {
		log.Fatalf("main.db.connect: %v", err)
	}

	fmt.Println("database connected")
	fmt.Println("starting olx server")
	
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /listings", handler.List(db))

	srv := http.Server{
		Addr: ":" + cfg.PORT,
		Handler: mux,
		ReadTimeout: time.Second*10,
		WriteTimeout: time.Second*30,
		IdleTimeout: time.Second*60,
	}

	fmt.Printf(`Server started on: https://localhost:%v`,os.Getenv("PORT"))
	
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
 