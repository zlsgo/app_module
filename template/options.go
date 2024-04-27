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

var options = Options{
	Dir: "./views",
}
