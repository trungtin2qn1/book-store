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

	var addItemToCartBtn = document.getElementsByClassName('btn-add-to-cart');
	for (var i = 0; i < addItemToCartBtn.length ; i ++) {
		var btn = addItemToCartBtn[i];
		btn.addEventListener('click', addCartItemClicked);
	} 

}

function addCartItemClicked(event){
	var btnClicked = event.target;
	var pItems = btnClicked.parentElement.parentElement;
	var pro_title = pItems.getElementsByClassName('product-title')[0].innerText;
	var pro_price = pItems.getElementsByClassName('product-price')[0].innerText;
	var pro_img =  pItems.getElementsByClassName('product-image')[0].src;
	console.log(pro_title, pro_price, pro_img);
	addItemToCart(pro_title,pro_price, pro_img);
}

function addItemToCart(title,price,imgsrc){
	var cartRow = document.createElement('tr');
	// cartRow.classList.add('cart-row');
	cartRow.innerText = title;
	var cartItems = document.getElementsByClassName('cart-row').innerHTML;
	console.log(cartItems);
	// var cartRowContent =`
	// 						<td data-th="Product">
 //                                <div class="row">
 //                                    <div class="col-sm-2 hidden-xs">
 //                                        <img class="cart-product-image" src="${imgsrc}" alt="7-Day" class="img-responsive" />
 //                                    </div>
 //                                    <div class="col-sm-10">
 //                                        <h4 class="cart-product-title">${title}</h4>
 //                                        <p class="cart-product-des">Quis aute iure reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Lorem ipsum dolor sit amet.</p>
 //                                    </div>
 //                                </div>
 //                            </td>
 //                            <td class="cart-price" data-th="Price">${price}</td>
 //                            <td  data-th="Quantity">
 //                              <input type="number" class="cart-quantity" value="1">
 //                          </td>
 //                          <td data-th="Subtotal" class="text-center">1.99</td>
 //                          <td class="actions" data-th="">
 //                              <button class="btn btn-info btn-sm"><i class="fa fa-refresh"></i></button>
 //                              <button class="btn gray-btn btn-danger btn-sm"><i class="fa fa-trash-o"></i></button>
 //                          </td>
	// `; 
	// cartRow.innerHTML = cartRowContent;
	// cartItems.append(cartRow);

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