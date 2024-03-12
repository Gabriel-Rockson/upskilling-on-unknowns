const set = new Set<number>();

set.add(1)
set.add(4)

// This error is meant to happen
set.add("abc")

console.log(set)
