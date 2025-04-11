package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

const tpl = `
<!DOCTYPE html>
<html>
    <head>
		        <meta charset="UTF-8">
						        <title>{{toUpper .Title}}</title>
										    </head>
												    <body>
				{{range .Items}}<ul><li>{{ . }}</li></ul>{{else}}<div><strong>no rows</strong></div>{{end}}
																		    </body>
																				</html>`

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func animeHandler(w http.ResponseWriter, r *http.Request) {

	funcMap := map[string]interface{}{
		"toUpper": strings.ToUpper,
	}

	//	t, err := template.New("webpage").Parse(tpl)
	t := template.Must(template.New("webpage").Funcs(funcMap).Parse(tpl)) // it will panic immediately when there is an error , so no need to declare an err as I use it in template.New alone.
	//	checkError(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My Favorate animes",
		Items: []string{
			"Berzerk",
			"Trigun",
			"Inial D",
			"Vinland Saga",
			"Cowboy bebop",
			"Monster",
			"Kaiji",
			"Akagi",
		},
	}

	err := t.Execute(w, data)
	checkError(err)

}
func mangaHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.New("webpage").Parse(tpl)
	checkError(err)

	data := struct {
		Title string
		Items []string
	}{
		Title: "My Favorate manga",
		Items: []string{},
	}

	err = t.Execute(w, data)
	checkError(err)

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "6969"
	}
	fmt.Println("port is: ", port)
	log.Println("starting the server on port")
	http.HandleFunc("/anime", animeHandler)
	http.HandleFunc("/manga", mangaHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
