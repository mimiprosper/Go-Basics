package main

import "fmt"

func main() {

	for day := 1; day <= 12; day++ {

		fmt.Println("on the", day, "day of xmas, my true love sent to: ")

		switch day {
		case 12:
			fmt.Println("12 drummers drumming")
			fallthrough
		case 11:
			fmt.Println("11 pipers piping")
			fallthrough
		case 10:
			fmt.Println("10 lords a leaping")
			fallthrough
		case 9:
			fmt.Println("9 ladies dancing")
			fallthrough
		case 8:
			fmt.Println("8 maids milking")
			fallthrough
		case 7:
			fmt.Println("7 swans swimming")
			fallthrough
		case 6:
			fmt.Println("6 geese laying")
			fallthrough
		case 5:
			fmt.Println("5 golden rings")
			fallthrough
		case 4:
			fmt.Println("4 calling birds")
			fallthrough
		case 3:
			fmt.Println("3 french hens")
			fallthrough
		case 2:
			fmt.Println("2 turtle doves")
			fallthrough
		case 1:
			fmt.Println("a partridge in a pear tree")

		}

		fmt.Println("")
	}

}
