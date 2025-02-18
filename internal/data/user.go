package data

import (
	"crypto-news-server/internal/models"
	"encoding/json"

	"fmt"
	"os"
)

var Users []models.User

func LoadUsers() {
	file, err := os.Open("data/users.json")

	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Users); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func SaveUser() {
	file, err := os.Create("data/users.json")

	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(Users); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}