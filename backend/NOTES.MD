# Golang handbook

In Go, similar to c++, referencing and dereferencing of pointers are handled with `*` and `&` keywords.

```go
// Standard way of declaring a variable
let n int = 10

// Standard way of declaring a pointer
// Here & is used to reference the memory address of the n variable
let p *int = &n

// This returns the actual memory address
log.Println(p)

// This dereferences the pointer and logs the value
log.Println(*p)
```
