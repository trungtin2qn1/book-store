// var api = require('api');



var password = document.getElementById("txtPassword"),
     confirm_password = document.getElementById("txtConfirmPassword"),
     email = document.getElementById('txtEmail');

 function validatePassword() {
     if (password.value != confirm_password.value) {
         confirm_password.setCustomValidity("Your password Don't Match");
     } else {
         confirm_password.setCustomValidity('');
     }
 }
 password.onchange = validatePassword;
 confirm_password.onkeyup = validatePassword;

 function register(){
    var res = registera(email,password);
    console.log(res);
 }

registera = async function (email, password) {
        //Truyền vào email và password
        //email password

        var myBody = {
            email: email,
            password: password,
        };

        console.log(myBody);

        //Tạo ra 1 http request
        //url: http://127.0.0.1:2013/api.sign_up
        //method: POST
        //body:{
        // "email": email,
        // "password": password
        //}
        //header: "Content-Type": "application/json"

        const response = await fetch('http://127.0.0.1:2013/api/sign_up', {
            method: 'POST',
            body: JSON.stringify(myBody), // string or object
            headers: {
                'Content-Type': 'application/json'
            }
        });
        const myJson = await response.json(); //extract JSON from the http response
        // do something with myJson
        console.log(myJson)

        return (true, myJson)
    }