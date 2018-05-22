package recommender

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/cdipaolo/goml/text"
	"github.com/gaspiman/cosine_similarity"
	"github.com/jdkato/prose/tokenize"
)

/*
	https://en.wikipedia.org/wiki/Bag-of-words_model
	https://nlp.stanford.edu/IR-book/html/htmledition/stemming-and-lemmatization-1.html
	https://machinelearningmastery.com/gentle-introduction-bag-words-model/
	https://www.linkedin.com/pulse/content-based-recommender-engine-under-hood-venkat-raman/
*/

/*
	Flow of reccomendation:

	1. Get words, ngrams
	2. Get coisne similarity
	3.
*/

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
	Stemming      bool
	NgramsSizes   []int
	tokenizer     tokenize.ProseTokenizer
}

/*
	Setup of the recommender
*/

func New(tokenizerType TokenizerType, stemming bool, ngramsSizes []int) Recommender {
	r := Recommender{
		TokenizerType: tokenizerType,
		Stemming:      stemming,
		NgramsSizes:   ngramsSizes,
	}

	return r
}

/*
	Section deals with the frequency of the words
*/

func (r *Recommender) Frequency(words []string) []float64 {
	var freqArray []float64
	freq := text.TermFrequencies(words)
	for _, v := range freq {
		freqArray = append(freqArray, v.Frequency)
	}

	return freqArray
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

func (r *Recommender) CosineSimilarity(a, b []float64) {
	cosine, _ := cosine_similarity.Cosine(a, b)
	fmt.Println(cosine)
}
