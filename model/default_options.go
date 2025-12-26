package model

var DefaultSchemaOptions = SchemaOptions{}

func SetDefaultSchemaOptions(fn func(*SchemaOptions)) {
	if fn == nil {
		return
	}
	fn(&DefaultSchemaOptions)
}
