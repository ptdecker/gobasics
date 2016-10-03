Worked Excercises from _The Go Programming Language_, by Alan A. A. Donovan and Brian W. Kernighan 
=======

This repository contains solutions to the examples presented in the book
[_The Go Programming Language_](http://www.gopl.io) by Alan A. A. Donovan [Brian W. Kerningham](https://en.wikipedia.org/wiki/Brian_Kernighan).

## Worked Exercises

### Chapter 1 - "A Tutorial Introduction"

* 1-01 - echo4:  Variation of 'echo3' which also prints the name of the envoking command
* 1-02 - echo5:  Variation of 'echo2' which also prints the index of each argument one per line
* 1-03 -
  * echo6 (+):  Variation of 'echo2' that also displays timing. Compare against 'echo7' for exercise 1-03 solution.
  * echo7 (+):  Variation of 'echo3' that also displays timing. Execution is slightly faster than 'echo6'
* 1-04 - dup4:  Prints the names of all files in which a duplicate line appears
* 1-05 - lissajous2:  Variation o 'lissajous1' with a color variation of a traditional oscilliscope image
* 1-06 - lissajous3:  Variation of 'lissajous1' with a dynamic color variation based upon current cycle
* 1-07 - fetch2:  Variation of 'fetch1' that leverages io.Copy to avoid the need for a buffer
* 1-08 - fetch3:  Variation of 'fetch2' that adds an 'http://' prefix if it is missing from the URL
* 1-09 - fetch4:  Variation of 'fetch3' that also prints the HTTP status code
* 1-10 - fetchall2:  Variation of 'fetch1' that fetches the URLs twice without exiting to shell. To satisfy the excercise requirement of writting the output to a file, run with `bin/fetchall2 http://facebook.com > results.txt` to pipe output to results.txt.
* 1-11 - Use 'fetchall1' to run a sample of sites from [Alexa](http://www.alexa.com/topsites). A possible run could be `bin/fetchall1 http://google.com http://youtube.com http://facebook.com http://baidu.com http://yahoo.com http://wikipedia.org http://amazon.com http://twitter.com http://qq.com > results.txt`. I was not able to find a site that would not respond.  So, you can use extras/norespserver as a testing tool.
* 1-12 - lissajous5:  Variation of 'lissajous4' adding support for passing a 'cycles' parameter to set number of cycles in addition to the basic web server 'lissajous4' added to 'lissajous3'

### Chapter 2 - "Program Structure"

* 2-01 -
  * tempconv1 [p]:  Extends 'tempconv' to include types, constants, and functions to support the Kelvin scale
  * cfk [tempconv1]: Variation of 'cf' extended to support Kelvin conversions using 'tempconv1'  
* 2-02 -
  * genconv [p]:  Generic conversions package
  * gounits [genconv]:  Converts a broad range of units from one to the other

### Extras

The following programs are provided as extras.

* norespserver.go:  A basic http server that never responds to requests sent to it.  Use to test exercise 1-11

## Compiling

A bash shell script is provided to compile and install all the exercise packages.  Successfully installed packages are placed in the 'bin' sub-directory.

    $ ./build.sh

## Solution Notes

Each worked excercise attempts to only use capabilities that have been introduced so far in the book (i.e. no forward references to knowledge).  Where this is deliberately not the case, a comment in the code is provided.

## Book Examples

This repository also includes working examples from the book.  These are also available from the authors, but the versions that I hand entered (hard way as opposed to pulling them down from the author) are included for completness. "[p]" is used to indicate packages which contain no 'main' and thus will appear directy in `bin`.  For convienence, when such packages are used, the package which ccontains 'main' lists the other packages it uses.

### Chapter 1

* Section 1.1 - "Hello World"
  * helloworld
* Section 1.2 - "Command-Line Arguments"
  * echo1
  * echo2
  * echo3
* Section 1.3 - "Finding Duplicate lines"
  * dup1
  * dup2
  * dup3
* Section 1.4 - "Animaged GIFs"
  * lissajous1
* Section 1.5 - "Fetching a URL"
  * fetch1
* Section 1.6 - "Fetching URLs Concurrently"
  * fetch2
* Section 1.7 - "A Web Server"
  * server1
  * server2
  * server3
  * lissajous4 - 'lissajous3' with built-in web server 

## Chapter 2

* Section 2.2 - "Declarations"
  * boiling
  * ftoc
* Section 2.3.2 - "Pointers"
  * echo8 [named 'echo4' in the text]
* Section 2.6 - "Packages and Files"
  * tempconv [p]
  * cf [tempconv]

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

* hasdups.txt: a text file containing duplicate lines
* nodups.txt: a text file that does not contain duplicate lines
* numbers.txt: a text file containing one number (positive and negative integers and floating point) per line

##License

These solutions are copyright under the terms of the GNU GENERAL PUBLIC LICENSE v3 (see LICENSE) with the following exceptions:

* By definition, the code presented here draws from, and directly answers the excercises presented within, Donovan's and Kernighan's and Ritchie's "The Go Programing Language."
