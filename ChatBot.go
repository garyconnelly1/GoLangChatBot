// An echo web application.
// problem sheet 2

//https://github.com/data-representation
//https://developer.mozilla.org/en-US/docs/AJAX/Getting_Started

// "fmt"

//"net/http"
//"html/template"//add html/template package 
//	"bytes"

package main

import (
	"html/template"
	"fmt"
	
	"net/http"
	


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

	userInput := r.URL.Query().Get("userInput")
	//reply := eliza.AskEliza(userInput)
	fmt.Fprintf(w, userInput)

	elizaResponse := "Hello User, how are you?"

	userResponse := r.FormValue("userResponse")

	t, _ := template.ParseFiles("ChatBot.html")

	t.Execute(w, &myMsg{Output: elizaResponse, Input: userResponse})
	

	//create and initialise string
	//output		:= "Howya lad"
	//input		:= r.FormValue("chat")
	//previous	:= input
	
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
//	t, _ := template.ParseFiles("chat.html")

	//execute template and pass pointer to myMsg 	struct
//	t.Execute(w, &myMsg{Input:input,Output:output,Previous:previous})
}//chatHandler

func main() {
	// handles root page
	http.HandleFunc("/", requestHandler)

	//handle /chat page
	http.HandleFunc("/Chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}