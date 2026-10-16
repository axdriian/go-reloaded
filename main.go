package main

import (
	"fmt"
	"os"
	"strings"
)
	
	func main() {
		if len(os.Args) < 3 {
			fmt.Println("Usage: program <fichierEntree> <fichierSortie>")
			os.Exit(1)
		}
	
		fichierEntree := os.Args[1]		
		fichierSortie := os.Args[2]
	
		data, err := os.ReadFile(fichierEntree)
		if err != nil {
			fmt.Println("Erreur lors de la lecture du fichier :", err)
			return
		}

		var contenu string = string(data)
		mots := strings.Fields(contenu)
		for i, leMot := range mots {
			 if leMot ==("(up)"){
			 mots[i - 1] = strings.ToUpper(mots[i - 1])
			 }
			 	if leMot ==("(low)"){
					mots[i - 1] = strings.ToLower(mots[i - 1])
				}
					if leMot ==("(cap)"){
						mots[i - 1] = strings.Title(mots[i - 1])
					}
						if leMot ==("(hex)"){
						 mots[i - 1] = fmt.Sprintf("%x", mots[i - 1])
					}
					voyelles := "aeiouyAEIOUY"
					if mots[i+1][0] == voyelles {
						mots[i] = fmt.Print("an")
					} else {
						mots[i] = fmt.Print("a")
					}
			}
		}
	
	}
	err := os.WriteFile(os.Args[2], []byte(strings.Join(mots, " ")), 0644); err != nil {
		fmt.Println("Erreur lors de l'écriture du fichier :", err)
		return
	
}