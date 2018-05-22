package recommender

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBinaryRepresentation(t *testing.T) {
	vector := NewVector()
	bow := NewBOW(false, []int{2, 3})

	vocabulary, err := bow.Get(posts)
	assert.NoError(t, err)

	_, err = vector.getBinaryRepresentation(posts[0], vocabulary)
	assert.NoError(t, err)

	//fmt.Println(binaryRepresentation)
}

func TestBinaryRepresentation(t *testing.T) {
	vector := NewVector()
	bow := NewBOW(false, []int{2, 3})

	vocabulary, err := bow.Get(posts)
	assert.NoError(t, err)

	_, err = vector.BinaryRepresentation(posts, vocabulary)
	assert.NoError(t, err)

	//fmt.Println(binaryRepresentation)
}
