package main

import (
	"fmt"
	"strings"
)

type Name struct {
	FullName string
}

const nameSearched = "Lorenzo Teixeira"

func parseNamesFromString(nameString string) []Name {
	splitedNames := strings.Split(nameString, ",")

	names := []Name{}
	for _, value := range splitedNames {
		trimmedNames := strings.TrimLeft(value, " ")
		names = append(names, Name{FullName: trimmedNames})
	}
	return names
}

func searchNames() {
	fullNames := "Arthur Silva, Bernardo Santos, Davi Costa, Enzo Oliveira, Gabriel Souza, Heitor Rodrigues, Lorenzo Nunes, Miguel Ferreira, Pedro Pereira, Rafael Oliveira, Lucas Gomes, Matheus Silva, Felipe Souza, Guilherme Oliveira, Arthur Mendes, Gabriel Almeida, Bernardo Soares, Davi Cardoso, Enzo Ferreira, Heitor Martins, Lorenzo Teixeira, Miguel Costa, Pedro Santos, Rafael Pereira, Lucas Souza, Matheus Oliveira, Felipe Gomes, Guilherme Silva, Arthur Santos, Bernardo Costa, Davi Almeida, Enzo Mendes, Heitor Cardoso, Lorenzo Ferreira, Miguel Teixeira, Pedro Soares, Rafael Martins, Lucas Costa, Matheus Santos, Felipe Oliveira, Guilherme Gomes, Arthur Almeida, Bernardo Soares, Davi Ferreira, Enzo Martins, Heitor Teixeira, Lorenzo Cardoso, Miguel Costa, Pedro Santos, Rafael Oliveira, Alice Silva, Beatriz Santos, Clara Costa, Eduarda Oliveira, Helena Souza, sabella Rodrigues, Júlia Nunes, Laura Ferreira, Luísa Oliveira, Manuela Pereira, Maria Eduarda Gomes, Mariana Silva, Melissa Souza, Nicole Oliveira, Sophia Mendes, Valentina Almeida, Alice Soares, Beatriz Cardoso, Clara Ferreira, Eduarda Martins, Helena Teixeira, Isabella Costa, Júlia Santos, Laura Pereira, Luísa Oliveira, Manuela Gomes, Maria Eduarda Silva, Mariana Souza, Melissa Oliveira, Nicole Mendes, Sophia Almeida, Valentina Soares, Alice Ferreira, Beatriz Martins, Helena Teixeira, Isabella Cardoso, Júlia Costa, Laura Santos, Luísa Oliveira, Manuela Gomes, Maria Eduarda Silva, Mariana Souza, Melissa Oliveira, Nicole Mendes, Sophia Almeida, Valentina Soares, Alice Ferreira, Beatriz Martins, Helena Teixeira, Isabella Cardoso"
	nameList := parseNamesFromString(fullNames)
	fmt.Println(len(nameList))

	if len(nameSearched) == 0 {
		fmt.Println("empty search name, enter a valid search name.")
		return
	}

	index := linearSearch(nameList, nameSearched)

	if index != -1 {
		fmt.Println("name found in the list:", nameList[index].FullName)
	} else {
		fmt.Println("name not found in the list:")
	}
}

func linearSearch(nameList []Name, name string) int {
	for index, currentName := range nameList {
		if currentName.FullName == name {
			fmt.Println("name found in the list:", currentName.FullName)
			return index
		}
	}
	fmt.Println("name not found in the list:")
	return -1
}

func main() {
	var nameList []Name
	var name string

	searchNames()
	linearSearch(nameList, name)
}
