package main

import "fmt"

type WordVerifier interface {
	IsValid(word string) bool
}

type wordMapVerifier struct {
	wordMap map[string]bool
}

func (w *wordMapVerifier) IsValid(word string) bool {
	_, ok := w.wordMap[word]
	return ok
}

func createWordVerifier(wordList []string) WordVerifier {
	wordMap := make(map[string]bool)
	for _, word := range wordList {
		wordMap[word] = true
	}
	return &wordMapVerifier{wordMap: wordMap}
}

func buildNewWord(word string, i int, ch rune) string {
	return word[:i] + string(ch) + word[i+1:]
}

func bfs(startWord string, endWord string, verifier WordVerifier, wordList []string) int {
	if !isValidWordList(wordList) {
		return 0
	}

	queue := []string{startWord}
	steps := 0
	for len(queue) > 0 {
		for _, word := range queue {
			if word == endWord {
				return steps + 1
			}
			for i := 0; i < len(word); i++ {
				for ch := 'a'; ch <= 'z'; ch++ {
					newWord := buildNewWord(word, i, ch)
					if verifier.IsValid(newWord) {
						queue = append(queue, newWord)
					}
				}
			}
		}
		steps++
		queue = queue[1:]
	}
	return 0
}

func isValidWordList(wordList []string) bool {
	return len(wordList) > 0
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	verifier := createWordVerifier(wordList)
	return bfs(beginWord, endWord, verifier, wordList)
}

func main() {
	beginWord := "hit"
	endWord := "cog"
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}

	minLength := ladderLength(beginWord, endWord, wordList)

	if minLength > 0 {
		fmt.Printf("minimum word ladder size: %d\n", minLength)
	} else {
		fmt.Println("It is not possible to transform the initial word into the final one using the list provided.")
	}
}
