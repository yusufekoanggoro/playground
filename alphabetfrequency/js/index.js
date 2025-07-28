function main(){
    const alphabet = "backend developer".replaceAll(" ", "");

    const split = alphabet.split("")

    const counter = {};

    for (let i = 0; i < split.length; i++){
        if(counter[split[i]]){
            counter[split[i]]++;
        }else {
            counter[split[i]] = 1;
        }
    }

    console.log(counter);

}

main();