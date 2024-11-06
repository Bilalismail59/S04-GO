package main

import "log"

var (
	token = "c732a4f732342956ec521490b59a7dce"
)

func main() {
	ville := "najac"

	coord, err := getCoordFfromCity(ville)
	if err != nil {
		log.Fatalln(err)

	}
	log.Printf("%#v", coord)
}
