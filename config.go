package main

// Configuration holds all global configs
type Configuration struct {
	LogToScreen      bool
	LogToFile        bool
	LogToDb          bool
	DbHost           string
	DbPort           uint16
	DbName           string
	DbUser           string
	DbPassword       string
	DbMaxConnections int
	LogFilePath      string
}
