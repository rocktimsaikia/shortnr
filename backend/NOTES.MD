# Golang handbook

In Go, similar to c++, referencing and dereferencing of pointers are handled with `*` and `&` keywords.

```go
// Standard way of declaring a variable
let n int = 10

// Standard way of declaring a pointer
// Here & is used to reference the memory address of the n variable
let p *int = &n

// Or the above can be done with 'short variable declaration operator' (:=)
// Using this operator, Go infers the type of the variable implicitly
n := 10
p := &n

// This returns the actual memory address
log.Println(p)

// This dereferences the pointer and logs the value
log.Println(*p)
```

Writing conditional `if-else` statements in Go similar any other languages. Looks like a mix of Python and Javascript.

```go
if score > 80 {
    log.Println("Best")
} else if score > 50 {
    log.Println("Moderate")
} else {
    log.Println("Not Good!")
}
```

> !NOTE
> `nil` is used for nullish values in Go istead of `null` or `None`.

We can also write one line `if` statements with "short statement":
```go
if err := doSomething(); err != nil {
    // Handle error
}
```
