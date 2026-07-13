package altStruct

import (
	"reflect"
	"strings"
	"testing"
)

func TestVertoProfileParameterUniqueKeyIncludesSecure(t *testing.T) {
	typ := reflect.TypeOf(ConfigVertoProfileParameter{})

	for _, fieldName := range []string{"Name", "Secure", "Parent"} {
		field, ok := typ.FieldByName(fieldName)
		if !ok {
			t.Fatalf("ConfigVertoProfileParameter.%s field is missing", fieldName)
		}
		if tag := field.Tag.Get("customsql"); !strings.Contains(tag, "unique_1") {
			t.Fatalf("ConfigVertoProfileParameter.%s customsql tag %q does not include unique_1", fieldName, tag)
		}
	}
}
