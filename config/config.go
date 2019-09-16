package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

var ServerConfig = Server{
	"debug",
	8080,
	60 * time.Second,
	60 * time.Second,
}

func init() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	runMode := os.Getenv("RUN_MODE")
	port, _ := strconv.Atoi(os.Getenv("HTTP_PORT"))
	read, _ := strconv.Atoi(os.Getenv("READ_TIMEOUT"))
	write, _ := strconv.Atoi(os.Getenv("WRITE_TIMEOUT"))

	readTime := time.Duration(read) * time.Second
	writeTime := time.Duration(write)  * time.Second

	ServerConfig.RunMode = runMode
	ServerConfig.HttpPort = port
	ServerConfig.ReadTimeout = readTime
	ServerConfig.WriteTimeout = writeTime
}