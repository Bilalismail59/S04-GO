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
