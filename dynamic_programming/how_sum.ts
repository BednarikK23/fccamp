// very similar to can_sum.js i didint do the naive solution here...

const howSum = (targetSum: number, numbers: number[], bad: Set<number> = new Set()): number[] | null => {
    if (!numbers) return null;
    if (targetSum === 0) return [];
    if (targetSum < 0) return null;

    for (let num of numbers) {
        if (bad.has(targetSum - num)) return null;

        const res = howSum(targetSum - num, numbers, bad)
        if (res != null) {
            res.push(num);
            return res; // [ ...res, num ]; - could be used spread syntax but its creating new array
        }
        bad.add(targetSum);
    }

    return null;
}


console.log(howSum(7, [2, 3]));
console.log(howSum(7, [5, 3, 4, 7]));
console.log(howSum(8, [2, 3, 5]));
console.log(howSum(300, [7, 14]));
