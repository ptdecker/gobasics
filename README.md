# Worked Excercises from _The Go Programming Language_, by Alan A. A. Donovan and Brian W. Kernighan

This repository contains solutions to the examples presented in the book
[_The Go Programming Language_](http://www.gopl.io) by Alan A. A. Donovan [Brian W. Kerningham](https://en.wikipedia.org/wiki/Brian_Kernighan).

## Worked Exercises

Exercises postfixed with a "(+)" indicate ones for which a benchmark test has been provided (see 'Benchmarks' section below)

### Chapter 1 - "A Tutorial Introduction"

- 1-01 -
  - echo4: Variation of 'echo3' which also prints the name of the envoking command
- 1-02 -
  - echo5: Variation of 'echo2' which also prints the index of each argument one per line
- 1-03 -
  - echo6 (+): Variation of 'echo2' that also displays timing. Compare against 'echo7' for exercise 1-03 solution.
  - echo7 (+): Variation of 'echo3' that also displays timing. Execution is slightly faster than 'echo6'
- 1-04 -
  - dup4: Prints the names of all files in which a duplicate line appears
- 1-05 -
  - lissajous2: Variation o 'lissajous1' with a color variation of a traditional oscilliscope image (generates a GIF file, e.g. `bin/lissajous1 > out.gif && open out.gif`)
- 1-06 -
  - lissajous3: Variation of 'lissajous1' with a dynamic color variation based upon current cycle (generates a GIF file)
- 1-07 -
  - fetch2: Variation of 'fetch1' that leverages io.Copy to avoid the need for a buffer
- 1-08 -
  - fetch3: Variation of 'fetch2' that adds an 'http://' prefix if it is missing from the URL
- 1-09 -
  - fetch4: Variation of 'fetch3' that also prints the HTTP status code
- 1-10 -
  - fetchall2: Variation of 'fetch1' that fetches the URLs twice without exiting to shell. To satisfy the excercise requirement of writting the output to a file, run with `bin/fetchall2 http://facebook.com > results.txt` to pipe output to results.txt.
- 1-11 -
  - Use 'fetchall1' to run a sample of sites from [Alexa](http://www.alexa.com/topsites). A possible run could be `bin/fetchall1 http://google.com http://youtube.com http://facebook.com http://baidu.com http://yahoo.com http://wikipedia.org http://amazon.com http://twitter.com http://qq.com > results.txt`. I was not able to find a site that would not respond. So, you can use extras/norespserver as a testing tool.
- 1-12 -
  - lissajous5: Variation of 'lissajous4' adding support for passing a 'cycles' parameter to set number of cycles in addition to the basic web server 'lissajous4' added to 'lissajous3'. Generates a GIF file

### Chapter 2 - "Program Structure"

- 2-01 -
  - tempconv1 (p): Extends 'tempconv' to include types, constants, and functions to support the Kelvin scale
  - cfk [tempconv1]: Variation of 'cf' extended to support Kelvin conversions using 'tempconv1'
- 2-02 -
  - genconv (p): Generic conversions package
  - gounits (genconv): Converts a broad range of units from one to the other
- 2-03 -
  - popcount1 (p, +): Variation on 'popcount' which uses a loop as opposed to a single expression
  - popcount1demo: Demo program for popcount1
- 2-04 -
  - popcount2 (p, +): Variation on 'popcount' which uses bit shifting as opposed to a lookup table
  - popcount2demo: Demo program for popcount2
- 2-05 -
  - popcount3 (p, +): Variation on 'popcount' which uses bit clearing as opposed to a lookup table
  - popcount3demo: Demo program for popcount3

### Chapter 3 - "Basic Data Types"

- 3-01 -
  - surface1 (+): Variation on 'surface' that adds guard code for invalid polygons. Generates an SVG (`bin/surface1 > out.png && -a "Google Chrome" out.svg`)
  - surface2 (+): A significantly faster variation of 'surface1' that outputs integers instead of floats. Generates an SVG
- 3-02 - (all answers generate an SVG)
  - surface3: Variation of 'surface1' that renders an egg box-like surface
  - surface4: Variation of 'surface1' that renders a mogul-like surface
  - surface5: Variation of 'surface1' that renders a "monkey saddle" surface
- 3-03 - (all answers generate an SVG)
  - surface6: Variation of 'surface1' that colors the polygons from red at peaks to blue
- 3-04 - (all answers generate an SVG)
  - surface7: Variation of 'surface6' plus the addition of a built-in basic web server
- 3-05 - (all answers generate an PNG `bin/mandelbrot1 > out.png && open -a "Google Chrome" out.png`)
  - mandelbrot1: Adds full-color rendering to 'mandelbrot'
- 3-06 - (all answers generate an PNG)
  - mandelbrot2: Adds supersampling anti-aliasing to 'mandelbrot1'
- 3-07 - (all answers generate an PNG)
  - newton1: Uses Newton's method to render z^4 - 1 = 0 Julia set fractal
- 3-08 - ** UNSOLVED ** (all answers generate an PNG)
  - mandelbrot3: Re-implements 'mandelbrot1' using complex64 with benchmarks
  - mandelbrot4: Re-implements 'mandelbrot1' using complex128 with benchmarks
  - mandelbrot5: Re-implements 'mandelbrot1' using big.Float with benchmarks
  - mandelbrot6: Re-implements 'mandelbrot1' using big.Rat with benchmarks
- 3-09 - (all answers generate an SVG)
  - mandelbrot7: Re-implements 'mandelbrot1' as a web service supporting x,y coordinates and zoom parameters. For example, try 'http://localhost:8000/?zoom=1000&x=-0.7325&y=-0.241'.
- 3-10 -
  - comma1: A non-recursive version of 'comma'
- 3-11 -
  - comma2: Adds support of floating-point numbers to 'comma1'
- 3-12 -
  - isanagram1: Tests two strings to see if they are anagrams of each other
- 3-13 -
  - const: Defines decimal orders of magnitude constants

### Chapter 4 - "Composite Types"

- 4-01 -
  - sha256-2: determines the number of bits that differ between two SHA256 hashes
- 4-02 -
  - sha: prints SHA256, SHA384, or SHA512 hashes. SHA256 is default. Use '-384' or '-512' to select alternate approaches.
- 4-03 -
  - rev1: reverses the elements of an integer array passed to the reverse function as a pointer
- 4-04 -
  - rotate1: rotates an array using a single pass
- 4-05 -
  - dedup: an in-place function that eliminates adjacent duplicates in a []string slice

### Chapter 8 - "Goroutines and Channels"

- 8-01 -
  - clock3: A concurent time server on a given TCP port
  - clockwall: A clock client that subscribes to a list of time servers

### Data

The following test data files are provided to support manually testing the exercises. Specifically, the ones that draw from standard input can be manually tested with these.

- 448bits.txt - Contains exactly 448 bits of test text data with no trailing new line. Complies with one of the test vectors for testing the 'sha' program (c.f. https://www.di-mgt.com.au/sha_testvectors.html)
- 896bits.txt - Contains exactly 896 bits of test text data with no trailing new line. Complies with one of the test vectors for testing the 'sha' program (c.f. https://www.di-mgt.com.au/sha_testvectors.html)
- abc.txt - Contains text string of "abc" with no trailing new line (24 bits). Complies with one of the test vectors for testing the 'sha' program (c.f. https://www.di-mgt.com.au/sha_testvectors.html)
- empty.txt - Contains nothing--A totally empty file with no new line nor end of file characters. Good for testing empty file handling. Complies with one of the test vectors for testing the 'sha' program (c.f. https://www.di-mgt.com.au/sha_testvectors.html)
- hasdups.txt - Contains duplicate lines of data.
- nodups.txt - Contains no duplicate lines of data.
- numbers.txt - Test numbers of various formats (positive and negative, integers and floating point)

### Extras

The following programs are provided as extras.

- norespserver.go: A basic http server that never responds to requests sent to it. Use to test exercise 1-11

## Compiling

A bash shell script is provided to compile and install all the exercise packages. Successfully installed packages are placed in the 'bin' sub-directory.

    $ ./build.sh

## Solution Notes

Each worked excercise attempts to only use capabilities that have been introduced so far in the book (i.e. no forward references to knowledge). Where this is deliberately not the case, a comment in the code is provided.

## Book Examples

This repository also includes working examples from the book. These are also available from the authors, but the versions that I hand entered (hard way as opposed to pulling them down from the author) are included for completness. "(p)" is used to indicate packages which contain no 'main' and thus will appear directy in `bin`. For convienence, when such packages are used, the package which ccontains 'main' lists the other packages it uses.

### Chapter 1

- Section 1.1 - "Hello World"
  - helloworld
- Section 1.2 - "Command-Line Arguments"
  - echo1
  - echo2
  - echo3
- Section 1.3 - "Finding Duplicate lines"
  - dup1
  - dup2
  - dup3
- Section 1.4 - "Animaged GIFs"
  - lissajous1
- Section 1.5 - "Fetching a URL"
  - fetch1
- Section 1.6 - "Fetching URLs Concurrently"
  - fetch2
- Section 1.7 - "A Web Server"
  - server1
  - server2
  - server3
  - lissajous4 - 'lissajous3' with built-in web server

### Chapter 2

- Section 2.2 - "Declarations"
  - boiling
  - ftoc
- Section 2.3.2 - "Pointers"
  - echo8 (named 'echo4' in the text)
- Section 2.6 - "Packages and Files"
  - tempconv (p)
  - cf [tempconv]
- Section 2.6.2 - "Package Initialization
  - popcount (p)
  - popcountdemo

### Chapter 3

- Section 3.2 - "Floating-Point Numbers"
  - surface
- Section 3.3 - "Complex Numbers"
  - mandelbrot

### Chapter 4

- Section 4.1 - "Arrays"
  - rev

### Chapter 8

- Section 8.2 - "Goroutines"
  - spinner
- Section 8.3 - "Exmaple: Concurent Clock Server"
  - clock1 - single threaded basic clock server
  - clock2 - multi-threaded clock server
  - netcat1 - read-only TCP net client

## Benchmarking

Some of the exercises also are provided with benchmark tests. These exercises are indicated in the list above with a '(+)' after the program name. Section 11.4 discusses how to execute these tests. But, in general, they can be executed using the following commands. The last command returns you to the main GOPATH directory.

    $ go test -bench=. <chapter>/<exercise program or package name>

For example, to run the benchmark associated with 'echo6' from exercise 1-03 you would:

    $ go test -bench=.

In addition, a script named 'benchmare.sh' has been provided to execute all benchmarks

## License

These solutions are copyright under the terms of the GNU GENERAL PUBLIC LICENSE v3 (see LICENSE) with the following exceptions:

- By definition, the code presented here draws from, and directly answers the excercises presented within, Donovan's and Kernighan's and Ritchie's "The Go Programing Language."

## Handy Commands

Hex Dumps

    hexdump -C -n128 data/448bits.txt

Removing Trailing New Line Characters

    tr -d '\n' < data/448bits.txt > data/448bits2.txt
