# model

## use

```go
model.New(func(o *model.Options) {
    o.SchemaApi = "/api/v1/schema"

    o.SchemaMiddleware = func() []znet.Handler {
    	return []znet.Handler{}
    }

    o.Schemas.Append(define.Schema{})
})
```