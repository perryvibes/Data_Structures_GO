package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Car struct {
	id         int
	nrDoors    int
	price      float32
	model      string
	driverName string
	series     rune
}

func showVector(vector []*Car) {
	for _, value := range vector {
		fmt.Println(*value)
	}
}

func addNewCar(vector *[]*Car, newCar *Car) []*Car {
	*vector = append(*vector, newCar)
	return *vector
}

func readCarsFromFile(fptr *os.File) {
	scanner := bufio.NewScanner(fptr)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %s", err)
	}
}

func main() {

	cars := []*Car{
		{1, 3, 5000, "Astra", "Ionescu", 'A'},
		{2, 5, 8000, "Mokka", "Vasilescu", 'S'},
		{3, 4, 9000, "Passat", "Gigel", 'I'},
	}
	showVector(cars)

	fmt.Println("==============================================")

	newCar := Car{4, 4, 10000, "A7", "July", 'S'}
	cars = addNewCar(&cars, &newCar)
	showVector(cars)

	fmt.Println("==============================================")

	file, err := os.Open("cars.txt")
	if err != nil {
		fmt.Println("Error opening file...")
	}
	defer file.Close()
	readCarsFromFile(file)

	fmt.Println("==============================================")
}
