package main

import "fmt"

type Dinner struct {
	eaterID    int
	foodMenuID int
}

type Menu struct {
	foodMenuId int
	count      int
}

func main() {
	fmt.Println("Hello!")

	newDinner := []Dinner{
		{eaterID: 1, foodMenuID: 101},
		{eaterID: 2, foodMenuID: 105},
		{eaterID: 3, foodMenuID: 103},

		{eaterID: 1, foodMenuID: 101},
		{eaterID: 4, foodMenuID: 102},
		{eaterID: 3, foodMenuID: 104},
		{eaterID: 2, foodMenuID: 101},
		{eaterID: 1, foodMenuID: 106},
	}
	:
	fmt.Println(newDinner)

	duplicateEntries := make(map[string]bool)

	for _, val := range newDinner {
		if duplicateEntries[val] {
			fmt.Println("", val.eaterID, val.foodMenuID)
		}
		
	}

}
