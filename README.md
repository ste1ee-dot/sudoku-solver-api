# SUDOKU SOLVER API

This project was made on purpose of learning [Go](https://go.dev/). It's probably really unoptimized, any criticizm is welcome. Also it is worth to mention that this backend API is stateless and made by only using standard go libraries.


## Usage

Use my project with go

```bash
go run .
```

to run it, from projects directory.
URL format should be localhost:8080/sudoku?x1y1=1&x2y1=5.......


## Packages used

 -	[fmt                    - implements formatted I/O with functions analogous to C's printf and scanf](https://pkg.go.dev/fmt)
 -	[log                     - defines a type, Logger, with methods for formatting output](https://pkg.go.dev/os)
 -	[strconv                - implements conversions to and from string representations of basic data types](https://pkg.go.dev/strconv)
 -	[net/http         - provides HTTP client and server implementations](https://pkg.go.dev/text/tabwriter)
