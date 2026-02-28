package main

import (
	"fmt"
)

type Movie struct {
	id     int
	time   int
	name   string
	budget float64
	minAge rune
}

func ShowVector(vector []*Movie) {
	for _, value := range vector {
		fmt.Println(*value)
	}
}

func CopyNElements(m []*Movie, nrCopiedElements int) []*Movie {
	if nrCopiedElements > len(m) {
		nrCopiedElements = len(m)
	}
	newVector := make([]*Movie, nrCopiedElements)
	copy(newVector, m[:nrCopiedElements])
	return newVector
}

func CopyOnlySomeElements(vector []*Movie, maximBudget float64, newVector []*Movie, dimension *int) []*Movie {
	*dimension = 0
	for _, value := range vector {
		if value.budget < maximBudget {
			*dimension++
		}
	}
	newVector = make([]*Movie, *dimension)
	copy(newVector, vector[:(*dimension)])
	return newVector
}

func getFirstConditionedElement(vector []*Movie, searchedName string) Movie {
	for _, value := range vector {
		if value.name == searchedName {
			return *value
		}
	}
	return Movie{id: -1, name: "n/a"}
}

func main() {

	f1 := Movie{1, 120, "Don't look up", 60, 16}
	fmt.Println(f1)
	fmt.Println("=========================================================================")

	// slices
	movies := []*Movie{
		&f1,
		{2, 100, "Interstellar", 30.5, 16},
		{3, 90, "Fight Club", 20, 16},
	}

	// arrays (with strict length)
	// but you have to modify ShowVector parameter from vector[]*Movie to vector[3]*Movie
	//movies := [3]*Movie{}
	//movies[0] = &f1
	//movies[1] = &Movie{2, 100, "Interstellar", 30.5, 16}
	//movies[2] = &Movie{3, 90, "Fight Club", 20, 16}
	ShowVector(movies)

	fmt.Println("=========================================================================")

	copiedMovies := 2
	newVector := CopyNElements(movies, copiedMovies)
	ShowVector(newVector)

	fmt.Println("=========================================================================")

	var cheapVector []*Movie
	var limit float64 = 100
	var cheapVectorDimension int

	cheapVector = CopyOnlySomeElements(movies, limit, cheapVector, &cheapVectorDimension)
	ShowVector(cheapVector)

	fmt.Println("=========================================================================")

	fmt.Println(getFirstConditionedElement(movies, "Interstellar"))

}
