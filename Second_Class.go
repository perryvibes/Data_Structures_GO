package main

import "fmt"

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
}
