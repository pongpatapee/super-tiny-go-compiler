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

The compiler essentially comprises of 3 main steps:
1. *Parsing* 

Parsing takes the raw strings of the code and parses it into a more abstract representation. E.g., an Abstract Syntax Tree (AST)

2. *Transformation*

The abstract representation of the code is then *transformed* into a modified abstract representation of the output we want to generate. In our example of turning LISP function calls into C function calls. We would take the AST representing a LISP function call and *transform*/*modify* the representation into an AST of a C function call. 

3. *Code Generation* 

Lastly, we use the final abstract representation and generate new code based on that representation. A traditional compiler would generate assembly or machine code from the abstract representation. 

### Parsing

Parsing typically gets broken down into 2 main steps: *Lexical Analysis* and *Syntactic Analysis*

1. *Lexical Analysis*

Lexical analysis is the process of taking the raw code, splitting it apart, and assigning meaning to isolated pieces. This is typically done through something called a *Tokenizer* or a *Lexer*

The *Tokenizer* generates an array of objects called tokens. These tokens describe the isolated piece of syntax. This could be a number, label, punctuation, operator, etc. 

2. *Syntactic Analysis*

The tokens are then used to build an abstract representation of the code called an Abstract Syntax Tree (AST). The AST holds more information than the tokens and is able to describe the syntax and their relationship between one another. 

An AST is a deeply nested object that is easy to work with and helps describe the information as mentioned above. 

A token might look like:
```
[
  { type: 'paren',  value: '('        },
  { type: 'name',   value: 'add'      },
  { type: 'number', value: '2'        },
  { type: 'paren',  value: '('        },
  { type: 'name',   value: 'subtract' },
  { type: 'number', value: '4'        },
  { type: 'number', value: '2'        },
  { type: 'paren',  value: ')'        },
  { type: 'paren',  value: ')'        }
]
```

An AST might look like:

```
 {
   type: 'Program',
   body: [{
     type: 'CallExpression',
     name: 'add',
     params: [{
       type: 'NumberLiteral',
       value: '2'
     }, {
       type: 'CallExpression',
       name: 'subtract',
       params: [{
         type: 'NumberLiteral',
         value: '4'
       }, {
         type: 'NumberLiteral',
         value: '2'
       }]
     }]
   }]
 }
```

### Transformation
### Code Generation




