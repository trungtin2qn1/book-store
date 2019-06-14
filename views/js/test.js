const TestgetAllBooks = async function(input){
    var res = await getAllBooks(input);
    console.log("res:", res);
    console.log("res[0]:", res[0]);
    console.log("res[1]:", res[1]);
}

var count = 0;
var div = document.getElementById('book-recommendes');
const loadMoreBooks = async function(){
	let res = await fetch('http://127.0.0.1:2013/api/auth/customer/1/books');
	let posts = await res.json();
	return posts;
}
function loadBooks(){
	loadMoreBooks()
	.then((posts) => {
		for (var i = 0; i < 10; i++) {
			div.innerHTML +=`
			<div class="col-lg-2 col-md-3 col-sm-4">
                    <div class="item">
                        <a href="product-single.html"><img class="product-image" src="${posts[count].images}" alt="img">
                            <h3 class="product-title">${posts[count].title}</h3>
                        </a>
                        <h6>Price: <span class="product-price">${posts[count].price}</span></h6>
                        <button class="btn btn-add-to-cart">Add to Cart</button>
                    </div>
                </div>
			`
			count = count + 1;
		}
	})
	.catch((error) => {
		console.log(error);
	})
}