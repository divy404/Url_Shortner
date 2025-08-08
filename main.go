package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"
)

type URL struct {
	ID           string    `json:"id"`
	OriginalURL  string    `json:"original_url"`
	ShortURL     string    `json:"short_url"`
	CreationDate time.Time `json:"creation_date"`
}

var urlDB = make(map[string]URL)

func generateShortURL(OriginalURL string) string {
	hasher := md5.New()               // this is used to give hashing
	hasher.Write([]byte(OriginalURL)) // it converts the originalURL string to a byte slice
	data := hasher.Sum(nil)
	hash := hex.EncodeToString(data)
	fmt.Println("EncodeToString: ", hash)
	fmt.Println("final string: ", hash[:8])
	return hash[:8]
}

func createURL(originalURL string) string {
	shortURL := generateShortURL(originalURL)
	id := shortURL
	urlDB[id] = URL{
		ID: id,
		OriginalURL: originalURL,
		ShortURL: shortURL,
		CreationDate: time.Now(),
	}
	return  shortURL
}

func getURL (id string) (URL, error) {
	url,ok := urlDB[id]
	if !ok {
		return  URL{}, errors.New("URL not found")
	}
	return  url, nil 
}
//  we will continue from here 
func main() {
	fmt.Println("Url Shortner")
	OriginalURL := "https://github.com/divy404"
	generateShortURL(OriginalURL)
}
