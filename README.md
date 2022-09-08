# go-generics-workshop

This repository holds the code written during [Go generics O'Reilly workshop](https://learning.oreilly.com/live-events/go-generics-in-2-hours/0636920072269/0636920075658/) by Johny Boursiquot.

## Introduction
Go 1.18 introduces the ability to write generic behavior where you declare and use functions or types that are written 
to work with any of a set of types provided when you invoke them. Go developers can remove complexity from their code and 
avoid the risks associated with using the empty interface (interface{}) or reflection to achieve what Go now offers out of the box with generics.

Generics in Go brings a much-requested update to the language but also the need to understand how the type system has evolved to support the capability. 
This training will introduce you to the new syntax you must understand before using generics and important concepts such as type parameters, 
type constraints, type approximations, built-in constraints, and more.

## What you’ll learn and how you can apply it

By the end of this live, hands-on online course, you’ll understand:

- How to use type parameters
- How to declare and use type constraints
- How to invoke type-constrained functions
- How to use the built-in constraints the standard library makes available

And you’ll be able to:

- Refactor code that once used interface{} and reflection to be simpler with generics
- Explore slices, maps, and channels, whichwhat generics now make possible with other areas of the language, including slices, maps, and channels

---
## This live event is for you because…

- You’re an existing user of Go or are at least comfortable with it
- You want to understand what generics bring to the table
- You want to know when and when not to use generics

### Prerequisites
Intermediate understanding of Go syntax

### Recommended preparation:
Read [Learning Go: Chapter 15. A Look at the Future: Generics in Go](https://learning.oreilly.com/library/view/learning-go/9781492077206/ch15.html) (book)

### Recommended follow-up:
Watch [Robert Griesemer and Ian Lance Taylor: Generics AMA](https://learning.oreilly.com/videos/go-generics-extensibility/0636920551942/0636920551942-video333639/) (video)

---
## Schedule
The timeframes are only estimates and may vary according to how the class is progressing.

### Getting to know Go Generics (60 minutes)

- Presentation: Introducing type parameters
- Presentation: Introducing type constraints
- Exercise: Writing generic functions
- Q&A: 5 minutes
- Break (5 minutes)

### Working with Constraints (45 minutes)

- Presentation: Custom type constraints
- Presentation: Built-in type constraints
- Exercise: Working with type constraints
- Q&A (5 mins)

### What’s Next for Generics (15 minutes)

- Presentation: How language elements like slices, maps, and channels might benefit from generics
Q&A

---

## Resources

- [Go in 3 Weeks Generics](https://github.com/jboursiquot/go-in-3-weeks/wiki/Generics) by Johny Boursiquot