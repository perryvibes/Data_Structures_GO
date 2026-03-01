package main

import (
	"bufio"
	"fmt"
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

//// Reading files requires checking most calls for errors
//func check(e error) {
//	if e != nil {
//		panic(e)
//	}
//}

func readCarFromFile(fptr *os.File) {
	reader := bufio.NewReader(fptr)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
		}
		fmt.Println(line)
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

	readCarFromFile(file)

	fmt.Println("==============================================")
}
