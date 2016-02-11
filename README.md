# gofra

`gofra` is simple fractal render library.

## Installation

To install `gofra`, simply run:
```
go get github.com/gitchander/gofra
```

## Getting Started

Сonsole rendering program exists in the directory `gofra/fractus`

You will build console program:
```
$ go build
```

### Make default
For make default file config:
```
$ ./fractus default
```
in this case the file will be created `fractal.json`

### Render fractal
For render first fractal, run:
```
$ ./fractus render
```
created image file `fractal.png`

you can watch it in any viewer program.

### Scale fractal

```
$ ./fractus scale 2
```
changed scale factor in configuration file.
for rebuild fractal you will render again.

### Move position

```
$ ./fractus move 0.5 0
```

```
(-1, 1) ------ (0, -1)
    |             |
    |             |
    |             |
(-1, 0) -------- (0, 0)
    |             |
    |             |
    |             |
(-1, -1)---------(0, -1)


  | -1 | 0 | 1 |
----------------
1 |
----------------
0 |
----------------
-1 |

```

