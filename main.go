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
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {

		log.Println(req.Host)
		log.Println(req.Form)
		log.Println(req.Proto)
		log.Println(req.URL)
		log.Println(req.Referer())

		tmpl := template.New("template")
		_, err := tmpl.Parse(htmlString)
		if err != nil {
			log.Println("Error:", err)
		}

		data := htmlData{
			Method:  req.Method,
			Headers: req.Header,
		}

		var body []byte
		_, err = req.Body.Read(body)
		if err != nil {
			log.Println("Error:", err)
		}

		if len(body) == 0 {
			data.Body = "Empty body"

		} else {
			data.Body = string(body)
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			log.Println("Error:", err)
		}
	})

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
