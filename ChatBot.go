// Gary Connelly
// The Eliza chatbot project
// class: Data representation and querying year 3 semester 1
//Lecturer: Ian Mcloughlin

//https://github.com/data-representation
//https://developer.mozilla.org/en-US/docs/AJAX/Getting_Started


package main

import (
	
	"fmt"
	
	"net/http"

	"./askMe"
	


)
	
type myMsg struct {
	Input string
	Output string
	Previous string
}

//function to process the user input
func chatHandler(w http.ResponseWriter, r *http.Request) {

	//get the input from the html form
	userInput := r.URL.Query().Get("userInput")
	//send that to the askMe folder for further processing
	response := askMe.ElizaResponse(userInput)
	//output the response
	fmt.Fprintf(w, response)

	

}//chatHandler

func main() {
	
	//serve the files from the /webApp folder
	directory := http.Dir("./webApp")
	fileServer := http.FileServer(directory)

	// handles root page
	http.Handle("/", fileServer)

	//handle /chat page
	http.HandleFunc("/chat", chatHandler)
	//serve it on port 8080
	http.ListenAndServe(":8080", nil)
}





