package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"sort"
)

func PgxPlay() {
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database %v\n", err)
		os.Exit(1)

	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "")
	defer rows.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	sort.SliceIsSorted()
}
