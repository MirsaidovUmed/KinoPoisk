package pkg

import "kinopoisk/pkg/models"

func FillCinemas() {
	Cinemas = append(Cinemas, models.Cinema{
		Name:  "Человек-паук",
		Genre: "фантастика",
		Year:  2010,
	})

	Cinemas = append(Cinemas, models.Cinema{
		Name:  "Звездные воины",
		Genre: "фантастика",
		Year:  1998,
	})
}
