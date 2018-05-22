package recommender

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetContent(t *testing.T) {
	tfidf := NewTFIDF()
	//bow := NewBOW(false, []int{2, 3})

	_, err := tfidf.getContent(posts)
	assert.NoError(t, err)

	//fmt.Println(result)
}

func TestAnalyze(t *testing.T) {
	tfidf := NewTFIDF()
	bow := NewBOW(false, []int{2, 3})

	words, err := bow.Get(posts)
	assert.NoError(t, err)

	similarities, err := tfidf.Amalyze(words, posts)
	assert.NoError(t, err)

	for _, similarity := range similarities {
		fmt.Println(similarity)
	}
}
