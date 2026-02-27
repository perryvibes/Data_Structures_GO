package main

import (
	"fmt"
	"strings"
)

type Movie struct {
	id     int
	time   int
	name   string
	budget float64
	minAge rune
}

func ShowVector(vector []*Movie, nrElements int) {
	for i := 0; i < nrElements; i++ {
		fmt.Println(*vector[i])
	}
}

func CopyNElements(m []*Movie, nrElements int, nrCopiedElements int) []*Movie {
	newVector := make([]*Movie, nrElements)
	for i := 0; i < nrCopiedElements; i++ {
		newVector[i] = m[i]
	}
	return newVector
}

func CopyOnlySomeElements(vector []*Movie, nrElements int, maximBudget float64, newVector *[]*Movie, dimension *int) {
	*dimension = 0
	for i := 0; i < nrElements; i++ {
		if vector[i].budget < maximBudget {
			*dimension++
		}
	}
	*newVector = make([]*Movie, *dimension)
	var k int = 0
	for i := range nrElements {
		if vector[i].budget < maximBudget {
			(*newVector)[k] = vector[i]
			k++
		}
	}
}

func getFirstConditionedElement(vector []*Movie, nrElements int, searchedName string) Movie {
	var s Movie
	s.id = -1
	s.name = "n/a"
	for i := range nrElements {
		if strings.Compare(vector[i].name, searchedName) == 0 {
			return *vector[i]
		}
	}
	return s
}

func main() {

	f1 := Movie{1, 120, "Don't look up", 60, 16}
	fmt.Println(f1)
	fmt.Println("=========================================================================")

	// slices
	var nrMovies int = 3
	var movies []*Movie

	movies = append(movies, &f1)
	movies = append(movies, &Movie{2, 100, "Interstellar", 30.5, 16})
	movies = append(movies, &Movie{3, 90, "Fight Club", 20, 16})

	// arrays (with strict length)
	// but you have to modify ShowVector parameter from vector[]*Movie to vector[3]*Movie
	//movies := [3]*Movie{}
	//movies[0] = &f1
	//movies[1] = &Movie{2, 100, "Interstellar", 30.5, 16}
	//movies[2] = &Movie{3, 90, "Fight Club", 20, 16}
	ShowVector(movies, nrMovies)

	fmt.Println("=========================================================================")

	var copiedMovies int = 2
	newVector := CopyNElements(movies, nrMovies, copiedMovies)
	ShowVector(newVector, copiedMovies)

	fmt.Println("=========================================================================")

	var cheapVector []*Movie
	var limit float64 = 50
	var cheapVectorDimension int

	CopyOnlySomeElements(movies, nrMovies, limit, &cheapVector, &cheapVectorDimension)
	ShowVector(cheapVector, cheapVectorDimension)

	fmt.Println("=========================================================================")

	fmt.Println(getFirstConditionedElement(movies, nrMovies, "Interstellar"))

}
