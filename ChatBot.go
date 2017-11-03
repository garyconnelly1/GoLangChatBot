// An echo web application.
// problem sheet 2

//https://github.com/data-representation
//https://developer.mozilla.org/en-US/docs/AJAX/Getting_Started

// "fmt"
//"net/http"

package main

import (
	"html/template"//add html/template package 
	"net/http"

//	"bytes"
)
	
type myMsg struct {
	Input string
	Output string
	Previous string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	//serve the homepage.html file
	http.ServeFile(w, r, "ChatBot.html")
}

func chatHandler(w http.ResponseWriter, r *http.Request) {

	//create and initialise string
	output		:= "Howya lad"
	input		:= r.FormValue("chat")
	previous	:= input
	
//	// checking for chat URL encoded variable
//	input, err := r.URL.Query()["guess"]
//	// if not found execute the template and exit
//	if !err || len(guess) < 1 {
//		log.Println("Url Param 'guess' is missing")
//		// execute the template with the message
//		tmpl.Execute(w, m)
//		return
//	}// if
	
	// Query()["guess"] will return an array of items, 
	// we only want the single item.
//	g := guess[0]
	
		// adding the guess value to the user value
	//read the contents of chat.html and return a template
	t, _ := template.ParseFiles("chat.html")

	//execute template and pass pointer to myMsg 	struct
	t.Execute(w, &myMsg{Input:input,Output:output,Previous:previous})
}//chatHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /chat page
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}