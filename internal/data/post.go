package data

import (
	"crypto-news-server/internal/models"
	"encoding/json"

	"fmt"
	"os"
)

var Posts []models.Post

func LoadPosts() {
	file, err := os.Open("data/posts.json")

	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Posts); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}

func SavePost() {
	file, err := os.Create("data/posts.json")

	if err != nil {
		fmt.Println("Error file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)

	if err := encoder.Encode(Posts); err != nil {
		fmt.Println("Error decoding JSON: ", err)
	}
}