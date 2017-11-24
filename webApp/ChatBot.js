

//everything in this file only occures once the butten is presses

  function myFunction(){
      //constants to get the ids of the input and future output
      const form = $("#userInput");
      const listItem = $("#list-group");


       //prevent default function prevents the page being reloaded once the button is pressed
       //check https://www.w3schools.com/jsref/event_preventdefault.asp for more info on this
        event.preventDefault();
        
        //create constant to get the form value
          const input = form.val();
          
        //to clear the input box after every inout  
        form.val(" ");

        //append the input to the chat list on the right hand side of the page
        listItem.append("<li class='list-group-item list-group-item-primary text-right'>" + "User : " + input + "</li>");
        
        
        // GET/POST methods
        //create a constant for the query parameters
        const queryParams = {
            "userInput": input
        }
        //jquery for getting the chat
        $.get("/chat", queryParams)
            .done(function(resp) {
                //next list item is elizas response which is served on the left hand side of the page
                var nextListItem = "<li class='list-group-item list-group-item-info'>" + "ELiza : " + resp + "</li>";
                //function to delay the responce by a second to simulate somone typing back
                setTimeout(function() {
                    listItem.append(nextListItem)
                    //auto fix to the bottom of the page so you dont have to scroll down to see the chat
                    // for the auto fix to the bottom go to https://stackoverflow.com/questions/47425453/html-css-auto-scroll-page-to-bottom
                     $("html, body").scrollTop($("body").height());
                }, 1000); //set timeout to give wait to response
            }).fail(function() {
                var nextListItem = "<li class='list-group-item list-group-item-danger' >Eliza has gone to sleep. Try again later.</li>";
                listItem.append(nextListItem);
            });
            
           
    
  }



