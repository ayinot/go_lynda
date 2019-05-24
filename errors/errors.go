package main

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	//COnfiguration file
}

func readConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()
	cfg := &Config{}

	return cfg, nil
}

func setUpLogging() {
	out, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return
	}
	log.SetOutput(out)
}

func main() {
	setUpLogging()
	cfg, err := readConfig("./errors/config.toml")
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		log.Printf("error: %v", err)
		os.Exit(1)
	}
	fmt.Println(cfg)
}
