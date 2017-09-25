# gofra

`gofra` is simple fractal render library.

## Installation

To install `gofra`, simply run:
```
go get github.com/gitchander/gofra
```

## Getting started console util

Сonsole rendering program exists in the directory `gofra/cmd/gofra`

Util used package: [cli.go](https://github.com/urfave/cli)

You will build console program:
```
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
$ ./gofra render
```
created image file `fractal.png`

you can watch it in any viewer program.

### Scale fractal

```
$ ./gofra scale 2
```
changed scale factor in configuration file.
for rebuild fractal you will render again.

### Move position

For move center position used command move x y, where x and y relative coordinate values in range [0 ... 10].

to left border:
```
$ ./gofra move 0 5
```
or
```
$ ./gofra move 0 -
```

to right border:
```
$ ./gofra move 10 -
```

to top border:
```
$ ./gofra move - 0
```

to bottom border:
```
$ ./gofra move - 10
```

without move:
```
$ ./gofra move 5 5
```
or
```
$ ./gofra move - -
```

### Images
![default fractal](images/default-ms.png)