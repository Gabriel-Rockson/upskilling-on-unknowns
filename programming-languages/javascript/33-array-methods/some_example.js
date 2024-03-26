const food = require("./food")

// The some method will go through each array and check if at least 
// one element meets the condition that the some method returns
// When an element meets the condition, some exists.
// The return of some is true or false

const hasAFavorite = food.some((foodItem) => foodItem.isFavorite)

console.log(hasAFavorite)
