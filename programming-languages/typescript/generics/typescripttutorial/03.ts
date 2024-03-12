// Generic interfaces that describe index types

interface Options<T> {
  [name: string]: T
}

let inputOptions: Options<boolean> = {
  'disabled': false,
  'visible': true
}
