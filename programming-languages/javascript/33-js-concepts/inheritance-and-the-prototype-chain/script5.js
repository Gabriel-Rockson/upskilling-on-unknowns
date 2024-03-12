function Base() {

}

function Derived() {

}

Object.setPrototypeOf(Derived.prototype, Base.prototype)

const object = new Derived()

console.log(Object.getPrototypeOf(object))
