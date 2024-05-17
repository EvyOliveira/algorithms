package main

import (
	"fmt"
	"strings"
	"time"
)

type Searcher interface {
	Search(target string, data []Name) int
}

type BinarySearch struct{}

type Name struct {
	FullName string
}

func main() {
	sortedNames := generateRandomNames()
	fmt.Println("the size of the ordered list is:", len(sortedNames))

	targetName := "Davi Costa"

	startTime := time.Now()

	binarySearcher := BinarySearch{}
	index := binarySearcher.Search(targetName, sortedNames)

	endTime := time.Now()

	executionTimeInNanosecounds := endTime.Sub(startTime)
	executionTimeInMilliseconds := float64(executionTimeInNanosecounds) / float64(time.Millisecond)

	if index != -1 {
		fmt.Printf("Execution time: %.3f milliseconds\n", executionTimeInMilliseconds)
	} else {
		fmt.Println("Name not found in the list:")
		fmt.Printf("Execution time: %.3f milliseconds\n", executionTimeInMilliseconds)
	}
}

func generateRandomNames() []Name {

	fullNames := "Arthur Silva, Bernardo Santos, Davi Costa, Enzo Oliveira, Gabriel Souza, Heitor Rodrigues, Lorenzo Nunes, Miguel Ferreira, Pedro Pereira, Rafael Oliveira, Lucas Gomes, Matheus Silva, Felipe Souza, Guilherme Oliveira, Arthur Mendes, Gabriel Almeida, Bernardo Soares, Davi Cardoso, Enzo Ferreira, Heitor Martins, Lorenzo Teixeira, Miguel Costa, Pedro Santos, Rafael Pereira, Lucas Souza, Matheus Oliveira, Felipe Gomes, Guilherme Silva, Arthur Santos, Bernardo Costa, Davi Almeida, Enzo Mendes, Heitor Cardoso, Lorenzo Ferreira, Miguel Teixeira, Pedro Soares, Rafael Martins, Lucas Costa, Matheus Santos, Felipe Oliveira, Guilherme Gomes, Arthur Almeida, Bernardo Soares, Davi Ferreira, Enzo Martins, Heitor Teixeira, Lorenzo Cardoso, Miguel Costa, Pedro Santos, Rafael Oliveira, Alice Silva, Beatriz Santos, Clara Costa, Eduarda Oliveira, Helena Souza, sabella Rodrigues, Júlia Nunes, Laura Ferreira, Luísa Oliveira, Manuela Pereira, Maria Eduarda Gomes, Mariana Silva, Melissa Souza, Nicole Oliveira, Sophia Mendes, Valentina Almeida, Alice Soares, Beatriz Cardoso, Clara Ferreira, Eduarda Martins, Helena Teixeira, Isabella Costa, Júlia Santos, Laura Pereira, Luísa Oliveira, Manuela Gomes, Maria Eduarda Silva, Mariana Souza, Melissa Oliveira, Nicole Mendes, Sophia Almeida, Valentina Soares, Alice Ferreira, Beatriz Martins, Helena Teixeira, Isabella Cardoso, Júlia Costa, Laura Santos, Luísa Oliveira, Manuela Gomes, Maria Eduarda Silva, Mariana Souza, Melissa Oliveira, Nicole Mendes, Sophia Almeida, Valentina Soares, Alice Ferreira, Beatriz Martins, Helena Teixeira, Isabella Cardoso"

	splitedNames := strings.Split(fullNames, ",")

	names := []Name{}
	for _, value := range splitedNames {
		trimmedNames := strings.TrimLeft(value, " ")
		names = append(names, Name{FullName: trimmedNames})
	}

	return names
}

func (b BinarySearch) Search(target string, data []Name) int {
	start := 0
	end := len(data) - 1
	middle := (end + start) / 2
	operation := 0

	for start <= end {
		if data[middle].FullName == target {
			fmt.Println("Name found in the list:", data[middle].FullName)
			return middle
		} else if data[middle].FullName > target {
			end = middle - 1
		} else {
			start = middle + 1
		}
		middle = (end + start) / 2
		operation++
		fmt.Println("Number of binary search operations performed:", operation)
	}
	return -1
}
