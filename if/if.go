
package main
import "fmt"

func main()  {

	age := 100

	if age < 13 {
		fmt.Println("you are underage")
	} else if age < 20{
		fmt.Println("you are a teenager")
	} else if age < 30 {
		fmt.Println("you are an adult")
	} else if age < 40 {
		fmt.Println("you are getting there")
	} else if age < 50 {
		fmt.Println("you are about to retire")
	} else {
		fmt.Println("you are a senior staff")
	}
	
}