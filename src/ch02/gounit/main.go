// Gounit is a Go-based unit conversion tool that converts from one unit to another
// based upon command-line parameters or standard input
package main

import (
    "bufio"
    "fmt"
    "strconv"
    "os"
    "ch02/genconv"
)

func main() {
    if len(os.Args) == 3 {
        input := bufio.NewScanner(os.Stdin)
        for input.Scan() {
            f, err := strconv.ParseFloat(input.Text(), 64)
            if err != nil {
                fmt.Fprint(os.Stderr, "%s: error converting '%s' to floating point\n", os.Args[0], input.Text())
            } else {
                conv(f, os.Args[1], os.Args[2])
            }
        } 
    } else if len(os.Args) == 4 {
        f, err := strconv.ParseFloat(os.Args[1], 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "%s: error converting '%s' to floating point\n", os.Args[0], os.Args[1])
            os.Exit(1)
        }
        conv(f, os.Args[2], os.Args[3])
    } else {
        help()
    }
}

func conv(v float64, from string, to string) {

    if from == "F" && to == "C" {
        f := genconv.Fahrenheit(v)
        fmt.Printf("%s = %s\n", f, genconv.FToC(f))
        return
    }

    if from == "F" && to == "K" {
        f := genconv.Fahrenheit(v)
        fmt.Printf("%s = %s\n", f, genconv.FToK(f))
        return
    }

    if from == "C" && to == "F" {
        c := genconv.Celsius(v)
        fmt.Printf("%s = %s\n", c, genconv.CToF(c))
        return
    }

    if from == "C" && to == "K" {
        c := genconv.Celsius(v)
        fmt.Printf("%s = %s\n", c, genconv.CToK(c))
        return
    }

    if from == "K" && to == "C" {
        k := genconv.Kelvin(v)
        fmt.Printf("%s = %s\n", k, genconv.KToC(k))
        return
    }

    if from == "K" && to == "F" {
        k := genconv.Kelvin(v)
        fmt.Printf("%s = %s\n", k, genconv.KToF(k))
        return
    }

    if from == "ft" && to == "m" {
        ft := genconv.Foot(v)
        fmt.Printf("%s = %s\n", ft, genconv.FtToM(ft))
        return
    }

    if from == "m" && to == "ft" {
        m := genconv.Meter(v)
        fmt.Printf("%s = %s\n", m, genconv.MToFt(m))
        return
    }

    if from == "lb" && to == "kg" {
        lb := genconv.Pound(v)
        fmt.Printf("%s = %s\n", lb, genconv.LbToKg(lb))
        return
    }

    if from == "kg" && to == "lb" {
        kg := genconv.Kilogram(v)
        fmt.Printf("%s = %s\n", kg, genconv.KgToLb(kg))
        return
    }

    fmt.Fprintf(os.Stderr,"%s: unknown conversion from '%s' to '%s'\n", os.Args[0], from, to);
    os.Exit(1)

}
       
func help() {
    fmt.Println("usage:\tgounit [value] from-units to-units\n")
    fmt.Println("\twhere:\n")
    fmt.Println("\t\t[value] = value in from-units to convert to to-units")
    fmt.Println("\t\tfrom-units = units to convert from")
    fmt.Println("\t\tto-units = units to convert to\n")
    fmt.Println("\tunits supported:\n")
    fmt.Println("\t\t'F'  - Fahrenheit")
    fmt.Println("\t\t'C'  - Celsius")
    fmt.Println("\t\t'K'  - Kelvin")
    fmt.Println("\t\t'ft' - Feet")
    fmt.Println("\t\t'm'  - Meters")
    fmt.Println("\t\t'lb' - Pounds")
    fmt.Println("\t\t'kg' - Kilograms\n")
    fmt.Println("Numbers are read from stdin, one per line, if [value] is omitted\n")
}
