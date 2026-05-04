package mcpapi

import (
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	mcplib "github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Registrar func(*server.MCPServer)

func NewServer(name, version string, registrars ...Registrar) *server.MCPServer {
	mcpServer := server.NewMCPServer(
		name,
		version,
		server.WithToolCapabilities(false),
		server.WithRecovery(),
		server.WithInputSchemaValidation(),
	)

	for _, registrar := range registrars {
		registrar(mcpServer)
	}

	return mcpServer
}

func Mount(mux *http.ServeMux, endpointPath string, mcpServer *server.MCPServer) {
	mux.Handle(endpointPath, NewHTTPHandler(mcpServer))
}

func NewHTTPHandler(mcpServer *server.MCPServer) *server.StreamableHTTPServer {
	return server.NewStreamableHTTPServer(mcpServer)
}

func WithInputSchema[T any]() mcplib.ToolOption {
	return func(t *mcplib.Tool) {
		schema, err := jsonschema.For[T](&jsonschema.ForOptions{IgnoreInvalidTypes: true})
		if err != nil {
			return
		}

		applyEnumTags(schema, reflect.TypeFor[T]())

		rawSchema, err := json.Marshal(schema)
		if err != nil {
			return
		}

		t.InputSchema = mcplib.ToolInputSchema{}
		mcplib.WithRawInputSchema(rawSchema)(t)
	}
}

func applyEnumTags(schema *jsonschema.Schema, valueType reflect.Type) {
	valueType = dereferenceType(valueType)
	if schema == nil || valueType.Kind() != reflect.Struct {
		return
	}

	for i := range valueType.NumField() {
		field := valueType.Field(i)
		if !field.IsExported() {
			continue
		}

		propertyName, ok := jsonFieldName(field)
		if !ok {
			continue
		}

		propertySchema := schema.Properties[propertyName]
		if propertySchema == nil {
			continue
		}

		if tagValue := field.Tag.Get("enum"); tagValue != "" {
			propertySchema.Enum = parseEnumTag(tagValue, field.Type)
		}

		applyEnumTags(propertySchema, field.Type)
		itemType := dereferenceType(field.Type)
		if propertySchema.Items != nil && (itemType.Kind() == reflect.Array || itemType.Kind() == reflect.Slice) {
			applyEnumTags(propertySchema.Items, itemType.Elem())
		}
	}
}

func jsonFieldName(field reflect.StructField) (string, bool) {
	tag := field.Tag.Get("json")
	if tag == "-" {
		return "", false
	}

	name, _, _ := strings.Cut(tag, ",")
	if name != "" {
		return name, true
	}

	return field.Name, true
}

func parseEnumTag(tagValue string, valueType reflect.Type) []any {
	values := strings.Split(tagValue, ",")
	enum := make([]any, 0, len(values))
	valueType = dereferenceType(valueType)

	for _, value := range values {
		value = strings.TrimSpace(value)
		enum = append(enum, parseEnumValue(value, valueType.Kind()))
	}

	return enum
}

func parseEnumValue(value string, kind reflect.Kind) any {
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		parsed, err := strconv.ParseInt(value, 10, 64)
		if err == nil {
			return parsed
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		parsed, err := strconv.ParseUint(value, 10, 64)
		if err == nil {
			return parsed
		}
	case reflect.Float32, reflect.Float64:
		parsed, err := strconv.ParseFloat(value, 64)
		if err == nil {
			return parsed
		}
	case reflect.Bool:
		parsed, err := strconv.ParseBool(value)
		if err == nil {
			return parsed
		}
	}

	return value
}

func dereferenceType(valueType reflect.Type) reflect.Type {
	for valueType.Kind() == reflect.Pointer {
		valueType = valueType.Elem()
	}

	return valueType
}
