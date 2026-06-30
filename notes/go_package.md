# Go Package Notes

Go code is organized around packages.

The important rule:

> Files in the same directory with the same `package` name are compiled together as one package.

That means top-level functions, variables, constants, and types in one file can be used by another file in the same package.

## Same Directory, Same Package

Example:

```text
demo/
  a.go
  b.go
```

`a.go`:

```go
package main

var name = "Go"

func hello() string {
	return "hello " + name
}
```

`b.go`:

```go
package main

import "fmt"

func main() {
	fmt.Println(hello())
}
```

This works because `a.go` and `b.go` are in the same directory and both use `package main`.

Run it with:

```powershell
go run ./demo
```

## Top-Level Names Share One Scope

Because files in one package are compiled together, their top-level names share the same package scope.

This is useful:

```go
// a.go
package main

func add(a, b int) int {
	return a + b
}
```

```go
// b.go
package main

func testAdd() int {
	return add(1, 2)
}
```

But it also means names can collide:

```go
// a.go
package main

var nums = []int{1, 2, 3}
```

```go
// b.go
package main

var nums = []int{4, 5, 6}
```

This fails because `nums` is declared twice in the same package.

Functions, structs, interfaces, constants, and variables all follow this rule.

## Top-Level Declarations Use Keywords

At the top level of a Go file, you are in package scope. Package scope only accepts declarations such as `var`, `const`, `type`, and `func`.

This works:

```go
package main

var nums = []int{1, 1, 1}
```

This does not work:

```go
package main

nums := []int{1, 1, 1}
```

`:=` is the short variable declaration syntax. It is only allowed inside function bodies:

```go
package main

import "fmt"

var nums = []int{1, 1, 1}

func main() {
	other := []int{2, 2, 2}

	fmt.Println(nums)
	fmt.Println(other)
}
```

Rule of thumb:

- Use `var name = value` at the package top level.
- Use `name := value` for local variables inside functions.

## One Directory Is One Package

In normal Go projects, one directory usually represents one package.

Example:

```text
calculator/
  add.go
  sub.go
```

Both files might use:

```go
package calculator
```

Then other packages can import it by path.

## What Is `package main`

`package main` is special. It builds an executable program.

A runnable program needs:

```go
package main

func main() {
}
```

The `main()` function is the program entry point.

For learning and algorithm practice, using `package main` is convenient because each exercise can be run directly with `go run`.

## Why Algorithm Problems Use Separate Directories Here

At first, this layout looked simple:

```text
algorithms/
  main.go
  lc001.go
  lc002.go
```

But all files under `algorithms/` would be one package.

That means `lc001.go` and `lc002.go` would share all top-level names. If both files define `solve`, `test`, `nums`, or `ListNode`, they collide.

So this repo uses one directory per problem:

```text
algorithms/
  lc001/
    main.go
  lc002/
    main.go
```

Now each problem is an independent `package main`.

This allows every problem to freely use simple names:

```go
func solve() {
}

func test() {
}

func main() {
	test()
}
```

No other problem can collide with it.

## Rule Of Thumb

Use the same directory when files are part of the same feature or package.

Use different directories when examples should be isolated from each other.

For this learning repo:

- Algorithm problems: one directory per problem.
- Small demos that belong together: same directory.
- Shared reusable code: create a normal package later, only when it is actually needed.


