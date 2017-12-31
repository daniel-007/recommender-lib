package recommender

import (
	"fmt"
	"strings"
)

type Post struct {
	Title string
	Body  string
}

type replace struct {
	new string
	old string
}

func Train(posts []Post) {
	var wordsContainer [][]string
	for _, post := range posts {
		wordsContainer = append(wordsContainer, getWordsOfPost(post.Body))
	}
	words := getAllWords(wordsContainer)

	fmt.Println(len(words))

}

func getWordsOfPost(body string) []string {
	var wordsUniq []string
	var wordsMap = make(map[string]bool)
	body = replacer(strings.ToLower(body))
	words := strings.Split(body, " ")

	for _, word := range words {
		wordsMap[word] = true
	}
	for k := range wordsMap {
		wordsUniq = append(wordsUniq, k)
	}
	return wordsUniq
}

func getAllWords(container [][]string) []string {
	var wordsUniq []string
	var wordsMap = make(map[string]bool)

	for _, words := range container {
		for _, word := range words {
			wordsMap[word] = true
		}
	}
	for k := range wordsMap {
		wordsUniq = append(wordsUniq, k)
	}
	return wordsUniq
}

func replacer(body string) string {
	var replaces = []replace{
		replace{
			new: "",
			old: ",",
		},
		replace{
			new: "",
			old: ".",
		},
		replace{
			old: "“",
			new: "",
		},
		replace{
			old: ")",
			new: "",
		},
		replace{
			old: "(",
			new: "",
		},
		replace{
			old: "!",
			new: "",
		},
		replace{
			old: " — ",
			new: " ",
		},
		replace{
			old: "”",
			new: "",
		},
		replace{
			old: "?",
			new: "",
		},
		replace{
			old: ":",
			new: "",
		},
	}

	for _, r := range replaces {
		body = strings.Replace(body, r.old, r.new, -1)
	}
	return body
}
