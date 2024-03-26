const food = require("./food")

// Find will return the first element that meets the condition the callback returns

const favoriteFood = food.find((foodItem) => foodItem.isFavorite)

console.log(favoriteFood)
