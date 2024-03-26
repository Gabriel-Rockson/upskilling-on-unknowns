// copyWithin modifies the original array
// the first parameter is where the copied elements should start being filled in
// so for example, copyWithin(3) -> this means, start filling in the copied elements 
// from index 3
// copyWithin takes 3 parameters, targetIndex, startIndex, endIndex
// targetIndex -> where we should start filling after copying
// startIndex -> which index should the copying start from ( default 0 )
// endIndex -> which index should the copying end on (default length of array - targetIndex ), not inclusive

const numbers = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]

numbers.copyWithin(2)

console.log({ numbers })

const alphabets = ["a", "b", "c", "d", "e", "f", "g", "h"]

alphabets.copyWithin(3, 2, 6)

console.log({ alphabets })
