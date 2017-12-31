package recommender

import (
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

func Train(posts []Post) map[string]map[string]bool {
	var wordsContainer [][]string
	for _, post := range posts {
		wordsContainer = append(wordsContainer, getWordsOfPost(post.Body))
	}
	words := getAllWords(wordsContainer)

	return mapPostsByWords(posts, words)
}

func mapPostsByWords(posts []Post, words []string) map[string]map[string]bool {
	var trained = make(map[string]map[string]bool)
	for _, post := range posts {
		wordsOfPost := getWordsOfPost(post.Body)
		trained[post.Title] = make(map[string]bool)
		for _, word := range words {
			trained[post.Title][word] = contains(wordsOfPost, word)
		}
	}
	return trained
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

func contains(list []string, word string) bool {
	for _, w := range list {
		if w == word {
			return true
		}
	}
	return false
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
