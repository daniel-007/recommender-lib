package recommender

import (
	"errors"
	"reflect"

	"github.com/wilcosheh/tfidf"
	"github.com/wilcosheh/tfidf/similarity"
)

type Similarity struct {
	Term1ID string
	Term2ID string
	Sim     float64
}

type TFIDF struct {
}

func NewTFIDF() TFIDF {
	return TFIDF{}
}

func (t *TFIDF) Amalyze(docs []string, unprocessedContent interface{}) ([]Similarity, error) {
	f := tfidf.New()
	f.AddDocs(docs...)

	dataContainer, err := t.getContent(unprocessedContent)
	if err != nil {
		return nil, err
	}

	var weights = make(map[string]map[string]float64)
	for id, data := range dataContainer {
		weight := f.Cal(data)
		weights[id] = weight
	}

	// @TODO do the exact checks only once
	var result []Similarity
	for id, weight := range weights {
		for id2, weight2 := range weights {
			if id != id2 {
				sim := similarity.Cosine(weight, weight2)
				simObj := Similarity{
					Term1ID: id,
					Term2ID: id2,
					Sim:     sim,
				}
				result = append(result, simObj)
			}
		}
	}

	return result, nil
}

func (t *TFIDF) getContent(unprocessedContent interface{}) (map[string]string, error) {
	var result = make(map[string]string)
	rv := reflect.ValueOf(unprocessedContent)
	rt := rv.Type()

	if rt.Kind() != reflect.Slice {
		return result, errors.New("must be slice")
	}

	for i := 0; i < rv.Len(); i++ {
		rvInner := rv.Index(i)
		if rvInner.Kind() != reflect.Struct {
			return result, errors.New("slice elements must be structs")
		}

		rvInnerType := rvInner.Type()

		var id string
		var content string
		for j := 0; j < rvInnerType.NumField(); j++ {
			field := rvInner.Field(j)
			fieldType := rvInnerType.Field(j)

			indexable, ok := fieldType.Tag.Lookup("indexable")
			if ok {
				if indexable == "id" {
					if field.Kind() == reflect.String {
						id = field.String()
					}
				}
				if indexable == "content" {
					if field.Kind() == reflect.String {
						content += field.String() + " "
					}
				}
			}

		}
		result[id] = content
	}

	return result, nil
}
