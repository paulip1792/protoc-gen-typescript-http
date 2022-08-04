package plugin

import "google.golang.org/protobuf/reflect/protoreflect"

// NamingOptions is a configurable Generatior optiolns.
type GeneratorOptions struct {
	// UseProtoNames uses proto field name instead of lowerCamelCase name in JSON
	// field names.
	UseProtoNames bool
}

var (
	DefaultGeneratorOptions = GeneratorOptions{
		UseProtoNames: false,
	}
)

var GetFieldName = func(fd protoreflect.FieldDescriptor) string {
	if DefaultGeneratorOptions.UseProtoNames {
		return fd.TextName()
	}
	return fd.JSONName()
}
