// this is the first problem where the memo set makes really sense
// in the last two it was nice but not necessary we could have only 2 variables
// here we have to store all the paths we already checked...


const canSum = (targetSum: number, numbers: number[]): boolean => {
  if (targetSum === 0) return true;
  if (targetSum < 0) return false;

  for (const num of numbers) {
    const remainder = targetSum - num;
    if (canSum(remainder, numbers)) return true;
  }

  return false;
}

console.log(canSum(7, [2, 3])) // true
console.log(canSum(7, [5, 3, 4, 7])) // true
console.log(canSum(7, [2, 4])) // false
console.log(canSum(8, [2, 3, 5])) // true


const canSumMemo = (targetSum: number, numbers: number[], memo: Record<number, boolean> = {}): boolean => {
  if (targetSum === 0) return true;
  if (targetSum < 0) return false;

  for (const num of numbers) {
    const remainder = targetSum - num;
    if (remainder in memo) return memo[remainder];

    const result = canSumMemo(remainder, numbers, memo);
    memo[remainder] = result;
    if (result) return true;
  }

  return false;
}

console.log(canSumMemo(7, [2, 3])) // true
console.log(canSumMemo(7, [5, 3, 4, 7])) // true
console.log(canSumMemo(7, [2, 4])) // false
console.log(canSumMemo(8, [2, 3, 5])) // true
console.log(canSumMemo(300, [7, 14])) // false

