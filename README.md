
#  OpGame

![Math operators](https://upload.wikimedia.org/wikipedia/commons/thumb/1/12/Basic_arithmetic_operators.svg/330px-Basic_arithmetic_operators.svg.png)

## Overview

writing a Go program to solve a simple mathematical puzzle.  The program is to efficiently find the binary operators that complete the equation.  I'm considering the operators are addition, subtraction, multiplication, and division (when the divisor is non-zero and the remainder is zero).  For example, given `1 O 2 = 3` the only operator that works to put in the circle is "+".  So `1 + 2 = 3` solves the problem.

Here i'm finding find two operators that satisfy the equation.  The solution would be `1 + 2 / 3 = 1`. 
e the operators and supply that to your program.

## Implementing
- functions
- Tree searching
- Recursion

## My program example Output

```shell
$ echo "3 1 2" | go run OpGame.go
3-1=2

$ echo "5 4 2 22" | go run OpGame.go
5*4+2=22

$ (echo "3 1 2" && echo "9 2 18") | go run OpGame.go
3-1=2
9*2=18

$ echo "6 2 3 4" | go run OpGame.go
6*2/3=4

$ go run OpGame.go < testdata/basic.txt
3 - 1 = 2

9 + 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6
```
Multiple solutions

$ echo "7 3 3 7" | go run OpGame.go -all
7+3-3=7, 7-3+3=7, 7*3/3=7

$ go run OpGame.go -all < testdata/basic.txt
3 - 1 = 2

9 + 0 = 9, 9 - 0 = 9
2 * 3 = 6
4 + 5 = 9
4 + 5 + 6 = 15
2 + 3 + 1 = 6, 2 * 3 * 1 = 6, 2 * 3 / 1 = 6

