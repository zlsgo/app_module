# member

## use

```go
memberModule := member.New("xxx", func(o *member.Options) {
    o.ApiPrefix = "/app/member"
    o.Providers = []auth.AuthProvider{&auth.Weapp{
        AppId:     "wx55a57ece33099d66",
        AppSecret: "",
    }}
    o.EnabledProviders = []string{"weapp"}
})
```