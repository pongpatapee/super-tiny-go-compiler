# Super Tiny Go Compiler

I have been meaning to learn a little bit about compilers and get a basic understanding of how they work under the hood. In addition, this is a nice project to practice Go.

The original code is from [hazbo](https://github.com/hazbo/the-super-tiny-compiler), whose project is also listed under the [build your own x repo](https://github.com/codecrafters-io/build-your-own-x?tab=readme-ov-file#build-your-own-programming-language).

The resulting compiler will turn lisp-like function call into C-like function calls.
```
LISP                      C
(add 2 2)                 add(2, 2)
(subtract 4 2)            subtract(4, 2)
(add 2 (subtract 4 2))    add(2, subtract(4, 2))

```

## How does a compiler work?



