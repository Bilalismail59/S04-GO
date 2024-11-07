package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Player struct {
	Nom    string `yaml:"nom"`
	Pseudo string `yaml:"pseudo"`
	Age    int    `yaml:"age"`
	Health int    `yaml:"health"`
	Bilal  int    `yaml:"bilal"`
}

// Map pour stocker les joueurs indexés par leur nom
var players = make(map[string]*Player)

// Méthode pour sauvegarder un joueur dans un fichier .yml
func (p *Player) save() error {
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(p.Nom+".yml", data, 0644)
}

// Méthode pour supprimer un joueur de la map et son fichier .yml
func (p *Player) del() error {
	delete(players, p.Nom)
	return os.Remove(p.Nom + ".yml")
}

// Méthode pour afficher les informations du joueur
func (p *Player) display() string {
	return fmt.Sprintf("Nom: %s, Pseudo: %s, Age: %d, Health: %d, Bilal: %d", p.Nom, p.Pseudo, p.Age, p.Health, p.Bilal)
}

// Fonction pour charger un joueur
func playerLoad(name string) (*Player, error) {
	if player, exists := players[name]; exists {
		return player, nil
	}

	// Charger depuis le fichier YAML s'il existe
	data, err := ioutil.ReadFile(name + ".yml")
	if err == nil {
		var player Player
		if err := yaml.Unmarshal(data, &player); err == nil {
			players[name] = &player
			return &player, nil
		}
	}

	// Créer un joueur si inexistant
	var pseudo string
	fmt.Printf("Pseudo du joueur: ")
	fmt.Scan(&pseudo)

	player := &Player{
		Nom:    name,
		Pseudo: pseudo,
		Age:    0,
		Health: 100,
		Bilal:  100,
	}
	players[name] = player
	return player, player.save()
}
