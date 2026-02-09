//go:build js && wasm

package main

import (
	"math/rand"

	"github.com/flipfloppy1/storefront/client/js"
)

const css = `
h1, h2, h3, h4, h5 {
	margin: 0;
}

body {
	background-color: rgb(0, 200, 200);
	height: 100%;
	width: 100%;
	padding: 0;
	margin: 0;
}

main {
	margin: 20px;
	margin-top: 30px;
	user-select: none;
}
`

var randomPhrases = []string{
	"for real",
	"now with store stuff",
	"making bank",
	"idk what to put here",
	"ellipsis...",
	"great goin genghis",
	"myaaaan goes the cat",
	"pop goes the weasel... but why?",
	"you can sometimes get the same one twice",
	"click here to crash your browser",
}

func randomPhrase() string {
	return randomPhrases[rand.Intn(len(randomPhrases))]
}

func init() {
	js.Head().Append(js.CreateElement("style").SetText(css))
}

func main() {
	js.Main().Append(js.CreateElement("h1").SetText("Generic Storefront"))
	h2 := js.CreateElement("h2").SetText(randomPhrase())
	js.Main().Append(h2)
	currText := h2.Text()
	change := func() {
		text := randomPhrase()
		h2.SetText(text)
		for text == currText {
			text = randomPhrase()
			h2.SetText(text)
		}
		currText = text
	}
	js.RootNode().BindVoid("onclick", change)

	div := js.CreateElement("div").Styles(
		"minWidth", "300px",
		"minHeight", "200px",
		"alignContent", "center",
		"fontSize", "100px",
		"textAlign", "center",
	)

	items := js.CreateElement("div")
	items.Flex("row", js.FlexOptions{Justification: "space-between", Alignment: "center", Gap: "20px"}).Styles(
		"overflow", "hidden",
	)

	itemsItems := js.CreateElement("div").Append(
		items.Append(
			div.SetText("item 1").Style("backgroundColor", "red"),
			div.Clone().SetText("item 2").Style("backgroundColor", "blue"),
			div.Clone().SetText("item 3").Style("backgroundColor", "lightgreen"),
		),
		items.Clone(),
	).Flex("column", js.FlexOptions{Gap: "20px", Alignment: "center"})

	js.Main().Append(itemsItems)

	hang()
}

func hang() {
	<-make(chan int)
}
