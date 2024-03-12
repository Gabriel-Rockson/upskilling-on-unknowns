const object = { a: 1 }

console.log(Object.getPrototypeOf(object) === Object.prototype)

const array = [1, 2, 3];
console.log(Object.getPrototypeOf(array) === Array.prototype)

console.log(JSON.stringify(Object.prototype))
console.log(JSON.stringify(Array.prototype))

function MyConstructor(value){
  this.value = value
}

MyConstructor.prototype.getValue = function() {
  return this.value;
}

MyConstructor.prototype.setValue = function(newValue) {
  this.value = newValue
}


const a = new MyConstructor(20)
console.log(a.getValue())

console.log(a)

a.setValue(23)
console.log(a.getValue())
