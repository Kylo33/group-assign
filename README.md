CLI tool for assigning learning group problems for CSCI425
and similar courses.

# Features

- Make sure each problem is solved by exactly two people (this is
  configurable)!
- Make sure problems are divided evenly; nobody should solve more than one
  more problem than any other student in the group.

# Installation

The `group-assign` binary is available with `go install`:

```sh
go install github.com/Kylo33/group-assign@latest
```

> [!NOTE]
> Make sure your Go binaries are on your `$PATH`
> ([Download and Install: Go](https://go.dev/doc/install)).

# Usage

```
Usage of group-assign:
  -coverage int
    	number of people per problem (default 2)
  -people string
    	comma separated list of group members
  -problems int
    	number of problems to assign
```

## Example

```
$ group-assign -problems=5 -people=Alice,Bob,Charlie,David

Alice: [3 4 5]
Charlie: [1 4]
David: [1 2 3]
Bob: [2 5]
```
