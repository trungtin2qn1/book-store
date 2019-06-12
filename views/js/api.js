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

