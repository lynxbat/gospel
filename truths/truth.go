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
	Name() string
	Description() string
	Tags() []string
	Gather() interface{}
}

// A Truth which has been gathered and stores the truth value in Value()
type GatheredTruth interface {
	Truth
	Value() interface{}
}

// Struct providing access to truths and the ability to bind to them
type Truths struct {
	coll map[string]Truth
}

func New() *Truths {
	t := &Truths{}
	t.load()
	return t
}

type truth struct {
	name        string
	description string
	tags        []string
	gather      func() interface{}
	value       interface{}
}

func (t *truth) Name() string {
	return t.name
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

func (t *truth) Value() interface{} {
	return t.value
}

func (t *Truths) load() {
	t.coll = make(map[string]Truth)
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
					t.coll[tr.Name()] = tr
				}
			} else {
				// We cannot allow this.
				panic(fmt.Errorf("Bad signature for *truth function with Truth prefix (%s): %s does not equal <func(*Truths) (Truth, error)>", method.Name, method.Func.String()))
			}
		}
	}
}

func (t *Truths) Gather() map[string]interface{} {
	r := make(map[string]interface{})
	for _, tr := range t.coll {
		r[tr.Name()] = tr.Gather()
	}
	return r
}

func (t *Truths) List() map[string]Truth {
	return t.coll
}
