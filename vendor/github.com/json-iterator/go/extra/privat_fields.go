package extra

import (
	"unicode"

	"github.com/json-iterator/go"
)

// SupportPrivateFields include private fields when encoding/decoding
func SupportPrivateFields() {
	jsoniter.RegisterExtension(&privateFieldsExtension{})
}

type privateFieldsExtension struct {
	jsoniter.DummyExtension
}

func (extension *privateFieldsExtension) UpdateStructDescriptor(structDescriptor *jsoniter.StructDescriptor) {
	for _, binding := range structDescriptor.Fields {
		isPrivate := unicode.IsLower(rune(binding.Field.Name[0]))
		if isPrivate {
			binding.FromNames = []string{binding.Field.Name}
			binding.ToNames = []string{binding.Field.Name}
		}
	}
}
