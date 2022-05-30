package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/long74100/gophercises/cyoa/stories"
)

var defaultTemplate = `
<!DOCTYPE html>
<html>
<head>
	<title>CYOA</title>
</head>
<body>
	Hello World!
	{{ .Title }}
</body>
</html>
`

func CustomHandler(story stories.Story) http.Handler {
	return handler{story}
}

type handler struct {
	story stories.Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("").Parse(defaultTemplate))
	err := tpl.Execute(w, h.story["intro"])

	if err != nil {
		panic(err)
	}
}

func JSONToStory(reader io.Reader) (stories.Story, error) {
	decoder := json.NewDecoder(reader)
	var story stories.Story
	if err := decoder.Decode(&story); err != nil {
		return nil, err
	}

	return story, nil
}

func main() {
	fileName := flag.String("file", "gopher.json", "JSON for the story")
	flag.Parse()
	fmt.Printf("hello %s", *fileName)

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := JSONToStory(file)

	if err != nil {
		panic(err)
	}

	handler := CustomHandler(story)
	log.Fatal(http.ListenAndServe(":3000", handler))

}
