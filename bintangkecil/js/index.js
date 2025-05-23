function main(){
    let str = '';
    for(let i = 0; i < 5; i++){
        for(let j = 0; j < i; j++){
            str += '*';
        }
        str += '\n';
    }
    console.log(str);

    let str2 = '';
    for(let i = 0; i < 5; i++){
        for(let j = 5; j > i; j--){
            str2 += '*';
        }
        str2 += '\n';
    }
    console.log(str2);
}

main();