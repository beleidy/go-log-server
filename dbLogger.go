package main

import (
	"fmt"

	"github.com/jackc/pgx"
)

func dbLogger(e chan Entry) {
	pgxConfig := pgx.ConnConfig{
		Host:     config.DbHost,
		Port:     config.DbPort,
		Database: config.DbName,
		User:     config.DbUser,
		Password: config.DbPassword}

	poolConfig := pgx.ConnPoolConfig{
		ConnConfig:     pgxConfig,
		MaxConnections: config.DbMaxConnections}

	pool, err := pgx.NewConnPool(poolConfig)
	check(err)

	pool.Prepare("log", "INSERT INTO logs (crawler, time, url, message) VALUES ($1, $2, $3, $4)")

	for {
		entry := <-e
		fmt.Println(entry)
		go func() {
			_, err := pool.Exec(
				"log",
				entry.ID, entry.Time, entry.URL, entry.Message)
			check(err)
		}()
	}

}
