package truths

import (
	"fmt"
	"reflect"
	"strings"
)

const (
	// truthPrefix is the prefix used to select *truths.* functions used to provide truths
	truthPrefix = "Truth"
)

// Truth provides an interface for truths
type Truth interface {
	Key() string
	Description() string
	Tags() []string
	Gather() interface{}
}

// Struct providing access to truths and the ability to bind to them
type Truths struct{}

func New() *Truths {
	return &Truths{}
}

type truth struct {
	key         string
	description string
	tags        []string
	gather      func() interface{}
}

func (t *truth) Key() string {
	return t.key
}

func (t *truth) Description() string {
	return t.description
}

func (t *truth) Tags() []string {
	return t.tags
}

func (t *truth) Gather() interface{} {
	return t.gather()
}

func (t *Truths) List() {
	collection := make(map[string]Truth)
	v := reflect.TypeOf(t)

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)

		if strings.HasPrefix(method.Name, truthPrefix) {
			if method.Type.NumIn() == 1 &&
				method.Type.NumOut() == 2 &&
				method.Type.Out(0).String() == "truths.Truth" &&
				method.Type.Out(1).String() == "error" {
				m := method.Func.Interface().(func(*Truths) (Truth, error))
				tr, err := m(t)
				if err == nil {
					collection[tr.Key()] = tr
				}
			} else {
				// We cannot allow this.
				panic(fmt.Errorf("Bad signature for *truth function with Truth prefix (%s): %s does not equal <func(*Truths) (Truth, error)>", method.Name, method.Func.String()))
			}
		}
	}

	fmt.Println("Truths")
	for _, t := range collection {
		fmt.Printf("\t%s\n", t.Key())
		fmt.Printf("\t%s\n", t.Gather())
	}

}
