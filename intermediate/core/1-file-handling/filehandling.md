# File handling

- File handling essentially means working with it, such as getting metadata information, creating new files or simply reading and writing data to and from a file.
- In Golang, the API for file handling is well-knitted into the standard architecture and we can do it using standard library.
- The `OS` package provides and API interface for file handling which is uniform across all operation systems. It provides functionality such as creating, deletion, opening a file, modifying its permissions and so on.
- The `io` package provides interfaces for basic i/o primitives and wraps them into easy to use public interfaces.
- The `filepath` package that would provide functions to parse and construct file paths.
- We also use the `fmt` package to format i/o with functions to read and write to the standard input and output.
