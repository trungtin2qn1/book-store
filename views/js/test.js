const TestgetAllBooks = async function(input){
    var res = await getAllBooks('5cc967ab2daf5e525551304d');
    console.log("res:", res);
    console.log("res[0]:", res[0]);
    console.log("res[1]:", res[1]);
}