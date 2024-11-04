package main

import "fmt"

func main() {
	fmt.Println("Hello world !")

	somme := Ajouter(3, 4)
	fmt.Printf("La somme de 3 et 4 est : %v\n", somme)

	produit := multiplication(3, 4)
	fmt.Printf("La multiplication de 3 et 4 est : %v\n", produit)

	quotient := division(12, 4)
	fmt.Printf("La division de 12 par 4 est : %v\n", quotient)
}
