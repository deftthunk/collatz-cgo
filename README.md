Collatz Sequence Solver

My first foray into CGO. Solves sequences of the Collatz Conjecture, defined as the following:
- start with any positive integer n
- if n is even, the next term is n/2
- if n is odd, the next term is 3n + 1
- no matter the value of n, the sequence always reaches 1

This CGO version combines the native multithreading of Go with the speed of C. It's memory usage is light but will take as many cores as possible.
