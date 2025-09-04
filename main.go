package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
    _ "github.com/golang-migrate/migrate/v4/source/file"

    "github.com/Aditya7880900936/postgres-golang/routers" // ✅ Import router package
)

func runMigrations(db *sql.DB) {
    driver, err := postgres.WithInstance(db, &postgres.Config{})
    if err != nil {
        log.Fatalf("Could not start SQL driver: %v", err)
    }

    m, err := migrate.NewWithDatabaseInstance(
        "file://migrations",
        "postgres", driver)
    if err != nil {
        log.Fatalf("Migration failed: %v", err)
    }

    if err := m.Up(); err != nil && err != migrate.ErrNoChange {
        log.Fatalf("Could not run migration: %v", err)
    }

    fmt.Println("Migrations applied successfully ✅")
}

func main() {
    if err := godotenv.Load(); err != nil {
		log.Println("⚠️  Warning: .env file not found, using system environment variables")
	}
    dbURL := os.Getenv("POSTGRES_URL")
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Run migrations automatically
    runMigrations(db)

    // ✅ Use router.Router() instead of creating a new one
    r := routers.Router()

    fmt.Println("Starting Server on the PORT 8080.....")
    log.Fatal(http.ListenAndServe(":8080", r))
}
