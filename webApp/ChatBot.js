

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

        // GET/POST
        const queryParams = {
            "userInput": input
        }
        $.get("/Chat", queryParams)
            .done(function(resp) {
                var nextListItem = "<li class='list-group-item list-group-item-info'>" + "ELiza : " + resp + "</li>";
                setTimeout(function() {
                    listItem.append(nextListItem)
                }, 1000); //set timeout to give wait to response
            }).fail(function() {
                var nextListItem = "<li class='list-group-item list-group-item-danger' >Sorry I'm not home right now.</li>";
                list.append(nextListItem);
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

