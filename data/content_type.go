package data

type ContentType struct {
	ID        int
	Name      string
	ValueType string
	Children  []ContentType
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
	File     = "file"
)

type ContentTypeList []ContentType

var data = map[int]ContentType{
	0: {
		Name: "Product", ValueType: Object, Children: []ContentType{
			{Name: "Name", ValueType: String},
			{Name: "Description", ValueType: RichText},
			{Name: "Price", ValueType: Decimal},
			{Name: "Logo", ValueType: File},
		},
	},
}

var counter = 1

func GetContentTypes() ContentTypeList {
	result := make(ContentTypeList, len(data))

	i := 0
	for _, contentType := range data {
		result[i] = contentType
		i++
	}
	return result
}

func GetContentTypeByID(id int) (ContentType, bool) {
	contentType, exists := data[id]
	return contentType, exists
}

func CreateContentType(ct *ContentType) {
	ct.ID = counter
	counter++

	data[ct.ID] = *ct
}

func UpdateContentType(id int, ct ContentType) bool {
	_, update := data[id]
	ct.ID = id
	data[ct.ID] = ct
	return update
}

func DeleteContentType(id int) {
	delete(data, id)
}
