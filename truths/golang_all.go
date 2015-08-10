package truths

import (
	"runtime"
)

func (t *Truths) TruthGolangVersion() (Truth, error) {
	return &truth{
		key:         "golang-version",
		description: "The version of golang",
		gather: func() interface{} {
			return runtime.Version()
		},
		tags: []string{"lang", "version", "golang"},
	}, nil
}

func (t *Truths) TruthGolangOS() (Truth, error) {
	return &truth{
		key:         "golang-os",
		description: "The operation system golang detects",
		gather: func() interface{} {
			return runtime.GOOS
		},
		tags: []string{"lang", "os", "golang"},
	}, nil
}

func (t *Truths) TruthGolangArch() (Truth, error) {
	return &truth{
		key:         "golang-arch",
		description: "The architecture golang detects",
		gather: func() interface{} {
			return runtime.GOARCH
		},
		tags: []string{"lang", "architecture", "golang"},
	}, nil
}
