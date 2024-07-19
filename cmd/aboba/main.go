package main

import (
	"aboba"
	"aboba/web"
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var (
		addr       string
		sqlAddr    string
		sessionKey string
	)

	fs := flag.NewFlagSet("aboba", flag.ExitOnError)
	fs.StringVar(&addr, "addr", ":4000", "HTTP service address")
	fs.StringVar(&sqlAddr, "sql-addr", "postgresql://root@127.0.0.1:26257/defaultdb?sslmode=disable", "SQL address")
	fs.StringVar(&sessionKey, "session-key", "secretkeynotcommit", "Session Key")
	if err := fs.Parse(os.Args[1:]); err != nil {
		return fmt.Errorf("parse flags: %v", err)
	}

	db, err := sql.Open("postgres", sqlAddr)
	if err != nil {
		return fmt.Errorf("open db: %w", err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		return fmt.Errorf("ping db: %w", err)
	}

	if err := aboba.MigrateSQL(context.Background(), db); err != nil {
		return fmt.Errorf("migrate sql: %w", err)
	}

	queries := aboba.New(db)
	svc := &aboba.Service{
		Queries: queries,
	}

	logger := log.New(os.Stderr, "", log.Ldate|log.Ltime|log.Llongfile)
	handler := &web.Handler{
		Logger:     logger,
		Service:    svc,
		SessionKey: []byte(sessionKey),
	}

	srv := &http.Server{
		Handler: handler,
		Addr:    addr,
	}

	defer srv.Close()

	logger.Printf("listening on %s", addr)

	err = srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("listen and serve: %w", err)
	}
	return nil
}
