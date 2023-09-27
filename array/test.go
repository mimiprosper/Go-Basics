
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

// list groceries items
func listGroceries()  {
	fmt.Println("grocery list is as follows:")
	for i := 0; i < groceriesLen; i++ {
		fmt.Println(storeGroceries[i])
	}
}


func main()  {
	addGroceries("coffe")
	addGroceries("bread")
	addGroceries("egg")
	addGroceries("milk")
	addGroceries("fish")
	listGroceries()
	//addGroceries("carrot")
	//addGroceries("oil")
	
}