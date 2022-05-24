package main

import (
	"context"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nnabeyang/sakila-app/ent"
	"github.com/nnabeyang/sakila-app/ent/migrate"
)

func main() {
	entOptions := []ent.Option{}
	entOptions = append(entOptions, ent.Debug())
	client, err := ent.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME")), entOptions...)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background(),
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
