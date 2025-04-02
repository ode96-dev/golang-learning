package main

import "fmt"

/*
- unordered and changeable collection that does not allow duplicates
- its length is the number of its elements. utilize len() to find
- different ways to create maps: using var, :=, make()
*/

func main() {
	fmt.Println("======Maps, storing key-value pairs and are useful for quick lookups======")
	// 1. using :=
	employees := map[string]int{"Alice": 29, "Bob": 25}
	fmt.Println(employees)
	fmt.Println(employees["Alice"])

	// 2. using var
	var car = map[string]string{"brand": "Ford", "model": "Raptor", "year": "2025"}
	fmt.Println(car)

	// 3. using make()
	car2 := make(map[string]string) //here, car2 is empty
	car2["brand"] = "Toyota"
	car2["model"] = "Harrier"
	car2["year"] = "2025"
	fmt.Println(car2)

	// another way to make an empty map other than the make() function is through
	var x map[string]int
	fmt.Println(x)

	//updating and adding map elements -  map_name[key] = value
	car2["model"] = "Hilux" //updating an element
	car2["color"] = "Black" // adding an element
	fmt.Println(car2)

	//removing element from map - delete()
	delete(car, "year")
	fmt.Println(car)

	// checking for specific elements in a map = val, ok:=map_name[key]
	val1, ok1 := car2["brand"] // checks for existing key and its value
	val2, ok2 := car["color"]  // checks for non-exiting key and its value
	_, ok3 := car2["color"]    // only checks for existing key and not its value
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(ok3)

	//Maps are references
	// if we change car2 using b, car will change
	b := car2
	fmt.Println(b)
	b["price"] = "30,000"
	fmt.Println(car2)
	fmt.Println(b)

	//iterating over maps - using range, no specific order
	for k, v := range car2 {
		fmt.Printf("%v : %v, ", k, v)
	}

	//iterate over maps - in a specific order
	var o []string //defining the order
	o = append(o, "brand", "model", "color", "year", "price")

	for _, element := range o {
		fmt.Printf("%v : %v,", element, car2[element])
	}

}
