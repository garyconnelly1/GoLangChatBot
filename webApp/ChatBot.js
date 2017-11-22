

//add the keypress function for when the user types an input for eliza to compute

//$(document).ready(() => {
//function myFunction(){
   // document.addEventListener("keydown", function(event) {

        console.log("well");

  //  }

  function myFunction(){
      const form = $("#userInput");
const listItem = $("#list-group");


       console.log("well lawd");
        event.preventDefault();

          const input = form.val();
          
        form.val(" ");
console.log("befor")
console.log(input)
        listItem.append("<li class='list-group-item list-group-item-primary text-right'>" + "User : " + input + "</li>");
        /*
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        listItem.append("<li class='list-group-item list-group-item-success'>" + "User : " + input + "</li>");
        */
        console.log("after")
//hehe
        // GET/POST
        const queryParams = {
            "userInput": input
        }
        $.get("/chat", queryParams)
            .done(function(resp) {
                var nextListItem = "<li class='list-group-item list-group-item-info'>" + "ELiza : " + resp + "</li>";
                setTimeout(function() {
                    listItem.append(nextListItem)
                    // for the auto fix to the bottom go to https://stackoverflow.com/questions/47425453/html-css-auto-scroll-page-to-bottom
                     $("html, body").scrollTop($("body").height());
                }, 1000); //set timeout to give wait to response
            }).fail(function() {
                var nextListItem = "<li class='list-group-item list-group-item-danger' >Sorry I'm not home right now.</li>";
                listItem.append(nextListItem);
            });
            
           
    
  }

/*
    form.keypress(function(event) {
console.log("inside key press");
        //var keyCoder = 13;
    

       // let keyCodes = event.keyCode;

        //for enter
        if (event.keyCode != 13) {
            return;
        }
        */

       

      
//}

//});

