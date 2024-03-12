const boxes = [
  {value: 1, getValue() { return this.value }},
  {value: 2, getValue() { return this.value }},
  {value: 3, getValue() { return this.value }},
  {value: 4, getValue() { return this.value }},
]


// The above can be done a below this comment
// All the box objects will now have just one reference to getValue, reducing memory usage

const boxPrototype = {
  getValue() {
    return this.value
  }
}

const newBoxes = [
  {value: 1, __proto__: boxPrototype },
  {value: 2, __proto__: boxPrototype },
  {value: 3, __proto__: boxPrototype },
  {value: 4, __proto__: boxPrototype },
]

console.log(newBoxes[0].getValue === newBoxes[1].getValue)


// Using constructor function
function Box(value) {
  this.value = value
}


Box.prototype.getValue = function() {
  return this.value
}

Box.prototype.getValuePlus2 = function() {
  return this.value + 2
}

Box.prototype.getValue = function() {
  return this.value * 0.82
}

const otherBoxes = [new Box(1), new Box(2), new Box(3), new Box(4)]

console.log(otherBoxes[0].getValue === otherBoxes[3].getValue)
console.log(otherBoxes[3].getValuePlus2())
console.log(otherBoxes[3].getValue())
console.log(Object.getPrototypeOf(new Box()))
console.log(Box.prototype.constructor)


// Use a class

class NewBox {
  constructor(value) {
    this.value = value;
  }

  // methods that are created on Box.prototype
  getValue() {
    return this.value;
  }
}
