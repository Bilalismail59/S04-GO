package main

import "log"

type Player struct {
	pseudo string
	age    int
	hp     int
	mp     int
	sexe   string
	intel  int
}

func (p *Player) setAge(age int) {
	p.age = age

}

var (
	player = make(map[string]Player)
)

func main() {
	//var joueur1 Player
	joueur2 := Player{pseudo: "Bilal", age: 15, hp: 10, mp: 10, sexe: "m", intel: 12}
	player["saryon"] = joueur2
	log.Printf("%#v", joueur2)

	log.Printf("L'age de %v : %v", joueur2.pseudo, joueur2.age)
	joueur2.age = 16
	log.Printf("L'age de %v : %v", joueur2.pseudo, joueur2.age)
	joueur2.setAge(17)
	log.Printf("L'age de %v : %v", joueur2.pseudo, joueur2.age)
}
