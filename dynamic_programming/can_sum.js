// this is the first problem where the memo set makes really sense
// in the last two it was nice but not necessary we could have only 2 variables
// here we have to store all the paths we already checked...
var canSum = function (targetSum, numbers) {
    if (targetSum === 0)
        return true;
    if (targetSum < 0)
        return false;
    for (var _i = 0, numbers_1 = numbers; _i < numbers_1.length; _i++) {
        var num = numbers_1[_i];
        var remainder = targetSum - num;
        if (canSum(remainder, numbers))
            return true;
    }
    return false;
};
console.log(canSum(7, [2, 3])); // true
console.log(canSum(7, [5, 3, 4, 7])); // true
console.log(canSum(7, [2, 4])); // false
console.log(canSum(8, [2, 3, 5])); // true
var canSumMemo = function (targetSum, numbers, memo) {
    if (memo === void 0) { memo = {}; }
    if (targetSum === 0)
        return true;
    if (targetSum < 0)
        return false;
    for (var _i = 0, numbers_2 = numbers; _i < numbers_2.length; _i++) {
        var num = numbers_2[_i];
        var remainder = targetSum - num;
        if (remainder in memo)
            return memo[remainder];
        var result = canSumMemo(remainder, numbers, memo);
        memo[remainder] = result;
        if (result)
            return true;
    }
    return false;
};
console.log(canSumMemo(7, [2, 3])); // true
console.log(canSumMemo(7, [5, 3, 4, 7])); // true
console.log(canSumMemo(7, [2, 4])); // false
console.log(canSumMemo(8, [2, 3, 5])); // true
console.log(canSumMemo(300, [7, 14])); // false
