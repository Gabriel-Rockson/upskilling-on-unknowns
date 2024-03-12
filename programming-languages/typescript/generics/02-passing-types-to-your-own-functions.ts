const makeFetch = async <TData>(url: string): Promise<TData> => {
  return fetch(url)
    .then((res) => {
      return res.json()
    })
}

makeFetch<{ firstName: string; lastName: string }>("/some/endpoint")
  .then((res) => {
    console.log(res);
  })
