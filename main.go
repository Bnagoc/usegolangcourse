package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awsome site!</h1>")
}

func contactFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:uladzislau.samuseu@gmail.com\">uladzislau.samuseu@gmail.com</a>.")
}

func faqFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>FAQ Page</h1>
	<ul>
		<li><b>Is there a free version?</b> Yes! we offer a free trial for 30 days on any paid plans.</li>
		<li><b>What are your support hours?</b> We have support staff answering emails 24/7, though responce times may be a bit slower on weekends.</li>
		<li><b>How do I contact support?</b> Email us - <a href="mailto:uladzislau.samuseu@gmail.com">uladzislau.samuseu@gmail.com</a>.</li>
	</ul>
	`)
}

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactFunc(w, r)
	case "/faq":
		faqFunc(w, r)
	default:
		http.NotFound(w, r)
	}
}

func main() {
	var router Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", router)
}
