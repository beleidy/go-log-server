package main

import (
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

	pool.Prepare("log", "INSERT INTO logs (crawler, time, message) VALUES ($1, $2, $3)")

	for {
		entry := <-e
		go func() {
			_, err := pool.Exec(
				"log",
				entry.ID, entry.Time, entry.Message)
			check(err)
		}()
	}

}
