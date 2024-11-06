# gotest
Juste pour tester un environnement go

## paquet principal

importer  "journal"

type  Joueur  struct {
	pseudo  chaîne
	âge     int
	hp      int
	mp      int
	   chaîne de sexe
	intel   int
}

func ( p  * Joueur ) setAge ( âge  int ) {
	p . âge  =  âge
}

var (
	joueur  =  créer ( carte [ chaîne ] Joueur )
)

fonction  principale () {
	//var joueur1 Joueur
	joueur2  :=  Joueur { pseudo : "saryon" , age : 15 , hp : 10 , mp : 10 , sexe : "" , intel : 12 }
	joueur [ "saryon" ] =  joueur2
	enregistrer . Printf ( "%#v" , joueur2 )
	enregistrer . Printf ( "L'âge de %v : %v" , joueur2 . pseudo , joueur2 . age )
	joueur2 . âge  =  16
	enregistrer . Printf ( "L'âge de %v : %v" , joueur2 . pseudo , joueur2 . age )
	joueur2 . setAge ( 17 )
	enregistrer . Printf ( "L'âge de %v : %v" , joueur2 . pseudo , joueur2 . age )
}

## package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Coordinates struct {
	Lat float32 `json:"lat"`
	Lon float32 `json:"lon"`
}

func getCoordFfromCity(city string) (Coordinates, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/geo/1.0/direct?q=%s&appid=%s", city, token)
	resp, err := http.Get(url)
	if err != nil {
		return Coordinates{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Coordinates{}, err
	}

	var resultat []Coordinates
	if err := json.Unmarshal(body, &resultat); err != nil {
		return Coordinates{}, err
	}

	// Check if resultat is empty
	if len(resultat) == 0 {
		return Coordinates{}, fmt.Errorf("no coordinates found for city: %s", city)
	}

	return resultat[0], nil
}


## 
