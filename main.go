package main

import (
	"embed"
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

//go:embed cookies/*
var cookies embed.FS

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	dir, err := cookies.ReadDir("cookies")
	if err != nil {
		log.Fatal(err)
	}

	var fortunes []string // Slice to hold fortune phrases

	for _, d := range dir {
		if d.Type().IsDir() {
			sub, err := cookies.ReadDir("cookies/" + d.Name())
			if err != nil {
				log.Fatal(err)
			}
			for _, s := range sub {
				data, err := cookies.ReadFile("cookies/" + d.Name() + "/" + s.Name())
				if err != nil {
					log.Fatal(err)
				}
				content := string(data)
				// Split content by '%' and trim spaces
				phrases := strings.Split(content, "%")
				for _, phrase := range phrases {
					trimmedPhrase := strings.TrimSpace(phrase)
					if trimmedPhrase != "" {
						fortunes = append(fortunes, trimmedPhrase) // Add non-empty phrases to fortunes
					}
				}
			}
		} else {
			data, err := cookies.ReadFile("cookies/" + d.Name())
			if err != nil {
				log.Fatal(err)
			}
			content := string(data)
			// Split content by '%' and trim spaces
			phrases := strings.Split(content, "%")
			for _, phrase := range phrases {
				trimmedPhrase := strings.TrimSpace(phrase)
				if trimmedPhrase != "" {
					fortunes = append(fortunes, trimmedPhrase) // Add non-empty phrases to fortunes
				}
			}
		}
	}

	// Randomly select a fortune phrase
	if len(fortunes) > 0 {
		randomIndex := rand.Intn(len(fortunes)) // Get a random index
		// Print the selected fortune to stdout
		// fmt.Println("Random Fortune Cookie Phrase:")
		fmt.Println(fortunes[randomIndex]) // Output the selected fortune
	} else {
		fmt.Println("No fortunes found.")
	}
}
