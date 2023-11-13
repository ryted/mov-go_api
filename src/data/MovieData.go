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
		},
	}
)
