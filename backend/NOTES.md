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

Writing conditional `if-else` statements in Go is not similar to any other languages. Looks like a mix of Python and Javascript.

```go
if score > 80 {
    log.Println("Best")
} else if score > 50 {
    log.Println("Moderate")
} else {
    log.Println("Not Good!")
}
```

> [!NOTE]
> `nil` is used for nullish values in Go istead of `null` or `None`

We can also write one line `if` statements with "short statement"
```go
if err := doSomething(); err != nil {
    // Handle error
}
```


`panic` is a builtin function in Go to singal that something went wrong and the program can not continue normal execution anymore. It is a combination of logger (optinal) + stack trace + exit 1

```go
panic("Yo this is an unexpected error!")

# The above statement will:
# 1. print the passed string
# 2. show the stack-trace of the panic invocation
# 3. stop/exit the current running program
```
