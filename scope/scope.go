package main
import "fmt"

var wow int = 10   //global scope (package level)

func testexit()  {
	wow := 20      // function level (scope)
	fmt.Println(&wow, wow)
}

func main()  {
	fmt.Println(&wow, wow)
	testexit()
	fmt.Println(&wow, wow)
}