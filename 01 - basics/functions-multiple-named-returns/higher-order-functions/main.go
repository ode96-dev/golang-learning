package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

/*
- High Order Funcion is a function that takes another function as an argument or returns a function
- This enables dynamic behavior, callbacks, and functional programming patterns in Go.
- Why Use HOF:
 1. encapsulation - encapsulates logic into reusable functions
 2. code reusability - avoids repeating code
 3. customizable functions - pass different behaviors dynamically.
*/

/*
- Basic
1. applyOperation accepts two integers and a function, operation
2. add or subtract are passed as an argument
3. the function will execute dynamically based on the argument provided
*/
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// functions to be passed, add
func add(x, y int) int {
	return x + y
}

// functions to be passed, subtract
func subtract(x, y int) int {
	return x - y
}

/*
- Returning a function (Function Factory)
- HOF can return a function, allowing dynamic function creation
- with createMultiplier(), no need for double(), and triple() functions
*/
//function that returns another function
func createMultiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

/*
- Hashing utility
- HOF can be used for hashing algorithms
- getHasher("sha256")/getHasher("sha512") returns SHA-256/SHA-512 hashing function
- this allows for dynamic selection of encryption algorithm
*/

func getHasher(algorithm string) func(string) string {
	return func(data string) string {
		if algorithm == "sha256" {
			hash := sha256.Sum256([]byte(data))
			return fmt.Sprintf("%x", hash)
		} else if algorithm == "sha512" {
			hash := sha512.Sum512([]byte(data))
			return fmt.Sprintf("%x", hash)
		}
		return "Unknown algorithm"
	}
}

/*
- Filtering Data with HOF
- filteredUsers takes a filtering function as an argument
- isAdmin is used to filter on admins dynamically.
*/

type User struct {
	Name string
	Role string
}

// HOF to filter users
func filterUsers(users []User, criteria func(User) bool) []User {
	var filtered []User
	for _, user := range users {
		if criteria(user) {
			filtered = append(filtered, user)
		}
	}
	return filtered
}

/*
- Mapping Data (Transforming Data)
- Go does not have an in-built map() function, but it can be implemented
*/
//generic map function
func mapInts(numbers []int, transformer func(int) int) []int {
	var result []int
	for _, num := range numbers {
		result = append(result, transformer(num))
	}
	return result
}

func main() {

	// Basic
	fmt.Println("Addition: ", applyOperation(5, 4, add))
	fmt.Println("Subtraction: ", applyOperation(5, 4, subtract))

	/*
		- Anonymous/Lambda Functions as Arguments
		- instead of defining multiplication(i1,i2 int), we pass the lambda function directly
		- good for short-lived operations that do not need a separate function
	*/
	result := applyOperation(10, 4, func(i1, i2 int) int {
		return i1 * i2
	})
	fmt.Println("Multiplication: ", result)

	// Function that returns another function
	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Println("double 4:", double(4))
	fmt.Println("triple 3:", triple(3))

	// Hashing utility
	sha256Hasher := getHasher("sha256")
	sha512Hasher := getHasher("sha512")

	password := "securepassword"

	fmt.Println("SHA-256 Hash:", sha256Hasher(password))
	fmt.Println("SHA-512 Hash:", sha512Hasher(password))

	// Filtering Data with HOF
	users := []User{
		{"Alice", "Admin"},
		{"Bob", "User"},
		{"Charlie", "Admin"},
		{"David", "Admin"},
	}
	//filtering function
	isAdmin := func(u User) bool {
		return u.Role == "Admin"
	}

	admins := filterUsers(users, isAdmin)
	fmt.Print("Admins:", admins)

	//Mapping Data (Transforming Lists)
	numbers := []int{1, 2, 3, 4, 5}

	square := func(n int) int {
		return n * n
	}

	squaredNumbers := mapInts(numbers, square)
	fmt.Println("Square numbers:", squaredNumbers)

}
