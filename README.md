# Pay Stopwatch

Have you ever been working a shift and wondered how your time worked translates into actual money earned? (I certainly have)

This program provides a stopwatch that calculates how much money you've earned (based on a given hourly pay rate) during the elapsed timeframe.

## Building

Built using the [Fyne](https://fyne.io/) GUI framework for Go, this program can be compiled to a single binary with the following commands:

```
git clone https://github.com/odddollar/Pay-stopwatch.git
cd Pay-stopwatch
go install fyne.io/fyne/v2/cmd/fyne@latest // installs the necessary Fyne tooling
fyne package --release
```

## Screenshots

![Image 1](/screenshots/image1.png)

![Image 2](/screenshots/image2.png)
