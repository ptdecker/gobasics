Worked Excercises from _The Go Programming Language_, by Alan A. A. Donovan and Brian W. Kernighan 
=======

This repository contains solutions to the examples presented in the book
[_The Go Programming Language_](http://www.gopl.io) by Alan A. A. Donovan [Brian W. Kerningham](https://en.wikipedia.org/wiki/Brian_Kernighan).

## Worked Exercises

### Chapter 1 - "A Tutorial Introduction"

* 1-01 - echo4:  Variation on 'echo3' which also prints the name of the envoking command
* 1-02 - echo5:  Variation on 'echo2' which also prints the index of each argument one per line
* 1-03 -
  * echo6 (+):  Variation of 'echo2' that also displays timing. Compare against 'echo7' for exercise 1-03 solution.
  * echo7 (+):  Variation of 'echo3' that also displays timing. Execution is slightly faster than 'echo6'

## Book Examples

This repository also includes working examples from the book.  These are also available from the authors, but the versions that I hand entered (hard way as opposed to pulling them down from the author) are included for completness.

### Chapter 1

* Section 1.1 - "Hello World"
  * helloworld
* Section 1.2 - "Command-Line Arguments"
  * echo1
  * echo2
  * echo3

##Benchmarking

Some of the exercises also are provided with benchmark tests.  These exercises are indicated in the list above  with a '(+)' after the program name.  Section 11.4 discusses how to execute these tests.  But, in general, they can be executed using the following commands.  The last command returns you to the main GOPATH directory.

    $ cd $GOPATH/src/<chapter>/<exercise program name>
    $ go test -bench .
    $ cd $GOPATH

For example, to run the benchmark assoicated with 'echo6' from exercise 1-03 you would:

    $ cd $GOPATH/src/ch01/echo6
    $ go test -bench .
    $ cd $GOPATH

##Test Files

The following test data files are provided to support manually testing the exercises. Specifically, the ones that draw from standard input can be manually tested with these.

None yet

##License

These solutions are copyright under the terms of the GNU GENERAL PUBLIC LICENSE v3 (see LICENSE) with the following exceptions:

* By definition, the code presented here draws from, and directly answers the excercises presented within, Donovan's and Kernighan's and Ritchie's "The Go Programing Language."
