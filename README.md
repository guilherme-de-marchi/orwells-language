# orwells-language
Programming language for the Orwell's Blockchain

## Keywords and notations

```
var     // Variable declaration

if      // Conditional declaration

exec    // Execute some inherit function

=       // Assignment sign

==      // Equal comparator

+       // Sum sign

-       // Subtraction sign

;       // Instruction delimiter

$<ref>  // Reference to the value of the variable <ref>

<str>   // String literal

<int>   // Integer literal

<float> // Float literal
```

## Valid syntaxes


```<ANY>``` means a value or keyword of type ```ANY```

```<ANY>|<ANY>|...``` means a value or keyword of type ```ANY OR ANY OR ...```


#### Variable declaration
```
// LITERAL DECLARATION
var <ref> = <ref>|<str>|<int>|<float> ;

// BINARY DECLARATION
var <ref> = <ref>|<str>|<int>|<float> +|- <ref>|<str>|<int>|<float> ;
```

#### Conditionals
```
// CONDITIONAL DECLARATION
if <ref>|<str>|<int>|<float> == <ref>|<str>|<int>|<float> ;
```

## How conditionals work?

In this language, conditionals do not work as usual. Here, when a conditional is declared, all of the following statements will only occur if the conditional is valid. This conditional will not be considered when the instruction is its respective ENDIF keyword.

For example, in a language like Go, a chain of conditional blocks is written as follows:

```
if a == b {
    // do something
    
    if b == c {
        // do anything
    }
}
```

However, in Orwell's Language, we do as follows:

```
if a == b ;
// do something
if b == c ;
// do anything
endif ;
endif ;
```

Confused, I know. Imagine the brackets has been replaced by the endif keyword, and each if have its respective endif.
Omitting an endif will result in a compilation error

## Features

### Syntax

- [x] Conditionals
- [x] Variable declaration
- [x] Variable References
- [ ] Execute inherit functions
- [ ] Binary functions (Not all yet)
- [ ] Comments