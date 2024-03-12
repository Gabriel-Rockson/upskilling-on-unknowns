// Generic interfaces that describe methods

interface Collection<T> {
  getFirstItem(): T;
  add(object: T): void;
  remove(object: T): void;
}

class List<T> implements Collection<T> {
  private items: T[] = [];

  add(object: T): void {
    this.items.push(object);
  }

  remove(object: T): void {
    let index = this.items.indexOf(object);
 
    if (index > -1) {
      this.items.splice(index, 1)
    }
  }

  getFirstItem(): T {
    return this.items[0]
  }
}

interface Person {
  name: string;
  age: number;
}

const persons = new List<Person>()

persons.add({name: "Gabriel Rockson", age: 12})
