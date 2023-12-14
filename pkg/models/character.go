package models

type Character struct {
	IsAlive bool   // является ли персонаж живым
	IsEvil  bool   // является ли персонаж злодеем
	Name    string // имя персонажа
	Age     int    // возраст персонажа
	Cinema  Cinema // фильм в котором он участвует
}
