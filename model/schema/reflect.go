package schema

import (
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/zlsgo/zdb/schema"
)

func FieldsFromStruct[T any]() Fields {
	var t T
	return fieldsFromType(reflect.TypeOf(t))
}

func fieldsFromType(t reflect.Type) Fields {
	if t == nil {
		return nil
	}
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil
	}

	fields := make(Fields)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if !sf.IsExported() {
			continue
		}

		if sf.Anonymous {
			embedded := fieldsFromType(sf.Type)
			for k, v := range embedded {
				if _, exists := fields[k]; !exists {
					fields[k] = v
				}
			}
			continue
		}

		fieldTag := sf.Tag.Get("field")
		if fieldTag == "-" {
			continue
		}

		name := getFieldName(sf)
		if name == "-" || name == "" {
			continue
		}

		ft := sf.Type
		nullable := false
		if ft.Kind() == reflect.Ptr {
			ft = ft.Elem()
			nullable = true
		}

		field := Field{
			Type:     goTypeToSchemaType(ft),
			Label:    name,
			Nullable: nullable,
		}

		parseFieldTag(fieldTag, &field)
		fields[name] = field
	}
	return fields
}

func getFieldName(sf reflect.StructField) string {
	jsonTag := sf.Tag.Get("json")
	if jsonTag != "" {
		parts := strings.Split(jsonTag, ",")
		if parts[0] != "" {
			return parts[0]
		}
	}
	return strings.ToLower(sf.Name)
}

func goTypeToSchemaType(t reflect.Type) schema.DataType {
	if t == reflect.TypeOf(time.Time{}) {
		return Time
	}

	switch t.Kind() {
	case reflect.Bool:
		return Bool
	case reflect.Int:
		return Int
	case reflect.Int8:
		return Int8
	case reflect.Int16:
		return Int16
	case reflect.Int32:
		return Int32
	case reflect.Int64:
		return Int64
	case reflect.Uint:
		return Uint
	case reflect.Uint8:
		return Uint8
	case reflect.Uint16:
		return Uint16
	case reflect.Uint32:
		return Uint32
	case reflect.Uint64:
		return Uint64
	case reflect.Float32, reflect.Float64:
		return Float
	case reflect.String:
		return String
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return Bytes
		}
		return JSON
	case reflect.Map, reflect.Struct, reflect.Interface:
		return JSON
	default:
		return String
	}
}

func parseFieldTag(tag string, f *Field) {
	if tag == "" {
		return
	}

	for _, part := range strings.Split(tag, ",") {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		kv := strings.SplitN(part, ":", 2)
		key := strings.TrimSpace(kv[0])
		val := ""
		if len(kv) > 1 {
			val = strings.TrimSpace(kv[1])
		}

		switch key {
		case "type":
			f.Type = schema.DataType(val)
		case "size":
			f.Size, _ = strconv.ParseUint(val, 10, 64)
		case "label":
			f.Label = val
		case "default":
			f.Default = val
		case "comment":
			f.Comment = val
		case "nullable":
			f.Nullable = true
		case "unique":
			f.Unique = true
		case "index":
			f.Index = true
		case "readonly":
			f.Options.ReadOnly = true
		case "crypt":
			f.Options.Crypt = val
		case "array":
			f.Options.IsArray = true
		case "format":
			f.Options.FormatTime = val
		}
	}
}
