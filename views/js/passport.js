const doRegister = async function(){
    let email = document.getElementById('txtEmail').value;
    let password = document.getElementById('txtPassword').value;
    var res = await register(email, password);
}

const doSignIn = async function(){
    let email = document.getElementById('txtEmail').value;
    let password = document.getElementById('txtPassword').value;
    var res = await login(email, password);
}