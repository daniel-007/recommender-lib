package recommender

import (
	"errors"
	"reflect"
	"strings"

	"github.com/jdkato/prose/tokenize"
)

type TokenizerType int8

const (
	Blankline TokenizerType = iota
	PunkSentence
	TreebankWord
	WordBoundary
	WordPunct
)

type Recommender struct {
	TokenizerType      TokenizerType
	tokenizer          tokenize.ProseTokenizer
	unprocessedContent interface{}
}

func New(tokenizerType TokenizerType, unprocessedContent interface{}) Recommender {
	r := Recommender{
		TokenizerType:      tokenizerType,
		unprocessedContent: unprocessedContent,
	}

	r.getTokenizer()
	return r
}

func (r *Recommender) getTokenizer() {
	switch r.TokenizerType {
	case Blankline:
		r.tokenizer = tokenize.NewBlanklineTokenizer()
	case PunkSentence:
		r.tokenizer = tokenize.NewPunktSentenceTokenizer()
	case TreebankWord:
		r.tokenizer = tokenize.NewTreebankWordTokenizer()
	case WordBoundary:
		r.tokenizer = tokenize.NewWordBoundaryTokenizer()
	case WordPunct:
		r.tokenizer = tokenize.NewWordPunctTokenizer()
	}
}

func (r *Recommender) getWords() ([]string, error) {
	//var documentVectors []string

	rv := reflect.ValueOf(r.unprocessedContent)
	rt := rv.Type()

	if rt.Kind() != reflect.Slice {
		return nil, errors.New("must be slice")
	}

	var wholeContent string
	for i := 0; i < rv.Len(); i++ {
		rvInner := rv.Index(i)
		if rvInner.Kind() != reflect.Struct {
			return nil, errors.New("slice elements must be structs")
		}

		rvInnerType := rvInner.Type()

		for j := 0; j < rvInnerType.NumField(); j++ {
			field := rvInner.Field(j)
			fieldType := rvInnerType.Field(j)

			indexable, ok := fieldType.Tag.Lookup("indexable")
			if ok {
				if indexable == "content" {
					if field.Kind() == reflect.String {
						wholeContent += field.String()
					}
				}
				//fmt.Println(indexable)
			}

			//fmt.Println(field)
		}
	}

	return r.tokenize(wholeContent), nil
}

func (r *Recommender) tokenize(data string) []string {
	return r.tokenizer.Tokenize(data)
}

/*
	OLD
*/

type Post struct {
	Title string
	Body  string
}

type replace struct {
	new string
	old string
}

func Recommend(data map[string]map[string]bool) []string {
	return []string{}
}

func Train(posts []Post) map[string]map[string]bool {
	var wordsContainer [][]string
	for _, post := range posts {
		wordsContainer = append(wordsContainer, getWordsOfPost(post.Body))
	}
	words := getAllWords(wordsContainer)

	return mapPostsByWords(posts, words)
}

/*
	Recommend related
*/

/*
	Train related
*/

func tokenizer(data string) []string {
	//tokenizer := tokenize.NewTreebankWordTokenizer()
	tokenizer2 := tokenize.NewTreebankWordTokenizer()

	// @TODO setup tokenizer for different options

	//return tokenizer.Tokenize(data)
	return tokenizer2.Tokenize(data)
}

func indexing(dataSet []string) map[int]string {
	var indexedRepresentation map[int]string

	return indexedRepresentation
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
