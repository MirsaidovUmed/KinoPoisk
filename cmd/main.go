package main

import (
	"fmt"
	"kinopoisk/pkg"
)

func main() {
	pkg.FillCinemas()
	pkg.FillCharacters()

	fmt.Printf("%+v \n", pkg.Cinemas)

	fmt.Printf("%+v \n", pkg.Characters)

	goodCharacters := pkg.GetGoodCharacters()
	for _, character := range goodCharacters {
		fmt.Printf("Cinema Name: %s, Name: %s, IsAlive: %v\n", character.Cinema.Name, character.Name, character.IsAlive)
	}
}
