// very similar to can_sum.js i didint do the naive solution here...
var howSum = function (targetSum, numbers, bad) {
    if (bad === void 0) { bad = new Set(); }
    if (!numbers)
        return null;
    if (targetSum === 0)
        return [];
    if (targetSum < 0)
        return null;
    for (var _i = 0, numbers_1 = numbers; _i < numbers_1.length; _i++) {
        var num = numbers_1[_i];
        if (bad.has(targetSum - num))
            return null;
        var res = howSum(targetSum - num, numbers, bad);
        if (res != null) {
            res.push(num);
            return res; // [ ...res, num ]; - could be used spread syntax but its creating new array
        }
        bad.add(targetSum);
    }
    return null;
};
console.log(howSum(7, [2, 3]));
console.log(howSum(7, [5, 3, 4, 7]));
console.log(howSum(8, [2, 3, 5]));
console.log(howSum(300, [7, 14]));
