package main

import (
	"os"
	"testing"
)

// Test pour créer, afficher, sauvegarder et supprimer un joueur
func TestPlayerLifecycle(t *testing.T) {
	player, err := playerLoad("TestPlayer")
	if err != nil {
		t.Fatalf("Erreur lors de la création du joueur : %v", err)
	}

	// Test de la fonction display
	displayText := player.display()
	expected := "Nom: TestPlayer, Pseudo: , Age: 0, Health: 100, Mana: 100"
	if displayText != expected {
		t.Errorf("display() = %v, attendu %v", displayText, expected)
	}

	// Vérifie si le fichier a été créé
	if _, err := os.Stat("TestPlayer.yml"); os.IsNotExist(err) {
		t.Errorf("Le fichier TestPlayer.yml n'a pas été créé.")
	}

	// Supprime le joueur
	if err := player.del(); err != nil {
		t.Errorf("Erreur lors de la suppression du joueur : %v", err)
	}

	// Vérifie si le fichier a été supprimé
	if _, err := os.Stat("TestPlayer.yml"); err == nil || !os.IsNotExist(err) {
		t.Errorf("Le fichier TestPlayer.yml n'a pas été supprimé.")
	}
}
