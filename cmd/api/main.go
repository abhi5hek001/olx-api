package main

import (
	"fmt"
	"log"
	"log/slog"
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

	logHandler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level: slog.LevelDebug,
	})
	logger := slog.New(logHandler)
	slog.SetDefault(logger)

	fmt.Println("database connected")
	fmt.Println("starting olx server")

	lh := handler.NewListingHandler(db, logger)
	
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", handler.Health)
	mux.HandleFunc("GET /listings", lh.List)
	mux.HandleFunc("DELETE /listings/{id}", lh.Delete)

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
 