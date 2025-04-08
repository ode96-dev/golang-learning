package main

/*
Go modules
- a collection of Go packages that are versioned together
- contains the code, a go.mod file, and any dependencies required for the module
- allow for managing external dependecies(3rd party libaries)
- organizing code into self-contained units called modules
- they simplify versioning, compatibility, and the distribution of Go code.
- helps in managing dependencies without relying on a global workspace
go.mod file:
- specifies the module's name, its dependencies, and other module-specific metadata
Go module path:
- typically a url to the repository hosting the module
go.sum file
- records the cryptographic hashes of the module and their dependencies, ensuring reproducibility
*/

func main() {
	/*
		   Structure
		   - go.mod defines the module and its dependencies
		   - go.sum ensures the integrity of dependencies by storing their hash values.
		   - An example project structure willl be:
		   		/my-module
					go.mod - defines the module and dependencies
					go.sum - stores checksumsof dependencies for integrity verification
					main.go - entry point of the program
					/pkg
						util.go - contains utility function or other code used in the module
	*/
}
