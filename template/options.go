package template

type Options struct {
	// Dir 模板目录
	Dir string `json:"dir"`
	// Static 静态文件路由
	Static string `json:"static"`
	// StaticDir 静态文件目录
	StaticDir string `json:"static_dir"`
	// Reload 是否开启模板热更新
	Reload bool `json:"reload"`
}

func (Options) ConfKey() string {
	return "templates"
}

func (Options) DisableWrite() bool {
	return true
}

var options = Options{
	Dir: "./views",
}
