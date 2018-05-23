package recommender

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNgrams(t *testing.T) {
	bow := NewBOW(false, []int{12})

	_, err := bow.ngrams(posts)
	assert.NoError(t, err)

	/*for _, ngram := range resultSet {
		for key, ngramInner := range ngram {
			fmt.Printf("%s -> %d\n", key, ngramInner)
		}
	}*/
}

func TestGetWords(t *testing.T) {
	bow := NewBOW(false, []int{1})

	words, err := bow.getWords(posts)
	assert.NoError(t, err)
	assert.Equal(t, 7567, len(words))

	words = bow.getUniqueWords(words)
	assert.Equal(t, 2253, len(words))
}

func TestGetBOW(t *testing.T) {
	bow := NewBOW(false, []int{2, 3, 4})

	words, err := bow.Get(posts)
	assert.NoError(t, err)
	assert.Equal(t, 22984, len(words))

}
