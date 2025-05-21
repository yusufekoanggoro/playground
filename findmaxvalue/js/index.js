
function main(){
    const array = [1, 3, 454, 999, 765];

    let largest = array[0];

    for(let i = 0; i < array.length; i++){
        if(array[i] > largest){
            largest = array[i]; 
        }
    }

    console.log("max value in array of int: ", largest);

    const users = [
        {
            name: 'udin 1',
            age: 23
        },
        {
            name: 'udin 2',
            age: 22
        },
        {
            name: 'udin 3',
            age: 50
        },
        {
            name: 'udin 4',
            age: 46
        }
    ];

    let maxAge = users[0].age;

    for(let i = 0; i < users.length; i++){
        if(users[i].age > maxAge){
            maxAge = users[i].age;
        }
    }

    console.log('max value in array of object: ', maxAge);
}

main();