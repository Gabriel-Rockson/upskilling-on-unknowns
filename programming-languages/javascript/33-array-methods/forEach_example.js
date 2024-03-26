const food = require("./food")

console.log("Food before ", food)

food.forEach((foodItem) => {
  console.log(foodItem.name.toUpperCase())
})

console.log("Food after ", food)
