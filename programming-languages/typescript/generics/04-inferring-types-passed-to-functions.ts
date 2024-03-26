interface MyResponse<T> {
  data: T
}

class NewResponse implements MyResponse<{name: string, age: number}> {

}
