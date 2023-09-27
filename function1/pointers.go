
package main
import "fmt"

func max(i int, j int, k *int) int {
	fmt.Println("value of z:", k)
	if i > j {
		return i
	} else {
		return j
	}
}

func main()  {
	var c int
	fmt.Println("the address of c:", &c)
	max(10, 15, &c)
	fmt.Println(c)
}