package env

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv(filename string) {
	if err := godotenv.Load(filename); err != nil {
		msg := fmt.Sprintf("Error loading %s file", filename)
		log.Println(msg)
	}

	msg := fmt.Sprintf("Loaded %s file successfully", filename)
	log.Println(msg)
}
