// slice is an array that can grow or shrink
// array has fixed lenght of items

package main
import "fmt"

func main()  {
	cards := [] string {"aces of diamond", newCard()} // slice example
	// function to add new items to the slice
	cards = append(cards, "aces of gold")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "queen of hearts"
}