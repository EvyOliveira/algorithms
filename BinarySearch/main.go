package main

import (
	"fmt"
	"sort"
	"strings"
)

func generateRandomNames() []string {
	fullNames := "Arthur Silva, Bernardo Santos, Davi Costa, Enzo Oliveira, Gabriel Souza, Heitor Rodrigues, Lorenzo Nunes, Miguel Ferreira, Pedro Pereira, Rafael Oliveira, Lucas Gomes, Matheus Silva, Felipe Souza, Guilherme Oliveira, Arthur Mendes, Gabriel Almeida, Bernardo Soares, Davi Cardoso, Enzo Ferreira, Heitor Martins, Lorenzo Teixeira, Miguel Costa, Pedro Santos, Rafael Pereira, Lucas Souza, Matheus Oliveira, Felipe Gomes, Guilherme Silva, Arthur Santos, Bernardo Costa, Davi Almeida, Enzo Mendes, Heitor Cardoso, Lorenzo Ferreira, Miguel Teixeira, Pedro Soares, Rafael Martins, Lucas Costa, Matheus Santos, Felipe Oliveira, Guilherme Gomes, Arthur Almeida, Bernardo Soares, Davi Ferreira, Enzo Martins, Heitor Teixeira, Lorenzo Cardoso, Miguel Costa, Pedro Santos, Rafael Oliveira, Alice Silva, Beatriz Santos, Clara Costa, Eduarda Oliveira, Helena Souza, sabella Rodrigues, Júlia Nunes, Laura Ferreira, Luísa Oliveira, Manuela Pereira, Maria Eduarda Gomes, Mariana Silva, Melissa Souza, Nicole Oliveira, Sophia Mendes, Valentina Almeida, Alice Soares, Beatriz Cardoso, Clara Ferreira, Eduarda Martins, Helena Teixeira, Isabella Costa, Júlia Santos, Laura Pereira, Luísa Oliveira, Manuela Gomes, Maria Eduarda Silva, Mariana Souza, Melissa Oliveira, Nicole Mendes, Sophia Almeida, Valentina Soares, Alice Ferreira, Beatriz Martins, Helena Teixeira, Isabella Cardoso, Júlia Costa, Laura Santos, Luísa Oliveira, Manuela Gomes, Maria Eduarda Silva, Mariana Souza, Melissa Oliveira, Nicole Mendes, Sophia Almeida, Valentina Soares, Alice Ferreira, Beatriz Martins, Helena Teixeira, Isabella Cardoso"
	splitedNames := strings.Split(fullNames, ",")

	names := []string{}

	for _, value := range splitedNames {
		names = append(names, strings.TrimLeft(value, " "))
	}

	sort.Strings(names)
	return names
}

func main() {
	nameList := generateRandomNames()
	fmt.Println(len(nameList))

	nameSearched := "Lorenzo Teixeira"
	start := 0
	end := len(nameList) - 1
	operation := 0

	for start <= end {
		middle := (end + start) / 2
		currentName := nameList[middle]

		if currentName == nameSearched {
			fmt.Println("Name found in the list:", currentName)
			break
		}

		if currentName > nameSearched {
			end = middle - 1
		}

		if currentName < nameSearched {
			start = middle + 1
		}
		operation++
	}
	fmt.Println("Number of binary search operations performed:", operation)
}
