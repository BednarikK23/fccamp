// how many way can we travell the grid moving only right or down
// it helped me when i aknowledge that i shrink the grid by one line every move...
// until i reach the end some basecase...
// if we starting with 2,3 grod we can move to 1,3 or 2,2
//    1,3 we can move to 0,3 or 1,2 - 0,3 is out of grid so we can move only to 1,2
//    2,2 we can move to 1,2 or 2,1
//          1,2 we can move to 0,2 or 1,1 - 1,1 is basecase
// its spans like into this tree where left is down and right is right...
// and we can see its just twisted fibonacci problem...


// Naive, bruteforce - same problem as in fib.js - slow on bigones like 18,18
const gridTraveler1 = (m, n) => {
    if (m === 1 && n === 1) return 1;
    if (m === 0 || n === 0) return 0;
    return gridTraveler1(m - 1, n) + gridTraveler1(m, n - 1);
}

console.log(gridTraveler1(1, 1)) // 1
console.log(gridTraveler1(2, 3)) // 3
console.log(gridTraveler1(3, 2)) // 3
console.log(gridTraveler1(3, 3)) // 6
// console.log(gridTraveler(18, 18)) // 2333606220


// now lets use memo
// time complexity O(m*n)
// space complexity O(m+n) 

const gridTraveler = (m, n, memo = {}) => {
    if (m === 1 && n === 1) return 1;
    if (m === 0 || n === 0) return 0;

    const key = m + ',' + n; // separator, because without it: 185 - either 18;5 or 1;85...
    if (key in memo) return memo[key];

    const sndKey = n + ',' + m;
    if (sndKey in memo) return memo[sndKey]; // to have (2, 1) is the same to have (1, 2)   

    memo[key] = gridTraveler(m - 1, n, memo) + gridTraveler(m, n - 1, memo); 

    return memo[key];
}







console.log(gridTraveler(1, 1)) // 1
console.log(gridTraveler(2, 3)) // 3
console.log(gridTraveler(3, 2)) // 3
console.log(gridTraveler(3, 3)) // 6
console.log(gridTraveler(18, 18)) // 2333606220
