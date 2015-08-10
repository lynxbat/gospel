package truths

func (t *Truths) TruthGolangVersion() (Truth, error) {
	return &truth{
		key:         "golang-version",
		description: "Returns current version of golang detected",
		gather:      goVersion,
	}, nil
}

func goVersion() interface{} {
	return "v1"
}

/*

golang-version
description
tags [language, version, golang]

*/
