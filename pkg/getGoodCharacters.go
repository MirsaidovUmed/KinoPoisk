package pkg

import (
	"kinopoisk/pkg/models"
)

func GetGoodCharacters() []models.Character {
	result := make([]models.Character, 0)
	for _, character := range Characters {
		if character.IsEvil == false {
			result = append(result, character)
		}
	}

	return result
}
