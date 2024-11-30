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
    rand.New(rand.NewSource(time.Now().UnixNano()))

    dir, err := cookies.ReadDir("cookies")
    if err != nil {
        log.Fatal(err)
    }

    var fortunes []string

    // WARN: not scalable
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
                phrases := strings.Split(content, "%")
                for _, phrase := range phrases {
                    trimmedPhrase := strings.TrimSpace(phrase)
                    if trimmedPhrase != "" {
                        fortunes = append(fortunes, trimmedPhrase)
                    }
                }
            }
        } else {
            data, err := cookies.ReadFile("cookies/" + d.Name())
            if err != nil {
                log.Fatal(err)
            }
            content := string(data)
            phrases := strings.Split(content, "%")
            for _, phrase := range phrases {
                trimmedPhrase := strings.TrimSpace(phrase)
                if trimmedPhrase != "" {
                    fortunes = append(fortunes, trimmedPhrase)
                }
            }
        }
    }

    if len(fortunes) > 0 {
        randomIndex := rand.Intn(len(fortunes))
        fmt.Println(fortunes[randomIndex])
    } else {
        fmt.Println("No fortunes found.")
    }
}
