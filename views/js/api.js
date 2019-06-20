const getAllBooks = async function (input) {
    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + input + '/books', {
        method: 'GET',
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    // var object = await JSON.parse(JSON.stringify(myJson));
    // console.log(object)

    return [true, myJson]
}
//Return:
//[
//     {
//         "id": "5cfe859b2daf5e6ca2ffbfa2",
//         "title": "Đắc Nhân Tâm",
//         "description": "Tại sao Đắc Nhân Tâm luôn trong Top sách bán chạy nhất thế giới? Bởi vì đó là cuốn sách mọi người đều nên đọc.",
//         "organization": "First News - Trí Việt",
//         "author": "Dale Carnegie",
//         "inventory": 100,
//         "from_time": 1234,
//         "to_time": 2356,
//         "price": 76000,
//         "discount_price": 53200,
//         "category_name": "",
//         "images": [
//             "https://firebasestorage.googleapis.com/v0/b/book-store-5f677.appspot.com/o/dac_nhan_tam.gif?alt=media&token=b66e0b99-f23d-49f5-b54f-6dac25c0b156"
//         ],
//         "comments": []
//     },
//     {
//         "id": "5cfe85f42daf5e6ca2ffbfa3",
//         "title": "Sống Thực Tế Giữa Đời Thực Dụng",
//         "description": "THỰC DỤNG Ư? KHÔNG HỀ, TÔI CHỈ RẤT THỰC TẾ THÔI!",
//         "organization": "Vibooks",
//         "author": "Mễ Mông",
//         "inventory": 100,
//         "from_time": 1234,
//         "to_time": 2356,
//         "price": 98000,
//         "discount_price": 70000,
//         "category_name": "",
//         "images": [
//             "https://firebasestorage.googleapis.com/v0/b/book-store-5f677.appspot.com/o/song_thuc_dung.jpg?alt=media&token=7ab82c32-086f-4dbb-adae-b94412ac4b3f"
//         ],
//         "comments": []
//     }
// ]

//Input:
//email: string
//password: string
const register = async function (email, password) {
    //Truyền vào email và password
    //email password

    var myBody = {
        email: email,
        password: password,
    };

    // console.log(myBody);

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

    return [true, myJson]
}

const login = async function (email, password) {
    //Truyền vào email và password
    // email: string
    // password: string

    var myBody = {
        email: email,
        password: password,
    };

    console.log(myBody);

    const response = await fetch('http://127.0.0.1:2013/api/sign_in', {
        method: 'POST',
        body: JSON.stringify(myBody), // string or object
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Return:
// {
//     "id": "5d00d4862daf5e12a352312a",
//     "email": "test6@gmail.com",
//     "first_name": "",
//     "last_name": "",
//     "phone": "",
//     "username": "",
//     "avatar": "",
//     "carts": []
// }
//Input:
//"input": customer id (string)
const getCustomerInfo = async function (input) {
    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + input, {
        method: 'GET',
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
//{
//     "id": "5cc967ab2daf5e525551304d",
//     "email": "tin@gmail.com",
//     "first_name": "Harry",
//     "last_name": "Tin Huynh",
//     "phone": "",
//     "username": "",
//     "avatar": "",
//     "carts": []
// }

//Input: 
//"id": string
//"info": customer
const updateCustomerInfo = async function (id, info) {

    const temp = await getCustomerInfo(id);

    var myBody = {
        first_name: info.first_name,
        last_name: info.last_name,
        phone: info.phone,
        username: info.username,
        avatar: info.avatar,
        carts: temp.carts,
    };

    console.log(myBody);

    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + id, {
        method: 'PUT',
        body: JSON.stringify(myBody), // string or object
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
//{
//     "id": "5cc967ab2daf5e525551304d",
//     "email": "tin@gmail.com",
//     "first_name": "first_name",
//     "last_name": "last_name",
//     "phone": "phone",
//     "username": "username",
//     "avatar": "avatar",
//     "carts": []
// }


//Input:
//"customer_id": string
//"book_id": string 
const getBookInfo = async function (customer_id, book_id) {
    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + customer_id + '/book/' + book_id, {
        method: 'GET',
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
//{
//     "id": "5cfcc7152daf5e2226a62eb6",
//     "title": "title",
//     "description": "description",
//     "organization": "organization",
//     "author": "",
//     "inventory": 100,
//     "from_time": 1234,
//     "to_time": 2356,
//     "price": 15,
//     "discount_price": 10,
//     "category_name": "",
//     "images": [],
//     "comments": []
// }

//Input:
//"book_id": id of book
//"amount": amount of book want to buy
//"user_id": id of customer
const addBookToCart = async function (user_id, amount, book_id) {

    var myBody = {
        book_id: book_id,
        amount: amount
    };

    console.log(myBody);

    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + user_id + '/cart', {
        method: 'POST',
        body: JSON.stringify(myBody), // string or object
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
// "amount": 2,
// "book_id": "5cc967ab2daf5e525551304d",
// "id": "5d072b372daf5e6f63a24f46"

//Input:
//"cart_id": id of cart (each books in cart will have 1 id)
//"amount": amount of book want to buy
//"user_id": id of customer
const updateCartAmount = async function (user_id, cart_id, amount) {

    var myBody = {
        amount: amount
    };

    console.log(myBody);

    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + user_id + '/cart/' + cart_id, {
        method: 'PUT',
        body: JSON.stringify(myBody), // string or object
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
// "amount": 2,
// "book_id": "5cc967ab2daf5e525551304d",
// "id": "5d072b372daf5e6f63a24f46"

//Input:
// "customer_id": id of user
const getAllCartsInfo = async function (customer_id){
    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + customer_id + '/carts', {
        method: 'GET',
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
// [
//     {
//         "id": "5d072e5e2daf5e3fe44f4e50",
//         "amount": 3,
//         "book_id": "5cfcc7152daf5e2226a62eb6",
//         "title": "title",
//         "description": "description",
//         "price": 15,
//         "discount_price": 10
//     }
// ]

//Input:
//"customer_id": id of customer
const checkOut = async function (customer_id){
    var myBody = {
        note: "abc"
    };

    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + customer_id + '/order', {
        method: 'POST',
        body: JSON.stringify(myBody), // string or object
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
// "book_id": ["5cfe83312daf5e6ca2ffbfa0", ...]
// "customer_id": "5cc967ab2daf5e525551304d"
// "id": "5d073c3d2daf5e7a9ef94a60"
// "note": "abc"
// "order_date": 1560755261
// status: 0

//Input:
//"customer_id": id of customer
//"cart_id": id of cart
const deleteCart = async function (customer_id, cart_id){

    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + customer_id + '/cart/' + cart_id, {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
//"message": "Success"

//Input:
// "customer_id": id of user
const getAllOrdersInfo = async function (customer_id){
    const response = await fetch('http://127.0.0.1:2013/api/auth/customer/' + customer_id + '/orders', {
        method: 'GET',
    });
    const myJson = await response.json(); //extract JSON from the http response
    // do something with myJson
    console.log(myJson)

    return [true, myJson]
}
//Output:
// [
//     {
//         "id": "5cfe74e12daf5e3a67545b6f",
//         "note": "Note",
//         "order_date": 1560179937,
//         "shipping_fee": 10000,
//         "status": 0,
//         "book_id": [
//             "5cfcc7152daf5e2226a62eb6"
//         ],
//         "payment": {
//             "id": "",
//             "method": ""
//         },
//         "customer_id": "5cc967ab2daf5e525551304d"
//     },
//     {
//         "id": "5cfe75762daf5e532903c102",
//         "note": "Note",
//         "order_date": 1560180086,
//         "shipping_fee": 10000,
//         "status": 0,
//         "book_id": [
//             "5cfcc7152daf5e2226a62eb6"
//         ],
//         "payment": {
//             "id": "",
//             "method": ""
//         },
//         "customer_id": "5cc967ab2daf5e525551304d"
//     },
//     {
//         "id": "5d073c3d2daf5e7a9ef94a60",
//         "note": "abc",
//         "order_date": 1560755261,
//         "shipping_fee": 10000,
//         "status": 0,
//         "book_id": [
//             "5cfe83312daf5e6ca2ffbfa0"
//         ],
//         "payment": {
//             "id": "",
//             "method": ""
//         },
//         "customer_id": "5cc967ab2daf5e525551304d"
//     }
// ]

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

