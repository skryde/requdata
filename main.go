package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type htmlData struct {
	Method  string
	Body    string
	Headers map[string][]string
	QueryParams map[string][]string
	Hostname string
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		tmpl := template.New("template")
		_, err := tmpl.Parse(htmlString)
		if err != nil {
			log.Println("Error:", err)
		}

		hostname, _ := os.Hostname()
		data := htmlData{
			Method:  req.Method,
			Headers: req.Header,
			QueryParams: req.URL.Query(),
			Hostname: hostname,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error:", err)
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
