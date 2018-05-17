package recommender

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gaspiman/cosine_similarity"
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
	TokenizerType TokenizerType
	tokenizer     tokenize.ProseTokenizer
}

/*
	Setup of the recommender
*/

func New(tokenizerType TokenizerType) Recommender {
	r := Recommender{
		TokenizerType: tokenizerType,
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

/*
	Section deals with the words, getting the actual list of unique words
*/

func (r *Recommender) Vocabulary(unprocessedContent interface{}) ([]string, error) {
	words, err := r.getWords(unprocessedContent)
	if err != nil {
		return nil, err
	}

	return r.getUniqueWords(words), nil
}

func (r *Recommender) getWords(unprocessedContent interface{}) ([]string, error) {
	//var documentVectors []string

	rv := reflect.ValueOf(unprocessedContent)
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
			}
		}
	}

	return r.tokenize(wholeContent), nil
}

func (r *Recommender) getUniqueWords(words []string) []string {
	var wordsUniq []string
	var wordsMap = make(map[string]bool)

	for _, word := range words {
		wordsMap[word] = true
	}
	for k := range wordsMap {
		wordsUniq = append(wordsUniq, k)
	}
	return wordsUniq
}

func (r *Recommender) tokenize(data string) []string {
	return r.tokenizer.Tokenize(data)
}

/*
	Section deals with the binary vector representation of the given post
*/

func (r *Recommender) BinaryRepresentation(unprocessedContent interface{}, vocabulary []string) (map[string][]int8, error) {
	var representations = make(map[string][]int8)

	rv := reflect.ValueOf(unprocessedContent)
	rt := rv.Type()

	if rt.Kind() != reflect.Slice {
		return nil, errors.New("must be slice")
	}

	for i := 0; i < rv.Len(); i++ {
		rvInner := rv.Index(i)
		if rvInner.Kind() != reflect.Struct {
			return nil, errors.New("slice elements must be structs")
		}

		rvInnerType := rvInner.Type()
		var id string

	Loop:
		for j := 0; j < rvInnerType.NumField(); j++ {
			field := rvInner.Field(j)
			fieldType := rvInnerType.Field(j)

			indexable, ok := fieldType.Tag.Lookup("indexable")
			if ok {
				if indexable == "id" {
					if field.Kind() == reflect.String {
						id = field.String()
						break Loop
					}
				}
			}
		}

		representation, err := r.getBinaryRepresentation(rvInner.Interface(), vocabulary)
		if err != nil {
			return nil, err
		}

		representations[id] = representation
	}

	return representations, nil
}

func (r *Recommender) getBinaryRepresentation(unprocessedContent interface{}, vocabulary []string) ([]int8, error) {
	var binaryRepresentation []int8

	rv := reflect.ValueOf(unprocessedContent)
	rt := rv.Type()

	if rv.Kind() != reflect.Struct {
		return nil, errors.New("must be struct")
	}

	var content string
	for i := 0; i < rt.NumField(); i++ {
		if tag, ok := rt.Field(i).Tag.Lookup("indexable"); ok && tag == "content" {
			if rv.Field(i).Kind() != reflect.String {
				continue
			}
			content += rv.Field(i).String() + " "
		}
	}

	for _, word := range vocabulary {
		if strings.Contains(content, word) {
			binaryRepresentation = append(binaryRepresentation, 1)
		} else {
			binaryRepresentation = append(binaryRepresentation, 0)
		}
	}

	return binaryRepresentation, nil
}

func (r *Recommender) CosineSimilarity() {
	cosine_similarity.Cosine(nil, nil)
}
