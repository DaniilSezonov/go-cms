package data

import "testing"

func TestValidateString(t *testing.T) {
	contentType := ContentType{
		Name:      "TestName",
		ValueType: String,
	}
	value := "Test String"
	ok := contentType.Validate(value)
	if !ok {
		t.Fail()
	}
}

func TestValidateList(t *testing.T) {
	contentType := ContentType{
		Name:      "List of strings",
		ValueType: List,
		Children: []ContentType{
			{ValueType: String},
		},
	}

	value := []string{
		"Apple",
		"Orange",
		"Tomato",
	}

	ok := contentType.Validate(value)
	if !ok {
		t.Fail()
	}
}

func TestValidateListFail(t *testing.T) {
	contentType := ContentType{
		Name:      "List of strings",
		ValueType: List,
		Children: []ContentType{
			{ValueType: String},
		},
	}

	value := []interface{}{
		"Apple",
		"Orange",
		123.4,
	}

	ok := contentType.Validate(value)
	if ok {
		t.Fail()
	}
}

func TestValidateObject(t *testing.T) {
	contentType := ContentType{
		Name:      "Product",
		ValueType: Object,
		Children: []ContentType{
			{Name: "Name", ValueType: String},
			{Name: "Description", ValueType: RichText},
			{Name: "Price", ValueType: Decimal},
			{Name: "Visible", ValueType: Boolean},
			{Name: "Tags", ValueType: List, Children: []ContentType{
				{ValueType: String},
			}},
		},
	}

	value := map[string]interface{}{
		"Name":        "Apple",
		"Description": "Juicy. Tasty. Yummy.",
		"Price":       2.5,
		"Visible":     true,
		"Tags":        []string{"fruit", "healthy"},
	}

	ok := contentType.Validate(value)
	if !ok {
		t.Fail()
	}
}

func TestValidateObjectMissingField(t *testing.T) {
	contentType := ContentType{
		Name:      "Product",
		ValueType: Object,
		Children: []ContentType{
			{Name: "Name", ValueType: String},
			{Name: "Description", ValueType: RichText},
			{Name: "Price", ValueType: Decimal},
			{Name: "Visible", ValueType: Boolean},
			{Name: "Tags", ValueType: List, Children: []ContentType{
				{ValueType: String},
			}},
		},
	}

	value := map[string]interface{}{
		"Name": "Apple",
	}

	ok := contentType.Validate(value)
	if ok {
		t.Fail()
	}
}

func TestValidateObjectInvalidFieldType(t *testing.T) {
	contentType := ContentType{
		Name:      "Product",
		ValueType: Object,
		Children: []ContentType{
			{Name: "Name", ValueType: String},
			{Name: "Description", ValueType: RichText},
			{Name: "Price", ValueType: Decimal},
			{Name: "Visible", ValueType: Boolean},
			{Name: "Tags", ValueType: List, Children: []ContentType{
				{ValueType: String},
			}},
		},
	}

	value := map[string]interface{}{
		"Name":        "Apple",
		"Description": "Juicy. Tasty. Yummy.",
		"Price":       "2.5",
		"Visible":     true,
		"Tags":        []string{"fruit", "healthy"},
	}

	ok := contentType.Validate(value)
	if ok {
		t.Fail()
	}
}
