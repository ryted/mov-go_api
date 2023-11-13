package MovieType

import (
	"net/http"
)

type Movie struct {
	ID       string   `json:id`
	Isbn     string   `json:isbn`
	Title    string   `json:title`
	Director Director `json:director`
}

type Director struct {
	Firstname string `json:firstname`
	Lastname  string `json:lastname`
}

type Route struct {
	Type    string
	Path    string
	Handler http.HandlerFunc
	Methods []string
}
