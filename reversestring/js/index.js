function main(str){
    const split = str.split("");

    let result = '';
    for(let i = split.length - 1; i >= 0; i--){
        result += split[i];
    }

    console.log(result);
}

main("yusuf eko anggoro");