//go:build js && wasm

package js

import (
	"fmt"
	"strings"
	"syscall/js"
)

type HTMLElement struct {
	value js.Value
}

func (elem *HTMLElement) Type() string {
	return strings.ToLower(elem.value.Get("tagName").String())
}

func (elem *HTMLElement) Append(nodes ...*HTMLElement) *HTMLElement {
	for _, node := range nodes {
		elem.value.Call("append", node.value)
	}
	return elem
}

func (elem *HTMLElement) Style(style string, value string) *HTMLElement {
	elem.value.Get("style").Set(style, value)
	return elem
}

func (elem *HTMLElement) Styles(styles ...string) *HTMLElement {
	if len(styles)%2 != 0 {
		panic(fmt.Errorf("odd number of arguments passed to Styles"))
	}
	for i := range len(styles) / 2 {
		elem.value.Get("style").Set(styles[i*2], styles[i*2+1])
	}
	return elem
}

type FlexOptions struct {
	Alignment     string
	Justification string
	Gap           string
}

func (elem *HTMLElement) Flex(direction string, opts FlexOptions) *HTMLElement {
	elem.Style("display", "flex")
	elem.Style("flexDirection", direction)
	elem.Style("alignItems", opts.Alignment)
	elem.Style("justifyContent", opts.Justification)
	elem.Style("gap", opts.Gap)
	return elem
}

func (elem *HTMLElement) Clone() *HTMLElement {
	jsVal := elem.value.Call("cloneNode", true)
	cloned := new(HTMLElement)
	cloned.value = jsVal
	return cloned
}

func (elem *HTMLElement) Class(class string) *HTMLElement {
	elem.value.Get("classList").Call("add", class)
	return elem
}

func (elem *HTMLElement) SetText(content string) *HTMLElement {
	elem.value.Set("textContent", content)
	return elem
}

func (elem *HTMLElement) Text() string {
	return elem.value.Get("textContent").String()
}

func (elem *HTMLElement) Bind(event string, handler func(this js.Value, args []js.Value) any) *HTMLElement {
	elem.value.Set(event, js.FuncOf(handler))
	return elem
}

func (elem *HTMLElement) BindVoid(event string, handler func()) *HTMLElement {
	elem.value.Set(event, js.FuncOf(func(this js.Value, args []js.Value) any { handler(); return nil }))
	return elem
}

func CreateElement(tag string) *HTMLElement {
	jsElem := js.Global().Get("top").Get("document").Call("createElement", tag)
	if jsElem.IsUndefined() || jsElem.IsNull() {
		return nil
	}

	elem := new(HTMLElement)
	elem.value = jsElem
	return elem
}

func GetBody() *HTMLElement {
	elems, _ := GetElementsByTagName("body")
	return elems[0]
}

func Main() *HTMLElement {
	elems, _ := GetElementsByTagName("main")
	return elems[0]
}

func Head() *HTMLElement {
	elems, _ := GetElementsByTagName("head")
	return elems[0]
}

func RootNode() *HTMLElement {
	elems, _ := GetElementsByTagName("html")
	return elems[0]
}
