# Pay Stopwatch

Have you ever been working a shift and wondered how your time worked translates into actual money earned? (I certainly have)

This program provides a stopwatch that calculates how much money you've earned (based on a given hourly pay rate) during the elapsed timeframe.

## Building

Built using the [Fyne](https://fyne.io) GUI framework for Go, this program is able to be compiled for many different targets with relative ease (Although the program was developed and tested primarily on Windows).

To build from source, ensure that Go is installed and clone the repository:

```bash
git clone https://github.com/odddollar/Pay-stopwatch.git
cd Pay-stopwatch
```

To run, use the normal method of:

```bash
go run .
```

However, to package to an executable binary, it is necessary to install Fyne's tooling:

```bash
go install fyne.io/fyne/v2/cmd/fyne@latest
```

Then run the below command to package. This will compile Pay Stopwatch to a binary for the current OS. To cross-compile for other platforms, please refer to [Fyne's packaging documentation](https://developer.fyne.io/started/packaging):

```
fyne package -icon Icon.png
```

### Note: Static content

__This is only necessary if ``Icon.png`` has been modified.__

In order to use ``Icon.png`` as the icon for Pay Stopwatch, it is necessary to bundle it to a ``.go`` file. Fyne's tooling provides an easy way to do this (Please remove the old ``icon.go`` before running):

```bash
fyne bundle Icon.png >> icon.go
```

