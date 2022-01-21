'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');

let inputString = '';
let currentLine = 0;

process.stdin.on('data', function(inputStdin) {
    inputString += inputStdin;
});

process.stdin.on('end', function() {
    inputString = inputString.split('\n');

    main();
});

function readLine() {
    return inputString[currentLine++];
}

/*
 * Complete the 'insertionSort1' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. INTEGER_ARRAY arr
 */

function insertionSort1(n, arr) {
    // Write your code here
    
    let targetNum = arr[n-1];
    
    for (let i = n-2; i >= 0; i--){
        if( targetNum >= arr[i]){
            arr[i+1] = targetNum;
            console.log(...arr)
            break
        }else{
            if(i == 0){
                arr[i+1] = arr[i]
                console.log(...arr)
                arr[i] = targetNum;
                console.log(...arr)
            }else{
                arr[i+1] = arr[i]
                console.log(...arr)
            }
            
        }
    }
}

function main() {
    const n = parseInt(readLine().trim(), 10);

    const arr = readLine().replace(/\s+$/g, '').split(' ').map(arrTemp => parseInt(arrTemp, 10));

    insertionSort1(n, arr);
}
