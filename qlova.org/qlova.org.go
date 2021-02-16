package main

import (
	"log"
	"net/http"

	"qlova.tech/rgb"

	"qlova.org/seed/new/app"
	"qlova.org/seed/new/page"

	"qlova.org/site"
	"qlova.org/site/qlova.org/ui"
)

var modules = []string{
	"/seed",
	"/mirror",
	"/should",
	"/send",
}

var website = app.New("Qlova",
	app.SetIcon("logo.png"),
	app.SetColor(rgb.White),

	app.OnUpdateFound(app.Update()),

	page.AddPages(ui.QlovaSplash{}),
	page.Set(ui.QlovaSplash{}),
).Handler()

func main() {
	for _, module := range modules {
		http.HandleFunc(module, func(w http.ResponseWriter, r *http.Request) {
			if site.ServeModule(module, "qlova.org", "github.com/qlova", w, r) {
				return
			}

			website.ServeHTTP(w, r)
		})
	}

	http.Handle("/", website)

	if err := http.ListenAndServe(":21075", nil); err != nil {
		log.Fatalln(err)
	}
}
