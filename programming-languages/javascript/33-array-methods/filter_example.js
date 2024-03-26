const food = require("./food")

// Filter array method takes a callback and returns a condition
// The array is filtered based on this returned condition

const favoriteFoods = food.filter((foodItem) => foodItem.isFavorite)

const chewingFoods = food.filter((foodItem) => foodItem.type == "chewing")

console.log("Favorites: ", favoriteFoods)

console.log("Chewing Foods: ", chewingFoods)
