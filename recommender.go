package recommender

import (
	"fmt"

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

func (r *Recommender) CosineSimilarity(a, b []float64) {
	cosine, _ := cosine_similarity.Cosine(a, b)
	fmt.Println(cosine)
}
