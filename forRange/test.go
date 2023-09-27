
package main
import "fmt"

const groceriesCap int  = 5 // number of array elements possible
var groceriesLen int // number of items in our array currently
var storeGroceries [groceriesCap] string // array of strings

// add groceries items
func addGroceries(a string)  {
	if groceriesLen < groceriesCap  {
		storeGroceries[groceriesLen] = a  
		groceriesLen++
	} else {
		fmt.Println("error !!! you have go pass the limit")
	}
}

// list groceries items without index
func listGroceries()  {
	for _, grocery := range storeGroceries {
		fmt.Println( grocery)
	}
}

// list groceries items with index
// func listGroceries()  {
// 	for i, grocery := range storeGroceries {
// 		fmt.Println(i, grocery)
// 	}
// }


func main()  {
	addGroceries("coffee")
	addGroceries("bread")
	addGroceries("egg")
	addGroceries("milk")
	addGroceries("fish")
	listGroceries()
	//addGroceries("carrot")
	//addGroceries("oil")	
}