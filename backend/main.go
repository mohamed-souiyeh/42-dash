package main

import (
	"backend/handlers"
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	dashDB "backend/db"
	db "backend/db/repository"

	_ "github.com/mattn/go-sqlite3"

	"github.com/charmbracelet/log"
	"github.com/joho/godotenv"
)

// /data/provider-alpha.json,/data/provider-beta.json,/data/provider-gamma.json

const (
	host = "0.0.0.0"
	port = ":3000"
)

type App struct {
	db            *sql.DB
	Server        *http.Server
	ingestHandler *handlers.IngestHandler
}

func NewApp() *App {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, relying on environment variables", "error", err)
	}

	app := new(App)

	app.dbSetup()

	app.serverSetup()

	return app
}

func (a *App) serverSetup() {
	a.ingestHandler = &handlers.IngestHandler{
		IngestRepo: db.NewSQLiteIngestionRepository(a.db),
		GameRepo:   db.NewSQLiteGameRepository(a.db),
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handlers.HealthzHandler)
	mux.HandleFunc("/ingest", a.ingestHandler.HandleTriggerIngest)

	// --- 3. The Server Object (Configured Once) ---

	server := &http.Server{
		Addr:         port,
		Handler:      mux,              // Attach your router here
		ReadTimeout:  10 * time.Second, // Drop connections that take too long to send requests
		WriteTimeout: 10 * time.Second, // Drop connections if our response takes too long
	}

	a.Server = server
}

func (a *App) serverStart() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		log.Info("Server starting on port 3000")
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
	}()

	go func() {
		// This serves your local "mock_data" folder on a different port
		mockMux := http.NewServeMux()
		mockMux.Handle("/", http.FileServer(http.Dir("./data")))

		// Listen on port 8081 so it doesn't clash with your main API
		http.ListenAndServe(":4444", mockMux)
	}()

	<-stopChan
	log.Info("Shutting down server gracefully...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := a.Server.Shutdown(ctx); err != nil {
		log.Fatalf("Graceful shutdown failed: %v\n", err)
	}

	log.Info("Server stopped cleanly.")
}

func (a *App) dbSetup() {
	// 1. Get the local file path (or default to app.db)
	dbFilePath := os.Getenv("DB_FILE_PATH")
	if dbFilePath == "" {
		dbFilePath = "./app.db"
		log.Warn("DB_FILE_PATH not set, defaulting to './app.db'")
	}

	dsn := fmt.Sprintf("file:%s?_journal_mode=WAL&_busy_timeout=5000&_fk=1", dbFilePath)

	log.Info("Connecting to SQLite database...")

	database, err := sql.Open("sqlite3", dsn)
	if err != nil {
		log.Fatalf("Failed to open SQLite database: %v", err)
	}

	database.SetMaxOpenConns(50)
	database.SetMaxIdleConns(15)
	database.SetConnMaxLifetime(30 * time.Minute)

	log.Info("Pinging database...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := database.PingContext(ctx); err != nil {
		log.Fatalf("❌ Failed to ping SQLite database: %v", err)
	}
	log.Info("✅ SQLite Database connection successful.")

	// Note: Ensure your migration scripts inside RunMigration are updated to SQLite dialect!
	dashDB.RunMigration(database)

	a.db = database
}

func main() {
	app := NewApp()

	app.serverStart()
}
