package template

type Options struct {
	Funcs     map[string]interface{} `json:"-"`
	Dir       string                 `json:"dir"`
	Static    string                 `json:"static"`
	StaticDir string                 `json:"static_dir"`
	Reload    bool                   `json:"reload"`
}

func (Options) ConfKey() string {
	return "templates"
}

func (Options) DisableWrite() bool {
	return true
}

// defaultOptions 返回模板模块的默认配置。
//
// 每次调用都返回新值：选项会被 service.DefaultConf 长期持有，
// 不能在多次 New 之间共享同一份可变配置。
func defaultOptions() Options {
	return Options{
		Dir: "./views",
	}
}
