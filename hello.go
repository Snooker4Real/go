package main

import "fmt"

func main() {
	fmt.Println("Hi man!")

	name, age, likesGo := "Jonathan", 22, true

	if likesGo {
		fmt.Println("My name is " + name + ", I'm " + string(age) + " and I like Go !")
	} else {
		fmt.Println("My name is " + name + ", I'm " + string(age) + " and I does'nt like Go !")
	}

}
