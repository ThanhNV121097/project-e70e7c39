package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-e70e7c39/backend/migrations"
	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" { log.Fatal("DATABASE_URL is required") }
	connConfig, err := pgx.ParseConfig(databaseURL)
	if err != nil { log.Fatalf("invalid DATABASE_URL: %v", err) }
	db := sql.OpenDB(stdlib.GetConnector(*connConfig))
	defer db.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := migrate(ctx, db); err != nil { log.Fatalf("migration failed: %v", err) }
	if err := db.PingContext(ctx); err != nil { log.Fatalf("database unavailable: %v", err) }
	port := os.Getenv("PORT"); if port == "" { port = os.Getenv("APP_PORT") }; if port == "" { port = "8080" }
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) { if r.Method != http.MethodGet { http.NotFound(w, r); return }; if err := db.PingContext(r.Context()); err != nil { writeError(w, http.StatusServiceUnavailable, "database_unavailable", "Service unavailable."); return }; writeJSON(w, http.StatusOK, map[string]string{"status":"ok"}) })
	httpServer := &http.Server{Addr: ":" + port, Handler: mux, ReadHeaderTimeout: 5*time.Second}
	log.Printf("listening on %s", httpServer.Addr)
	log.Fatal(httpServer.ListenAndServe())
}

func migrate(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (version text PRIMARY KEY, applied_at timestamptz NOT NULL DEFAULT now())`); err != nil { return err }
	entries, err := migrations.Files.ReadDir("."); if err != nil { return err }; sort.Slice(entries, func(i, j int) bool { return entries[i].Name() < entries[j].Name() })
	for _, entry := range entries { name := entry.Name(); if entry.IsDir() || !strings.HasSuffix(name, ".up.sql") { continue }; var done bool; if err := db.QueryRowContext(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE version = $1)`, name).Scan(&done); err != nil || done { if err != nil { return err }; continue }; sqlBytes, err := migrations.Files.ReadFile(name); if err != nil { return err }; tx, err := db.BeginTx(ctx, nil); if err != nil { return err }; if _, err = tx.ExecContext(ctx, string(sqlBytes)); err == nil { _, err = tx.ExecContext(ctx, `INSERT INTO schema_migrations (version) VALUES ($1)`, name) }; if err != nil { tx.Rollback(); return fmt.Errorf("%s: %w", name, err) }; if err = tx.Commit(); err != nil { return err } }
	return nil
}

func writeJSON(w http.ResponseWriter, status int, value any) { w.Header().Set("Content-Type", "application/json"); w.Header().Set("Cache-Control", "no-store"); w.WriteHeader(status); _ = json.NewEncoder(w).Encode(value) }
func writeError(w http.ResponseWriter, status int, code, message string) { writeJSON(w, status, map[string]any{"error":map[string]string{"code":code,"message":message}}) }
