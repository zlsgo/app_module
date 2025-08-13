package html

type Options struct{}

func (Options) ConfKey() string {
	return "html"
}

func (Options) DisableWrite() bool {
	return true
}

var options = Options{}
