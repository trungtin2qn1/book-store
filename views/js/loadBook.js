

const loadBooks = async function () {
    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/1/books');
    const myJson = await response.json(); //extract JSON from the http response
    return myJson;
}

function load() {
	var count = 0;
	var div = document.getElementById('all-books');
    loadBooks()
        .then((books) => {
            for (var i = 0; i < books.length; i++) {
                div.innerHTML += `
                    <div class="col-lg-2 col-md-3 col-sm-4">
                            <div class="item">
                                <a href="product-single.html" onclick="bd('${books[count].id}')"><img class="product-image" src="${books[count].images}" alt="img">
                                    <h3 class="product-title">${books[count].title}</h3>
                                </a>
                                <h6>Price: <span class="product-price">${books[count].price}</span></h6>
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

function bd(id){
	setCache('id',id);
	console.log('id: ',id);
}

function getCache(key) {
    console.log(localStorage.getItem(key))
    return localStorage.getItem(key)
}
function deleteCache(key) {
    localStorage.removeItem(key)
    return true
}

function setCache(key, value) {
    localStorage.setItem(key, value)
    return true
}


window.onload = load;

