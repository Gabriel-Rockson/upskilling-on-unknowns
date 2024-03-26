const food = require("./food")

// push always modifies the original array, use with caution
// push will return the new length after adding the element to the end of the array


console.log("Old food: ", food)

food.push({
  name: "Kokonte",
  type: "swallow",
  isFavorite: false,
  dislike: true
})

console.log("New food: ", food)
