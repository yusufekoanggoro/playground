
function selectionSort(data, sort){
    if(sort === 'asc'){
        // [4, 5, 1, 10, 2, 3]
        for(let i = 0; i < data.length; i++){
            let minIndex = i;
            for(let j = i+1; j < data.length; j++){
                if(data[j] < data[minIndex]){ // asc
                    minIndex = j;
                }
            }
            let temp = data[i];
            data[i] = data[minIndex];
            data[minIndex] = temp;
        }

        return data;
    }

    if(sort === 'desc'){
        // [4, 5, 1, 10, 2, 3]
        for(let i = 0; i < data.length; i++){
            let minIndex = i;
            for(let j = i+1; j < data.length; j++){
                if(data[j] > data[minIndex]){ // desc
                    minIndex = j;
                }
            }
            let temp = data[i];
            data[i] = data[minIndex];
            data[minIndex] = temp;
        }

        return data;
    }
}

function selectionSortObject(data, sort){
    if(sort === 'asc'){
        // [4, 5, 1, 10, 2, 3]
        for(let i = 0; i < data.length; i++){
            let minIndex = i;
            for(let j = i+1; j < data.length; j++){
                if(data[j].age < data[minIndex].age){ // asc
                    minIndex = j;
                }
            }
            let temp = data[i];
            data[i] = data[minIndex];
            data[minIndex] = temp;
        }

        return data;
    }

    if(sort === 'desc'){
        // [4, 5, 1, 10, 2, 3]
        for(let i = 0; i < data.length; i++){
            let maxIndex = i;
            for(let j = i+1; j < data.length; j++){
                if(data[j].age > data[maxIndex].age){ // desc
                    maxIndex = j;
                }
            }
            let temp = data[i];
            data[i] = data[maxIndex];
            data[maxIndex] = temp;
        }

        return data;
    }
}


function main(){
    console.log("asc: ", selectionSort([4, 5, 1, 10, 2, 3], 'asc'));
    console.log("desc: ", selectionSort([4, 5, 1, 10, 2, 3], 'desc'));

    const data = [
        {
            name: "Yusuf 1",
            age: 40
        },
        {
            name: "Yusuf 2",
            age: 20
        },
        {
            name: "Yusuf 3",
            age: 23
        },
        {
            name: "Yusuf 4",
            age: 12
        },
    ];

    console.log("asc: ", selectionSortObject(data, 'asc'));
    console.log("desc: ", selectionSortObject(data, 'desc'));
}

main();