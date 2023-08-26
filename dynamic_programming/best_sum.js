// return an array containinng the 
// shortest combination of numbers that add up to exactly the targetSum
var __spreadArray = (this && this.__spreadArray) || function (to, from, pack) {
    if (pack || arguments.length === 2) for (var i = 0, l = from.length, ar; i < l; i++) {
        if (ar || !(i in from)) {
            if (!ar) ar = Array.prototype.slice.call(from, 0, i);
            ar[i] = from[i];
        }
    }
    return to.concat(ar || Array.prototype.slice.call(from));
};
// again start with naive brute force
var bestSum2 = function (targetSum, numbers) {
    if (targetSum === 0)
        return [];
    if (targetSum < 0)
        return null;
    var shorterstCombination = null;
    for (var _i = 0, numbers_1 = numbers; _i < numbers_1.length; _i++) {
        var num = numbers_1[_i];
        var reminder = targetSum - num;
        var remainderCombination = bestSum2(reminder, numbers);
        if (remainderCombination !== null) {
            var combination = __spreadArray(__spreadArray([], remainderCombination, true), [num], false);
            if (shorterstCombination === null || combination.length < shorterstCombination.length) {
                shorterstCombination = combination;
            }
        }
    }
    return shorterstCombination;
};
console.log(bestSum2(7, [2, 3]));
console.log(bestSum2(7, [5, 3, 4, 7]));
console.log(bestSum2(8, [2, 3, 5]));
var bestSum = function (targetSum, numbers, memo) {
    if (memo === void 0) { memo = {}; }
    if (targetSum in memo)
        return memo[targetSum];
    if (targetSum === 0)
        return [];
    if (targetSum < 0)
        return null;
    var shorterstCombination = null;
    for (var _i = 0, numbers_2 = numbers; _i < numbers_2.length; _i++) {
        var num = numbers_2[_i];
        var reminder = targetSum - num;
        var remainderCombination = bestSum(reminder, numbers, memo);
        if (remainderCombination !== null) {
            var combination = __spreadArray(__spreadArray([], remainderCombination, true), [num], false);
            if (shorterstCombination === null || combination.length < shorterstCombination.length) {
                shorterstCombination = combination;
            }
        }
    }
    memo[targetSum] = shorterstCombination;
    return shorterstCombination;
};
// time: O(m^2 * n)
console.log(bestSum(7, [2, 3]));
console.log(bestSum(7, [5, 3, 4, 7]));
console.log(bestSum(8, [2, 3, 5]));
console.log(bestSum(100, [1, 2, 5, 25]));
