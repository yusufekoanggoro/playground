
function main(){
    const word = "halo halo bandung ibu kota periangan";

    const split = word.split(' ');

    const counter = {};

    for(let i = 0; i < split.length; i++){
        if(counter[split[i]]){
            counter[split[i]]++;
        }else{
            counter[split[i]] = 1;
        }
    }

    console.log(counter);
}

main();