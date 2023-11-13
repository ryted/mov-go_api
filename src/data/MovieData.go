package MovieData

import (
	MovieType "github.com/ryted/mov-go_api/src/Model"
)

var (
	MovieList = []MovieType.Movie{
		{
			ID:    "1",
			Isbn:  "28471",
			Title: "The Last Heroes",
			Director: MovieType.Director{
				Firstname: "Jonathan",
				Lastname:  "Alexander",
			},
			ID:    "2",
			Isbn:  "28472",
			Title: "The Last Ronin",
			Director: MovieType.Director{
				Firstname: "Jonathan",
				Lastname:  "li",
			},
		},
	}
)
