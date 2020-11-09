package data

import (
	"errors"
	"fmt"
)

type Content struct {
	ID     int         `json:"id,omitempty"`
	TypeID int         `json:"typeID"`
	Value  interface{} `json:"value,omitempty"`
}

var contents = map[int]Content{
	1: {
		ID: 1, TypeID: 1, Value: product{
			Name:        "Apple",
			Description: "Juicy. Tasty. Yummy.",
			Price:       2.5,
			Visible:     true,
			Tags:        []string{"fruit", "healthy"},
		},
	},
}
var contentCounter = len(contents) + 1

func GetContentByTypeID(typeID int) chan Content {
	result := make(chan Content)

	// Just to test how channels work in Go
	// Can be replaced with an array of Content.
	go func() {
		for _, content := range contents {
			if content.TypeID == typeID {
				result <- content
			}
		}
		close(result)
	}()

	return result
}

func GetContentByID(id int) (Content, bool) {
	content, exists := contents[id]
	return content, exists
}

func CreateContent(content *Content) error {
	contentType, found := GetContentTypeByID(content.TypeID)
	if !found {
		return errors.New(fmt.Sprintf("Content type %d was not found.", content.TypeID))
	}

	ok := contentType.Validate(content.Value)
	if !ok {
		return errors.New("Invalid content.")
	}

	content.ID = contentCounter
	contentCounter++
	contents[content.ID] = *content
	return nil
}

func UpdateContent(id int, content Content) (bool, error) {
	contentType, found := GetContentTypeByID(content.TypeID)
	if !found {
		return false, errors.New(fmt.Sprintf("Content type %d was not found.", content.TypeID))
	}

	ok := contentType.Validate(content.Value)
	if !ok {
		return false, errors.New("Invalid content.")
	}

	content.ID = id
	contents[content.ID] = content
	return true, nil
}

func DeleteContent(id int) {
	delete(contents, id)
}
