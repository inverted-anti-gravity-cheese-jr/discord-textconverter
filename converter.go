package main

import ("bufio"
    "os"
    "fmt")

func toLowerCase(s []byte) {
    for i := 0; i < len(s); i++ {
        if(s[i] >= 65 && s[i] <= 90) {
            s[i] = s[i] + (97 - 65)
        }
    }
}

func printDiscordLetters(s []byte) {
    for i := 0; i < len(s); i++ {
        if((s[i] >= 97 && s[i] <= 122)) {
            fmt.Printf(":regional_indicator_%c: ", s[i])
        }
        if(s[i] == 32) {
            fmt.Printf("   ");
        }
        switch s[i] {
        case 48:
            fmt.Printf(":zero: ")
        case 49:
            fmt.Printf(":one: ")
        case 50:
            fmt.Printf(":two: ")
        case 51:
            fmt.Printf(":three: ")
        case 52:
            fmt.Printf(":four: ")
        case 53:
            fmt.Printf(":five: ")
        case 54:
            fmt.Printf(":six: ")
        case 55:
            fmt.Printf(":seven: ")
        case 56:
            fmt.Printf(":eight: ")
        case 57:
            fmt.Printf(":nine: ")
        }
    }
}

func main() {
    fmt.Println("Wpisz jakis tekst w ASCII")
    read := bufio.NewReader(os.Stdin)
    buf, _, _ := read.ReadLine()
    toLowerCase(buf)
    fmt.Println("Wklej to do Discord")
    printDiscordLetters(buf)
    fmt.Println()
    fmt.Println("Wpisz cokolwiek aby zakonczyc")
    read.ReadLine()
}