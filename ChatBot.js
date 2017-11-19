const form = $("#userInput");
const listItem = $(".list-group-item");

//add the keypress function for when the user types an input for eliza to compute
form.keypress(function(event){
    //for enter
    if(event.keyCode != 13){ 
        return;
    }

      event.preventDefault(); 

      const input = form.val();
      form.val(" ");

      //using GET method to get the form value data from the html form
      $.get('/chat', { value: form.val() } )
            .done(function (data) {
              $('listItem').val(data);
              });
      

 });