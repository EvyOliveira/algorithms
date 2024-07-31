package main

import (
	"fmt"
	"math"
)

type Book struct {
	Thickness int
	Height    int
}

type Shelf struct {
	Thickness int
	Height    int
}

func calculateMinimumHeight(books []Book, maximumShelfWidth int) int {
	if len(books) == 0 {
		return 0
	}
	minimumHeights := make([]int, len(books)+1)
	for i := range minimumHeights {
		minimumHeights[i] = math.MaxInt32
	}
	minimumHeights[0] = 0

	for i := 1; i <= len(books); i++ {
		currentShelf := Shelf{}
		for j := i - 1; j >= 0; j-- {
			if currentShelf.Thickness+books[j].Thickness > maximumShelfWidth {
				break
			}
			currentShelf.Thickness += books[j].Thickness
			currentShelf.Height = max(currentShelf.Height, books[j].Height)
			minimumHeights[i] = min(minimumHeights[i], minimumHeights[j]+currentShelf.Height)
		}
	}
	return minimumHeights[len(books)]
}

func minimumHeightShelves(books [][]int, shelfWidth int) int {
	bookStruct := make([]Book, len(books))
	for i, book := range books {
		bookStruct[i] = Book{book[0], book[1]}
	}
	return calculateMinimumHeight(bookStruct, shelfWidth)
}

func main() {
	books := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	shelfWidth := 4
	minHeight := minimumHeightShelves(books, shelfWidth)
	fmt.Println("the minimum possible height that the total bookshelf is:", minHeight)
}
