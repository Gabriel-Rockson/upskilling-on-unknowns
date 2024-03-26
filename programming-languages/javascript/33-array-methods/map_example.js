const food = require("./food")

// map returns a new array
// this new array can be stored in another variable

const newFood = food.map((foodItem) => {
  
  return foodItem.isFavorite 
  ? foodItem.name.toUpperCase() + " is a Favorite" 
  : foodItem.name.toUpperCase()
})

console.log(newFood)
