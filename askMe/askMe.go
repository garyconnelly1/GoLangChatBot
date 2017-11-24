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
	"strings"
	"time"
	
	
	
)

//function to generate the Eliza responses
func ElizaResponse(inputStr string) string{

	input :=inputStr

	//seed the random output to the time
	rand.Seed(time.Now().UTC().UnixNano())

	

//reflecting pronouns
		if strings.Contains(strings.ToLower(input), "I really am not sure that you understand the effect your questions are having on me."){

			//reflecting pronouns
			if strings.Contains(strings.ToLower(input), "your") {

				test := regexp.MustCompile(`your`)

				var reflectsResponse = test.ReplaceAllString(input, "my")
				return reflectsResponse
			}


			if strings.Contains(strings.ToLower(input), "me") {

				test := regexp.MustCompile(`me`)

				var reflectsResponse = test.ReplaceAllString(input, "you")
				return reflectsResponse
			}




			if strings.Contains(strings.ToLower(input), "you") {

				test := regexp.MustCompile(`you`)

				var reflectsResponse = test.ReplaceAllString(input, "I")
				return reflectsResponse
			}


			if strings.Contains(strings.ToLower(input), "my") {

				test := regexp.MustCompile(`my`)

				var reflectsResponse = test.ReplaceAllString(input, "your")
				return reflectsResponse
			}

		}
			

			//ending reflecting pronouns

			

    //to process hello
	if strings.Contains(strings.ToLower(input), "hello"){
		//return the string below
		return "Hello friend. What is your name?"
	}

	//if the user wants to hear a joke
	if matched,_ := regexp.MatchString(`(?i). *\bjokes\b.*`,input); matched {
		//return the string below
		return "Would you like to hear a joke?"
	}

	//telling the joke
	if matched,_ := regexp.MatchString(`(?i). *\bjoke\b.*`,input); matched {
		//return the string below
		return "What do you call a row of rabbits jumping backwards?"
	}

	//telling the punch line
	if matched,_ := regexp.MatchString(`(?i). *\bwhat\b.*`,input); matched {
		//return the string below
		return "A receeding hairline!"
	}


	//to capture I am
	//regex syntax for this was covered in the third problem sheet
	re := regexp.MustCompile("[Ii] (?:A|a)(?:M|m) ([^.!?]*)[.!?]?")
	
	//easy response for somone saying they are x
	if re.MatchString(input){	
		return re.ReplaceAllString(input, "How do you know you are $1?") 
	}

	//to capture i feel
	//very similar syntax to capturing I am
	re2 := regexp.MustCompile("[Ii] (?:F|f)(?:E|e)(?:E|e)(?:L|l) ([^.!?]*)[.!?]?")
	

	if re2.MatchString(input){	
		return re2.ReplaceAllString(input, "How long have you felt $1?") 
	}



	//capture age
	re4 := regexp.MustCompile("[Ii] (?:T|t)(?:U|u)(?:R|r)(?:N|n)(?:E|e)(?:D|d) ([0-9]+)[.!?]?")
	
	if re4.MatchString(input){	
		return re4.ReplaceAllString(input, "When did you turn $1?") 
	}


	//to capture the name of the user
	re3 := regexp.MustCompile("[Mm]y name is ([^.!?]*)[.!?]?")
	
	//hello gary. tell me about yourself.
	if re3.MatchString(input){	
		return re3.ReplaceAllString(input, "Hello $1. Tell me about yourself") 

	}

	//to capture I like
	reLike := regexp.MustCompile("[Ii] (?:L|l)(?:I|i)(?:K|k)(?:E|e) ([^.!?]*)[.!?]?")
	
	//oh nice, why do you like pizza/football etc
	if reLike.MatchString(input){	
		return reLike.ReplaceAllString(input, "Oh nice, why do you like $1?") 
	}


	//try to recognise emotions
	//if the input string ends in an exclamation mark, it typically means the user is angry about something
	reAngry := regexp.MustCompile("([^.!?]*)(!)")
	
	if reAngry.MatchString(input){	
		return "Why are you getting angry?" 
	}


	
//default answers if it doesnt recognise the user input
	answers := []string{
		"I'm not sure what you are trying to say. Could you explain it to me?",
		"How does that make you feel?",
		"Why do you say that?",
		"Would you like to hear a joke?",
		"How is your day going?",
		"Are you a college student like my creator?",
		"What hobbies do you have?",
		"Do you play any sport?",
		"What age are you?",

	}

	//returning a single string response
		return answers[rand.Intn(len(answers))]
}

//
/*
//adapted from https://gist.github.com/ianmcloughlin/c4c2b8dc586d06943f54b75d9e2250fe
func refPro(inputStr string) string{
	boundaries := regexp.MustCompile(`\b`)
	tokens := boundaries.Split(inputStr, -1)
	
	// List the reflections.
	reflections := [][]string{
		{`I`, `you`},
		{`me`, `you`},
		{`am`, `are`},
		{`you`, `I`},
		{`my`, `your`},
		{`your`, `my`},
		{`I am`,`you are`},
	}
	
	// Loop through each token, reflecting it if there's a match.
	for i, token := range tokens {
		for _, reflection := range reflections {
			if matched, _ := regexp.MatchString(reflection[0], token); matched {
				tokens[i] = reflection[1]
				
				break
			}
		}
	}
	
	// Put the tokens back together.
	return strings.Join(tokens, ``)
}
	//end part 5
*/