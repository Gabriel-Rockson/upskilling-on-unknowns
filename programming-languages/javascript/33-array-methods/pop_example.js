const numbers = [1, 3, 5, 7, 11, 13]

// pop removes the last element from the array and returns it, if array is emepty, it returns undefined
// pop modifies the original array

console.log("Before removing: ", { numbers })

const removedNumber = numbers.pop()

console.log({ removedNumber })

console.log("After removing: ", { numbers })
