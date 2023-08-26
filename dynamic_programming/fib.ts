// lets start with typical fibonacci problem
// we will use memo to get compfortable with it
// eventhough we dont need it here, we could use only 2 variables...

// slow solution because of recursion, we re doing a lot of repeated work
// time complexity O(2^n)
const fib1 = (n: number): number => {
    if (n <= 2) return 1;
    return fib1(n - 1) + fib1(n - 2);
}

console.log(fib1(6)); // 8
console.log(fib1(7)); // 13
console.log(fib1(8)); // 21


// memoization solution - storing duplicate subproblems to use them later
// we dont have to go deeper to recursion just look up in memo
//      js object, keys will be arg to fn and value will be return value
// time complexity O(n)
// space complexity O(n)
const fib = (n: number, memo: Record<number, number> = {}): number => {
    if (n in memo) return memo[n];
    if (n <= 2) return 1;
    memo[n] = fib(n - 1, memo) + fib(n - 2, memo);
    return memo[n];
}

console.log(fib(6)); // 8
console.log(fib(7)); // 13
console.log(fib(8)); // 21
console.log(fib(50)); // 12586269025
