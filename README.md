# GoLangChatBot
## Name: Gary Connelly.
## Student ID: G00336847
## Module: Data Representation and Querying.
## Lecturer: Ian Mcloughlin.

This is the readme file to go with the GoLang chatbot I created for a 3rd year project in Data Representation and Querying in GMIT. It is an implementation of the famous Eliza chatbot.

## How to run the program:
If your machine does not already have Git and Go installed, they can be installed by going to 
 https://golang.org/dl/
 https://git-scm.com/downloads
 
 To clone the repository to your local machine, in command prompt enter
 ```
 https://github.com/garyconnelly1/GoLangChatBot.git
 ```
### Step 1:
Build the ChatBot.go file.
```
go build ChatBot.go
```

###Step 2:
Execute the Chatbot.go file.
```
ChatBot.exe
```
###Step 3:
Serve the web app on your local host by entering the following code into the search bar on your preferred browser.
```
127.0.0.1:8080
```

## How it was made.
### Resources:
When we were first given this project I had to do a lot of research in order to educate myself on how to go about making a chat bot.
The first thing I did was google Eliza chatBot and I was presented with a very interesting artical on wikipedia.
https://en.wikipedia.org/wiki/ELIZA . Now that I knew what Eliza was I started thinking about how it could actually be built. I found a very good implementation of the eliza chatbot written in python.  
https://www.smallsurething.com/implementing-the-famous-eliza-chatbot-in-python/ . This gave me a very good understanding of how language processing might actually work in a computer program. Finally we were given lab problems to do which helped emensly with some of the vital syntax involved in coding a language processer. 
https://github.com/garyconnelly1/Go-Problems-3.

### Development:
After all my research I felt very comfortable that I could tackle the problem at hand. The first thing I did was create a simple go file that served a very simple html web page. I designed the front end of the webpage with a form that can take in input from the user.
https://www.w3schools.com/html/html_forms.asp . The hardest part was creating the javascript file to asynchronously append the users input along with elizas response. But with the help of a few friends we managed to get the webApp to behave the way we wanted it to. Then I made a second go file with a function in it called elizaResponse which takes the user input as a parameter and processes it before returning the appropriate response. The response is then sent back to the first go file which is served to the web application. 

### Features:
Start a conversation by simply saying hello. You can talk about hobbies you have or just how you are feeling or even ask her to tell you a joke. I have a set of default responses also which come into play if Eliza doesnt recognise what the input is. These default responses are engineered to try to get the user talking about topics which Eliza will be able to process and understand. One problem I had was when I was trying to reflect the pronouns. ie changing I to you etc. I was able to get the pronoun to reflect but they were starting the mess with the other responses I had programmed. An example of this would be trying to reflect the pronouns in "I am", however I already had a response ready for when the user enters "I am ....", which started behaving weird when implementing the pronouns. So in order to combat this, while it may be far from ideal, I just put all the reflections in one if statement which would only be triggered if the following sentence was entered - "I really am not sure that you understand the effect your questions are having on me.". This sentence was taken from the lab problems I spoke of earlier. Give it a try and see how it works! I also added a small feature which helps eliza detect if somone is angry. This is triggered when the users input ends with an exclamation mark.

## Final thoughts.
It was a fascinating experience creating this project. It was the first time I got to dip my toe in the realm of artificial intelligence, even though I am aware it is not real artificial intelligence it was the closest thing I have come to working with it. I can also see how a project like this can easily be scaled to the extent of seeming like a real human you are talking to. I learned a lot about language proccessing technology while researching this project and I hope to build on this knowledge in the future.


