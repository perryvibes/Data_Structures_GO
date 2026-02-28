package main

import "fmt"

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

}
