// Generic interfaces that describe object properties

interface Pair<K, V> {
  key: K;
  value: V;
}

let month: Pair<string, number> = {
  key: 'January',
  value: 1
}

console.log(month)
console.log(month.value)
