package main

import (
	"log"
	"net/http"

	"github.com/kataras/iris/v12/view"
)

func main() {
	engine := view.HTML("./views", ".html")
	engine.Layout("layouts/main.html") // set default layout
	err := engine.Load()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		err := engine.ExecuteWriter(w, "index.html", "", map[string]interface{}{
			"title":   "iris view",
			"message": "hello",
		})
		if err != nil {
			log.Fatal(err)
		}
	}))

	http.Handle("/set-layout", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// specific another layout
		err := engine.ExecuteWriter(w, "second.html", "layouts/another.html", map[string]interface{}{
			"title":   "iris view",
			"message": "hello",
		})
		if err != nil {
			log.Fatal(err)
		}
	}))

	log.Println(("serve 0.0.0.0:8080"))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
