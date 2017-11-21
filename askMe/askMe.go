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

	//back ticks instead of quotes to make sure it doesnt leave the characters first
	if matched,_ := regexp.MatchString(`(?i). *\bjoke\b.*`,input); matched {
		//return the string below
		return "What do you call a row of rabbits jumping backwards?"
	}

	if matched,_ := regexp.MatchString(`(?i). *\bwhat\b.*`,input); matched {
		//return the string below
		return "A receeding hairline!"
	}


	//to capture the name
	//re := regexp.MustCompile("([Mm]y name is) ([^.!?]*)[.!?]?)")
	re := regexp.MustCompile("[Ii] (?:A|a)(?:M|m) ([^.!?]*)[.!?]?")
	

	if re.MatchString(input){	
		return re.ReplaceAllString(input, "Hello $1. Tell me about yourself") 

	}

	//to capture i feel
	re2 := regexp.MustCompile("[Ii] (?:F|f)(?:E|e)(?:E|e)(?:L|l) ([^.!?]*)[.!?]?")
	

	if re2.MatchString(input){	
		return re2.ReplaceAllString(input, "How long have you felt $1?") 

	}
/*
	re3 := regexp.MustCompile("([^.!?]*) (?:J|j)(?:O|o)(?:K|k)(?:E|e) ([^.!?]*)[.!?]?")
	

	if re3.MatchString(input){	
		return re3.ReplaceAllString(input, "How long have you felt $1?") 

	}
	*/

	

	

	answers := []string{
		"I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		 "Why do you say that?",
		 "Would you like to hear a joke?",

	}

	//returning a single string response
		return answers[rand.Intn(len(answers))]
}



