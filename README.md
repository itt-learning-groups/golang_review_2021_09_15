# Golang Review, 09-15-2021: Concurrency Part 3

## Exercise 1: Review the channel-operations table

### blank (for practicing)

* Channel is    `open` + `empty / not full`

  We attempt to    `read` from it

  We get:

* Channel is    `open` + `empty / not full`

  We attempt to    `write` to it

  We get:

* Channel is    `open` + `empty / not full`

  We attempt to    `close` it

  We get:

* Channel is    `open` + `not empty / full`

  We attempt to    `read` from it

  We get:

* Channel is    `open` + `not empty / full`

  We attempt to    `write` to it

  We get:

* Channel is    `open` + `not empty / full`

  We attempt to    `close` it

  We get:

* Channel is    `closed`

  We attempt to    `read` from it

  We get:

* Channel is    `closed`

  We attempt to    `write` to it

  We get:

* Channel is    `closed`

  We attempt to    `close` it

  We get:

* Channel is    `nil`

  We attempt to    `read` from it

  We get:

* Channel is    `nil`

  We attempt to    `write` to it

  We get:

* Channel is    `nil`

  We attempt to    `close` it

  We get:

### completed

* Channel is    `open` + `empty / not full`

  We attempt to    `read` from it

  We get:    block / hang <-- (could be risk of goroutine leak if this happens in a goroutine)

* Channel is    `open` + `empty-and-reader-is-active / not full`

  We attempt to    `write` to it

  We get:    value added to the channel

* Channel is    `open` + `empty / not full`

  We attempt to    `close` it

  We get:    closed channel

* Channel is    `open` + `not empty / full`

  We attempt to    `read` from it

  We get:    the (only / next) value from the channel

* Channel is    `open` + `not empty / full`

  We attempt to    `write` to it

  We get:    block / hang <-- (could be risk of goroutine leak if this happens in a goroutine)

* Channel is    `open` + `not empty / full`

  We attempt to    `close` it

  We get:    closed channel

* Channel is    `closed`

  We attempt to    `read` from it

  We get:    get a value

* Channel is    `closed`

  We attempt to    `write` to it

  We get:    **PANIC!!** <-- container terminates, restarts; request gets a 500

* Channel is    `closed`

  We attempt to    `close` it

  We get:    **PANIC!!** <-- container terminates, restarts; request gets a 500

* Channel is    `nil`

  We attempt to    `read` from it

  We get:    block / hang

* Channel is    `nil`

  We attempt to    `write` to it

  We get:    block / hang

* Channel is    `nil`

  We attempt to    `close` it

  We get:    **PANIC!!** <-- container terminates, restarts; request gets a 500

## Exercise 2: Separating generator/consumer concerns; Practicing clear channel ownership; Avoiding mutability issues; Identifyig Goroutine leaks

(*See 1st section of `main` function in `main.go` and the `ex2` package.*)

## "Exercise" 3: More goroutine leak examples + incomplete work / shutdown behavior

Walk through some informative examples highlighted in the Ardan Labs blog:

* <https://www.ardanlabs.com/blog/2018/11/goroutine-leaks-the-forgotten-sender.html>
* <https://www.ardanlabs.com/blog/2018/12/goroutine-leaks-the-abandoned-receivers.html>
* <https://www.ardanlabs.com/blog/2019/04/concurrency-trap-2-incomplete-work.html>

## Exercise 4: Error handling

(*See 2nd section of `main` function in `main.go` and the `ex4` package.*)
