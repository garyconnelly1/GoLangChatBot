const form = $("#userInput");
//const listItem = $(".list-group-item");

//add the keypress function for when the user types an input for eliza to compute
form.keypress(function(event){
    //for enter
    if(event.keyCode != 13){ 
        return;
    }

      event.preventDefault(); 

 });