package main

import "fmt"

func main() {
	maps()
}

func slices() {
	names := make([]string, 10)
	fmt.Println(names)

	names[0] = "João"
	names[1] = "Pedro"
	names[3] = "Miguel"
	names[4] = "Maria"
	names[5] = "José"
	fmt.Println(names)

	twoNames := names[0:2]
	allNames := names[:len(names)]

	fmt.Println(twoNames)
	fmt.Println(allNames)

	names = append(names, "Fernando")

	fmt.Println("New names ", names)
}

func theMatrix() {
	var matrix [3][3]int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {
			matrix[i][j] = j + i
		}
	}
	fmt.Println("matrix", matrix)
}

func simpleLists() {
	var ids [5]int
	ages := [10]int{23, 22, 20}

	fmt.Println("array:", ids, "\nlength:", len(ids))
	fmt.Println("Ages", ages)
}

func playingWithRange() {

	var ids = []int{1, 2, 3, 4, 5}

	for range ids {
		println("I'm inside the range :D")
	}

	println("------- With Index and Value -------")

	for index, value := range ids {
		println("Index >>", index, "Value >>", value)
	}
}

// A lat, lng struct
type Coordinate struct {
	lat, lng float64
}

func maps() {
	var places = map[string]Coordinate{
		"Bell Labs": Coordinate{40.68433, -74.39967},
		"Google":    Coordinate{37.42202, -122.08408},
	}

	places["Another place"] = Coordinate{
		lat: 30.68433,
		lng: -54.39967,
	}

	fmt.Println(places)
	fmt.Println(places["Bell Labs"])
}
