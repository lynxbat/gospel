package truth

import (
	"fmt"
	"reflect"
)

type TruthCollection interface {
}

type Truth interface {
}

type truthControl struct {
	Collections map[string]func() (TruthCollection, []Truth, error)
}

type truthCollection struct {
	Name string
}

// truth family (string)

func truthHook(collectionName string) TruthCollection {
	return &truthCollection{Name: collectionName}
}

func ListTruths() {
	fmt.Println("Truth Collections")

	t := new(truthControl)
	t.Collections = make(map[string]func() (TruthCollection, []Truth, error))

	v := reflect.ValueOf(t)

	for i := 0; i < v.NumMethod(); i++ {
		method := v.Method(i)
		// fmt.Println(method.String())
		// fmt.Println(method.IsValid())
		rv := method.Call([]reflect.Value{})

		functionName := rv[0].Interface().(*truthCollection).Name

		// fmt.Println(functionName)

		t.Collections[functionName] = method.Interface().(func() (TruthCollection, []Truth, error))
	}

	for k, _ := range t.Collections {
		fmt.Printf("\t%s\n", k)
	}

}
