package main

import "fmt"

var spending = []float64{9.50, 8.00, 10.20, 7.40, 100.00}

func main() {

	activity1()
	spending = append(spending, 7.0)
	activity3()

	// a bit of activity 4 + removing an item from the slice?
	// removing entry 3
	set1 := spending[0:2]
	set2 := spending[3:5]
	for _, value := range set2 {
		set1 = append(set1, value)
	}
	spending = set1
	fmt.Println(spending)
}

func activity1() {
	fmt.Println(spending)
	fmt.Println(len(spending))
	fmt.Println(cap(spending))
}

func activity3() {
	spending[2] = 9.8
	fmt.Println(spending)
}
