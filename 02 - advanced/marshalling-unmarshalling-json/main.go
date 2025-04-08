package main

import (
	"encoding/json"
	"fmt"
)

// TODO
/*
JSON Marshalling and Unmarshalling in Go
- JSON marshalling is the process of converting Go data structures into JSON format
- JSON unmarshalling is the reverse process,
*/

// Marshaling JSON in Go (Encoding)
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Marshaling JSON in Go (Encoding)
	p := Person{Name: "Alice", Age: 30}

	// marshalling the struct into JSON
	jsonData, err := json.MarshalIndent(p, "", " ")

	if err != nil {
		fmt.Println("error marshalling:", err)
		return
	}
	fmt.Println(string(jsonData))

	// Unmarshalling JSON in Go (decoding)
	var p2 Person

	// unmarshalling the JSON data into the struct
	err2 := json.Unmarshal([]byte(jsonData), &p2)

	if err2 != nil {
		fmt.Println("error unmarshalling:", err)
		return
	}

	fmt.Println(p2)

	// handling missing or unknown fields in JSON - go will either ignore (if the value does not exist), or fill them with default values (if the value exists but is missing)
	jsonData2 := `{"name":"Alice"}`
	var p3 Person
	err3 := json.Unmarshal([]byte(jsonData2), &p3)
	if err3 != nil {
		fmt.Println("error unmarshalling", err3)
		return
	}
	fmt.Println(p3) // age defaults to 0

	// unknown fields
	jsonData3 := `{"name":"Alice", "age":25, "profession":"Programmer"}`
	var e Person
	err4 := json.Unmarshal([]byte(jsonData3), &e)

	if err4 != nil {
		fmt.Println("error", err4)
		return
	}
	fmt.Println(e) //profession is ignored

	// working with maps and slices in Go - json can be passed into maps and slices instead of structured types
	jsonData4 := `[{"name":"John","age":30}, {"name":"Jane","age":25}]`

	var f []Person

	err5 := json.Unmarshal([]byte(jsonData4), &f)
	if err5 != nil {
		fmt.Println("error marshalling", err5)
		return
	}

	fmt.Println(f)

	// handling custom json serialization - Go allows customization of structs, how they are marshaled and unmarshaled using custom methods
	//custom marshaling
	// func (p Person) MarshalJSON() ([]byte, error) {
	// 	return json.Marshal(struct {
	// 		FullName string `json:"full_name"`
	// 		Age int `json:"age"`
	// 	}{
	// 		FullName: p.Name,
	// 		Age: p.Age,
	// 	})

	// }

}
