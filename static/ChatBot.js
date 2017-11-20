
const form = $("#userInput");
const listItem = $("#list-group");

//add the keypress function for when the user types an input for eliza to compute

$(document).ready(()=>{

form.keypress(function(event){

    const keyCodes = {
    ENTER : 13
}

let keyCode = event.keyCode;

    //for enter
    if(event.keyCode !=  keyCodes.ENTER){ 
        return;
    }

      event.preventDefault(); 

      const input = form.val();
      form.val(" ");

       listItem.append("<li class='list-group-item list-group-item-success'>"+"User : " + input + "</li>");

      // GET/POST
    const queryParams = {"userInput" : input }
    $.get("/Chat", queryParams)
        .done(function(response){
            var nextListItem = "<li class='list-group-item list-group-item-info'>"+"ELiza : " + response + "</li>";
            setTimeout(function(){
                listItem.append(nextListItem)
            }, 1000);//set timeout to give wait to response
        }).fail(function(){
            var nextListItem = "<li class='list-group-item list-group-item-danger' >Sorry I'm not home right now.</li>";
            list.append(nextListItem);
        });
});
 
  });

  