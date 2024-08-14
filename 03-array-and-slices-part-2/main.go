package main

import "fmt"

var filmsOnDB = []string{
	"O poderoso chefão",
	"O poderoso chefão 2",
	"Batman: O Cavaleiro das Trevas",
	"12 Homens e uma Sentença",
	"A Lista de Schindler",
	"O Senhor dos Anéis: O Retorno do Rei",
	"Pulp Fiction: Tempo de Violência",
	"O Bom, o Mau e o Feio",
	"O Senhor dos Anéis: A Sociedade do Anel",
	"Forrest Gump",
}

func main() {
	// array of favorite movies_id from the user
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	//var films []
	films := make([]string, 0, 10) // this is a slice with a length of 0 and a capacity of 10
	// this makes in a way we don't need to resize the slice capacity every time we append a new element

	// IMP: this is good when we know the exact size of the slice we're going to use
	for _, id := range resultsFromApi {
		film := filmsOnDB[id-1]
		fmt.Println(len(films), cap(films))
		films = append(films, film)
		fmt.Println(len(films), cap(films))

	}

	fmt.Println(films)

}
