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

	"github.com/charmbracelet/log"
)

// /data/provider-alpha.json,/data/provider-beta.json,/data/provider-gamma.json

const (
	host = "0.0.0.0"
	port = ":3000"
)

type App struct {
	db     *sql.DB
	Server *http.Server
}

func NewApp() *App {
	// err := godotenv.Load()

	//if err != nil {
	//	log.Fatal("Error loading .env file, relying on environment variables", "error", err)
	//}

	app := new(App)

	app.serverSetup()

	app.dbSetup()

	return app
}

func (a *App) serverSetup() {
	mux := http.NewServeMux()

	mux.HandleFunc("/healthz", handlers.HealthzHandler)

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
		log.Info("Server starting on port 8080...")
		if err := a.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v\n", err)
		}
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
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbSSLMode := os.Getenv("DB_SSLMODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		log.Fatal("Database configuration environment variables are not fully set.")
	}
	if dbSSLMode == "" {
		dbSSLMode = "disable"
		log.Warn("DB_SSLMODE not set, defaulting to 'disable'. Ensure this is secure for production.")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbHost, dbPort, dbUser, dbPassword, dbName, dbSSLMode)

	log.Info("Connecting to PostgreSQL database...")
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed to open PostgreSQL database: %v", err)
	}

	// Configure connection pool
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(15)
	db.SetConnMaxLifetime(30 * time.Minute)
	db.SetConnMaxIdleTime(30 * time.Minute)

	log.Info("Pinging database...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("❌ Failed to ping PostgreSQL database: %v", err)
	}
	log.Info("✅ PostgreSQL Database connection successful.")

	dashDB.RunMigration(db)

	a.db = db
}

func main() {
	app := NewApp()

	http.HandleFunc("/healthz", healthzHandler)

	port := ":8080"
	fmt.Printf("Starting server on port %s...\n", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("Server crashed: ", err)
	}
}
