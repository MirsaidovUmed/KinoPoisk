package pkg

import "kinopoisk/pkg/models"

func FillCharacters() {
	spiderMan := Cinemas[0]
	starWars := Cinemas[1]

	Characters = append(Characters, models.Character{
		Name:    "Питер Паркер",
		IsAlive: true,
		IsEvil:  false,
		Age:     16,
		Cinema:  spiderMan,
	})

	Characters = append(Characters, models.Character{
		Name:    "Мери Джейн",
		IsAlive: true,
		IsEvil:  false,
		Age:     25,
		Cinema:  spiderMan,
	})

	Characters = append(Characters, models.Character{
		Name:    "Дарт Вейдер",
		IsAlive: false,
		IsEvil:  true,
		Age:     75,
		Cinema:  starWars,
	})

	Characters = append(Characters, models.Character{
		Name:    "Принцесса Лея",
		IsAlive: false,
		IsEvil:  false,
		Age:     45,
		Cinema:  starWars,
	})
}
