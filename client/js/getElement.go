//go:build js && wasm

package js

import (
	"fmt"
	"syscall/js"
)

func getElementsOfCollection(htmlCollection js.Value) []*HTMLElement {
	slice := make([]*HTMLElement, 0)
	for i := range htmlCollection.Length() {
		elem := new(HTMLElement)
		elem.value = htmlCollection.Index(i)
		slice = append(slice, elem)
	}

	return slice
}

func GetElementsByTagName(tag string) (_ []*HTMLElement, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = fmt.Errorf("error in GetElementsByTagName %v", rec)
		}
	}()
	return getElementsOfCollection(js.Global().Get("top").Get("document").Call("getElementsByTagName", tag)), nil
}

func GetElementsByClassName(class string) (_ []*HTMLElement, err error) {
	defer func() {
		rec := recover()
		if rec != nil {
			err = fmt.Errorf("error in GetElementsByTagName %v", rec)
		}
	}()
	return getElementsOfCollection(js.Global().Get("top").Get("document").Call("getElementsByClassName", class)), nil
}
