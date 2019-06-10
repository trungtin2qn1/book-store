if (document.readtState ==  'loading') {
	document.addEventListener('DOMContentLoaded', ready);
}
else
{
	ready();
}

function ready(){
	var removeItemsCartBtn = document.getElementsByClassName('btn-danger');
	console.log(removeItemsCartBtn);
	for (var i = 0; i < removeItemsCartBtn.length ; i ++) {
		var btn = removeItemsCartBtn[i];
		btn.addEventListener('click', removeCartItem)
	}

	var quantityInput = document.getElementsByClassName('cart-quantity');
	for (var i = 0; i < quantityInput.length ; i ++) {
		var input = quantityInput[i];
		input.addEventListener('change', quantityChanged)
	}

}


function quantityChanged(event){
	var input = event.target;
	if(isNaN(input.value) || input.value <= 0){
		input.value = 1;
	}

	updateCartTotal();
 }

function removeCartItem(event){
	var btnClicked = event.target;
	btnClicked.parentElement.parentElement.remove();
	updateCartTotal();
}

function updateCartTotal(){
	var cartItem = document.getElementsByClassName('cart-items')[0];
	var cartRows = cartItem.getElementsByClassName('cart-row');
	var total = 0;
	for (var i = 0; i < cartRows.length ; i ++) {
		var cartRow = cartRows[i];
		var priceItem = cartRow.getElementsByClassName('cart-price')[0];
		var quantityItem = cartRow.getElementsByClassName('cart-quantity')[0];

		var price = parseFloat(priceItem.innerText.replace('$','')); 
		var quantity = quantityItem.value;
		total = total + (price * quantity);
	}
	total = Math.round(total * 100)/100;
	document.getElementsByClassName('cart-total-price')[0].innerText = '$' + total;
}