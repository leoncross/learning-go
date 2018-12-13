# key learnings / takeaways


- functions (func) require declared outputs
- if func's require arguments, you need to specify the type (int, string etc.)



you can write benchmarks in Go (look in iteration for example); to run a benchmark:

```
go test -bench=.
```

you can write example code in your test to document functions (look in integers for example):

```
func ExampleAdd() {
  sum := Add(1, 5)
  fmt.Println(sum)
  // Output: 6
```
