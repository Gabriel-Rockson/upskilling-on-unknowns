// when 
//
const o = {
  a: 1,
  b: 2,
  __proto__: {
    b: 3,
    c: 4,
    __proto__: {
      d: 5,
    }
  }
}

console.log(o.a)
console.log(o.b)
console.log(o.c)
console.log(o.d)
