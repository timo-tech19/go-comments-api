package main

import "fmt"

// Run - responsible for startup and instantiation
// of our go application
func Run() error {
	fmt.Println("Starting app...")
	return nil
}

func main() {
	fmt.Println("Go Comments REST API")

	// in go you can declare a var in a control statement
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
