//package main

/*
	"math/rand"
	"time"
	"fmt"
	"bufio"
	"os"
	"strings"
*/
package askMe


//imports
import (
	
	"regexp"
	"math/rand"
	
	
	
)

func ElizaResponse(inputStr string) string{

	input :=inputStr

	re := regexp.MustCompile("[Ii] (?:A|a)(?:M|m) ([^.!?]*)[.!?]?")
	

	if re.MatchString(input){	
		return re.ReplaceAllString(input, "How do you know you are $1?") 

		

	}
	answers := []string{
		"I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		 "Why do you say that?",

	}

	//returning a single string response
		return answers[rand.Intn(len(answers))]
}



