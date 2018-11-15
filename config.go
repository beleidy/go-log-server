package main

// Configuration holds all global configs
type Configuration struct {
	DbHost           string
	DbPort           uint16
	DbName           string
	DbUser           string
	DbPassword       string
	DbMaxConnections int
}
