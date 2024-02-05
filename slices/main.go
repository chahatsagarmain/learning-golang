package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"Apples", "Mangoes"}
	fmt.Printf("The type of fruitList is %T\n", fruitList)
	fmt.Println(fruitList)

	fruitList = append(fruitList, "Pineapple")
	fmt.Println(fruitList)

	fruitList = fruitList[1:]
	fmt.Println(fruitList)

	// 5 Length and 10 as capacity
	var newList = make([]int, 5, 10)
	fmt.Printf("The type of newList is %T\n", newList)
	fmt.Println("The initialised values are : ", newList)

	var newnewList = new([]int)
	fmt.Printf("The type of newnewList is %T\n", newnewList)
	fmt.Println("The initialised values are : ", newnewList)

	newList[0] = 1
	newList = append(newList, 2, 3, 4, 5)

	fmt.Println("The values are : ", newList)

	fmt.Println("List is sorted ? : ", sort.SliceIsSorted(newList, func(i, j int) bool {
		return newList[i] < newList[j]
	}))

	sort.Ints(newList)

	fmt.Println("The sorted values are : ", newList)
	fmt.Println("List is sorted ? : ", sort.SliceIsSorted(newList, func(i, j int) bool {
		return newList[i] < newList[j]
	}))

	//remove a element at index i 
	var index int = 2;
	newList = append(newList[:index],newList[index+1:]...)
	fmt.Println("The values are : ", newList)
	
}
