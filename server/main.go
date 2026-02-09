package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:generate sh -c "GOOS=js GOARCH=wasm go build -o web/build.wasm ../client"

//go:embed web
var webfiles embed.FS

func main() {
	webfilesRoot, err := fs.Sub(webfiles, "web")
	if err != nil {
		log.Fatal(err)
	}

	http.ListenAndServe(":8080", http.FileServerFS(webfilesRoot))
}
