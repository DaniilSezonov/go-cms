package data

import (
	"reflect"
)

type ContentType struct {
	ID        int           `json:"id,omitempty"`
	Name      string        `json:"name,omitempty"`
	ValueType string        `json:"valueType"`
	Children  []ContentType `json:"children,omitempty"`
}

const (
	String   = "string"
	Int      = "int"
	Decimal  = "decimal"
	Boolean  = "boolean"
	Text     = "text"
	RichText = "rich-text"
	Object   = "object"
	List     = "list"
	// File     = "file"
)

type ContentTypeList []ContentType

var contentTypes = map[int]ContentType{
	1: {
		ID: 1, Name: "Product", ValueType: Object, Children: []ContentType{
			{Name: "Name", ValueType: String},
			{Name: "Description", ValueType: RichText},
			{Name: "Price", ValueType: Decimal},
			{Name: "Visible", ValueType: Boolean},
			{Name: "Tags", ValueType: List, Children: []ContentType{{ValueType: String}}},
			// {Name: "Logo", ValueType: File},
		},
	},
}

var contentTypeCounter = len(contentTypes) + 1

func GetContentTypes() ContentTypeList {
	result := make(ContentTypeList, len(contentTypes))

	i := 0
	for _, contentType := range contentTypes {
		result[i] = contentType
		i++
	}
	return result
}

func GetContentTypeByID(id int) (ContentType, bool) {
	contentType, exists := contentTypes[id]
	return contentType, exists
}

func CreateContentType(ct *ContentType) {
	ct.ID = contentTypeCounter
	contentTypeCounter++

	contentTypes[ct.ID] = *ct
}

func UpdateContentType(id int, ct ContentType) bool {
	_, update := contentTypes[id]
	ct.ID = id
	contentTypes[ct.ID] = ct
	return update
}

func DeleteContentType(id int) {
	delete(contentTypes, id)
}

func (ct ContentType) Validate(value interface{}) bool {
	var ok bool

	switch ct.ValueType {
	case String:
		_, ok = value.(string)
	case Text:
		_, ok = value.(string)
	case RichText:
		_, ok = value.(string)
	case Int:
		_, ok = value.(int)
	case Boolean:
		_, ok = value.(bool)
	case Decimal:
		_, ok = value.(float64)
	// case File:
	// 	_, ok = value.(io.Reader)
	case List:
		kind := reflect.TypeOf(value).Kind()
		if kind != reflect.Slice {
			return false
		}

		children := reflect.ValueOf(value)

		if len(ct.Children) == 1 {
			childType := ct.Children[0]

			for i := 0; i < children.Len(); i++ {
				item := children.Index(i)
				itemValue := item.Interface()
				ok = childType.Validate(itemValue)
				if !ok {
					return false
				}
			}
		} else {
			ok = false
		}
	case Object:
		content := reflect.ValueOf(value)
		for _, childType := range ct.Children {
			field := content.MapIndex(reflect.ValueOf(childType.Name))
			if !field.IsValid() {
				return false
			}

			ok = childType.Validate(field.Interface())
			if !ok {
				return false
			}
		}
	default:
		break
	}
	return ok
}

// For holding test content values
type product struct {
	Name        string
	Description string
	Price       float64
	Visible     bool
	Tags        []string
}
