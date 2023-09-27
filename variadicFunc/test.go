package main

import "fmt"

var Groceries [] string // array of strings

// variadic function (... like spread operator in JS)
func addGroceries(a ...string) {
	Groceries = append(Groceries, a...)
}

// list groceries items without index
func listGroceries() {
	for _, grocery := range Groceries {
		fmt.Println(grocery)
	}
}

func main() {
	// effect of the variadic function
	addGroceries("coffee", "bread", "egg", "milk", "fish") 
	addGroceries("sugar")
	listGroceries() 
}
