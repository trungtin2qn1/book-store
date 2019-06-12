
var password = document.getElementById("txtPassword"),
    confirm_password = document.getElementById("txtConfirmPassword"),
    email = document.getElementById("txtEmail");

function validatePassword() {
    if (password.value != confirm_password.value) {
        confirm_password.setCustomValidity("Your password Don't Match!");
    } else {
        confirm_password.setCustomValidity('');
    }
}
password.onchange = validatePassword;
confirm_password.onkeyup = validatePassword;


const doRegister = async function (email, password) {
    // var res = await register(email.value, password.value);
    // console.log(res);
    alert(`Sign Up Done! ${email.value}`);
}