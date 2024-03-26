const food = require("./food")

// findIndex is just like find, the only difference is that findIndex
// returns the index of the first element that meets the condition 
// the callback returns

const dislikedFoodIndex = food.findIndex((foodItem) => foodItem.dislike)

console.log(dislikedFoodIndex)
