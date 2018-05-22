package recommender

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strings"

	"github.com/caneroj1/stemmer"
	"github.com/jdkato/prose/tokenize"
)

type BOW struct {
	Stemming    bool
	NgramsSizes []int
	tokenizer   tokenize.ProseTokenizer
}

func NewBOW(stemming bool, ngramsSizes []int) BOW {
	b := BOW{
		Stemming:    stemming,
		NgramsSizes: ngramsSizes,
		tokenizer:   tokenize.NewWordBoundaryTokenizer(),
	}

	return b
}

func (b *BOW) Get(unprocessedContent interface{}) ([]string, error) {
	words, err := b.getWords(unprocessedContent)
	if err != nil {
		return nil, err
	}

	if b.Stemming {
		words = b.getUniqueWords(b.stem(words))
	}

	singleWords := b.getUniqueWords(words)

	ngrams, err := b.ngrams(unprocessedContent)
	if err != nil {
		return nil, err
	}
	ngramsSlice := b.ngramsToSlice(ngrams)

	return append(singleWords, ngramsSlice...), nil
}

func (b *BOW) getContent(unprocessedContent interface{}) (string, error) {
	rv := reflect.ValueOf(unprocessedContent)
	rt := rv.Type()

	if rt.Kind() != reflect.Slice {
		return "", errors.New("must be slice")
	}

	var wholeContent string
	for i := 0; i < rv.Len(); i++ {
		rvInner := rv.Index(i)
		if rvInner.Kind() != reflect.Struct {
			return "", errors.New("slice elements must be structs")
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

	return wholeContent, nil
}

// stem get the vocabulary and get the exact words like the pure meaning
// ex.: doing -> do
func (b *BOW) stem(words []string) []string {
	fmt.Println(stemmer.StemMultiple(words))
	return stemmer.StemMultiple(words)
}

func (b *BOW) getWords(unprocessedContent interface{}) ([]string, error) {
	content, err := b.getContent(unprocessedContent)
	if err != nil {
		return nil, err
	}

	return b.tokenize(content), nil
}

func (b *BOW) getUniqueWords(words []string) []string {
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

func (b *BOW) tokenize(data string) []string {
	return b.tokenizer.Tokenize(data)
}

func (b *BOW) tokenizeTreebankWord(data string) []string {
	t := tokenize.NewTreebankWordTokenizer()
	return t.Tokenize(data)
}

func (b *BOW) ngramsToSlice(ngrams []map[string]uint32) []string {
	var result []string

	for _, ngram := range ngrams {
		for words := range ngram {
			result = append(result, words)
		}
	}
	return result
}

func (b *BOW) ngrams(unprocessedContent interface{}) ([]map[string]uint32, error) {
	content, err := b.getContent(unprocessedContent)
	if err != nil {
		return nil, err
	}

	words := b.tokenize(content)

	var ngrams []map[string]uint32
	for _, size := range b.NgramsSizes {
		ngrams = append(ngrams, b.ngramsBySize(words, size))
	}

	return ngrams, nil
}

func (b *BOW) ngramsBySize(words []string, size int) map[string]uint32 {
	count := make(map[string]uint32, 0)
	offset := int(math.Floor(float64(size / 2)))

	max := len(words)
	for i := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		count[gram]++
	}

	return count
}
