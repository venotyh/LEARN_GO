# Project Instructions

This is a personal Go learning repository. Keep changes small, direct, and easy to understand.

## Repository Layout

- `basics/`: Go syntax and language basics.
- `algorithms/`: Algorithm practice.
- `patterns/`: Design pattern experiments in Go style.
- `stdlib/`: Standard library experiments.
- `projects/`: Small practice projects.
- `notes/`: Learning notes.

## Algorithm Exercises

Use one directory per problem under `algorithms/`.

Examples:

```text
algorithms/
  lc001/
    main.go
  lc002/
    main.go
  cf555/
    main.go
```

Each problem directory should be an isolated `package main` so top-level variables, functions, and types do not collide with other problems.

For beginner-friendly exercises, a single `main.go` is enough. It can include:

- The solution function.
- A small `test()` function with sample cases.
- A `main()` function that calls `test()`.

Do not split into `solution.go`, `solution_test.go`, and `README.md` unless explicitly requested.

## Verification

After changing an algorithm exercise, run the specific problem:

```powershell
go run ./algorithms/lc001
```

Before finishing broader changes, run:

```powershell
go test ./...
```

## Style

- Prefer simple Go over abstractions.
- Do not add frameworks or helper libraries for basic exercises.
- Match the current learning-oriented style.
- Keep explanations and comments short.
