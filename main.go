package main

import (
    "database/sql"
    "fmt"
    "log"
    "net/http"
    "os"

    _ "github.com/lib/pq"
    "github.com/gorilla/mux"
    "github.com/joho/godotenv"
    "github.com/golang-migrate/migrate/v4"
    "github.com/golang-migrate/migrate/v4/database/postgres"
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
    godotenv.Load()
    dbURL := os.Getenv("POSTGRES_URL")
    db, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    // Run migrations automatically
    runMigrations(db)

    r := mux.NewRouter()
    fmt.Println("Starting Server on the PORT 8080.....")
    log.Fatal(http.ListenAndServe(":8080", r))
}
