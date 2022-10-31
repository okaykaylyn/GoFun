// this is the starting package for the program
package main

//after the package declaration, we import several others
import (
	"fmt"
)

// this is the main function, which is the entry point for the program
func main() {
	//string in go are unicode, which means they can contain any character... like emojis!
	fmt.Println("Hello World! 😀")
}
