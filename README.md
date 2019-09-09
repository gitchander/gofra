# gofra

`gofra` is simple fractal render library.

![default fractal](images/default-ms.png)

## Installation

To install `gofra`, simply run:
```
go get github.com/gitchander/gofra
```

## Getting started console util

Сonsole rendering program exists in the directory `gofra/cmd/gofra`

It used package: [cli.go](https://github.com/urfave/cli)

You will build console program:
```
$ cd gofra/cmd/gofra
$ go build
```

### Make default
For make default file config:
```
$ ./gofra default
```
in this case the file will be created `fractal.json`

### Render fractal
For render first fractal, run:
```
$ ./gofra draw
```
created image file `fractal.png`

```
$ ./gofra --image my_fractal.png draw
```
created image file `my_fractal.png`


you can watch it in any viewer program.

### Scale fractal

```
$ ./gofra scale 2
```
changed scale factor in configuration file.
for make fractal image you will render again.

Or you can run:
```
$ ./gofra --render scale 2
```
Also it makes render.

### Move position

For move center position used command move x y, where x and y relative coordinates.
The coordinates encodes with next characters:
'm' - minus
'z' - 0.0   // zero    = 0
'w' - 1.0   // whole   = 1
'h' - 0.5   // half    = (1 / 2)
'q' - 0.25  // quarter = (1 / 4)
'e' - 0.125 // eighth  = (1 / 8)

Examples:

```
./gofra move z z
```
x = 0, y = 0

```
./gofra move h mw
```
x = 0.5, y = -1

```
./gofra move mq e
```
x = -0.25, y = 0.125

```
./gofra move hq z
```
x = 0.75, y = 0
