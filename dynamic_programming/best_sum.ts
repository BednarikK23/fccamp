// return an array containinng the 
// shortest combination of numbers that add up to exactly the targetSum


// again start with naive brute force
const bestSum2 = (targetSum: number, numbers: number[]):  number[] | null => {
    if (targetSum === 0) return [];
    if (targetSum < 0) return null;

    let shorterstCombination: number[] | null = null;

    for (let num of numbers) {
        const reminder = targetSum - num;

        const remainderCombination = bestSum2(reminder, numbers);
        if (remainderCombination !== null) {

            const combination = [...remainderCombination, num];
            if (shorterstCombination === null || combination.length < shorterstCombination.length) {
                shorterstCombination = combination;
            }
        }
    }

    return shorterstCombination;
}

console.log(bestSum2(7, [2, 3]));
console.log(bestSum2(7, [5, 3, 4, 7]));
console.log(bestSum2(8, [2, 3, 5]));



const bestSum = (targetSum: number, numbers: number[], memo: Record<number, (number[] | null)> = {}): number[] | null => {
    if (targetSum in memo) return memo[targetSum];
    if (targetSum === 0) return [];
    if (targetSum < 0) return null;

    let shorterstCombination: number[] | null = null;

    for (let num of numbers) {
        const reminder = targetSum - num;

        const remainderCombination = bestSum(reminder, numbers, memo);
        if (remainderCombination !== null) {

            const combination = [...remainderCombination, num];
            if (shorterstCombination === null || combination.length < shorterstCombination.length) {
                shorterstCombination = combination;
            }
        }
    }

    memo[targetSum] = shorterstCombination;
    return shorterstCombination;
}


// time: O(m^2 * n)
console.log(bestSum(7, [2, 3]));
console.log(bestSum(7, [5, 3, 4, 7]));
console.log(bestSum(8, [2, 3, 5]));
console.log(bestSum(100, [1, 2, 5, 25]));
