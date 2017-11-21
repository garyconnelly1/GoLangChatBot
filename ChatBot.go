// An echo web application.
// problem sheet 2

//https://github.com/data-representation
//https://developer.mozilla.org/en-US/docs/AJAX/Getting_Started

// "fmt"

//"net/http"
//"html/template"//add html/template package 
//	"bytes"
//"html/template"

package main

import (
	
	"fmt"
	
	"net/http"
	


)
	
type myMsg struct {
	Input string
	Output string
	Previous string
}

//func requestHandler(w http.ResponseWriter, r *http.Request) {
	//serve the homepage.html file
//	http.ServeFile(w, r, "ChatBot.html")
//}
//////////////////////////////////////////////////
func chatHandler(w http.ResponseWriter, r *http.Request) {

	//userInput := r.URL.Query().Get("userInput")
	//reply := eliza.AskEliza(userInput)
	response := "hello, tell me about yourself"
	fmt.Fprintf(w, response)
/*
	elizaResponse := "Hello User, how are you?"

	userResponse := r.FormValue("userResponse")

	t, _ := template.ParseFiles("ChatBot.html")

	t.Execute(w, &myMsg{Output: elizaResponse, Input: userResponse})
	*/
	

	//create and initialise string
	//output		:= "Howya lad"
	//input		:= r.FormValue("chat")
	//previous	:= input
	

}//chatHandler

func main() {
	// handles root page
	//http.HandleFunc("/", requestHandler)
	//serve the files from the /static folder
	directory := http.Dir("./webApp")
	fileServer := http.FileServer(directory)

	http.Handle("/", fileServer)

	//handle /chat page
	http.HandleFunc("/chat", chatHandler)
	http.ListenAndServe(":8080", nil)
}