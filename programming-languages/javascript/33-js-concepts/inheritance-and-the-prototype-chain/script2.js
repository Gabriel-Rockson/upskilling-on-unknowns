const parent = {
  value: 2,
  method() {
    return this.value + 1
  }
}

console.log(parent.method())

const firstChild = {
  __proto__: parent,
}

console.log(firstChild.method())

const secondChild = {
  value: 100,
  __proto__: parent,
}

console.log(secondChild.method())
