type MyGenericType<TData> = {
  data: TData;
}

type Example1 = MyGenericType<{
  firstName: string;
}>
