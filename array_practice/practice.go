package main

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	array := [3]string{"hobby 1", "hobby 2", "hobby 3"}
	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1:])

	slice := array[0:2]
	reSlice := slice[1:3]

	fmt.Println(slice)
	fmt.Println(reSlice)

	dynamicArray := []string{"goal 1", "goal 2"}
	dynamicArray[1] = "new goal 2"
	dynamicArray = append(dynamicArray, "goal 3")

	fmt.Println(dynamicArray)

	dynamicStructArray := []Product{
		{
			"title1",
			"id1",
			9.99,
		},
		{
			"title2",
			"id2",
			10.99,
		},
	}

	dynamicStructArray = append(dynamicStructArray, 
		Product{
			"title3",
			"id3",
			11.99,
		},
	)

	fmt.Println(dynamicStructArray)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
