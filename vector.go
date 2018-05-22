package recommender

import (
	"errors"
	"reflect"
	"strings"
)

type Vector struct {
}

func NewVector() Vector {
	return Vector{}
}

func (v *Vector) BinaryRepresentation(unprocessedContent interface{}, vocabulary []string) (map[string][]int8, error) {
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

		representation, err := v.getBinaryRepresentation(rvInner.Interface(), vocabulary)
		if err != nil {
			return nil, err
		}

		representations[id] = representation
	}

	return representations, nil
}

func (v *Vector) getBinaryRepresentation(unprocessedContent interface{}, vocabulary []string) ([]int8, error) {
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
