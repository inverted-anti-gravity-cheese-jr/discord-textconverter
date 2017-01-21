package main

import ("bufio"
    "os"
    "fmt"
    "github.com/atotto/clipboard")

func toLowerCase(s []byte) {
    for i := 0; i < len(s); i++ {
        if(s[i] >= 65 && s[i] <= 90) {
            s[i] = s[i] + (97 - 65)
        }
    }
}

func printDiscordLetters(s []byte) {
    fmt.Println(discordLettersToString(s))
}

func discordLettersToString(s []byte) string {
    var str = string("")
    for i := 0; i < len(s); i++ {
        if((s[i] >= 97 && s[i] <= 122)) {
            str += ":regional_indicator_" + string(s[i]) + ": "
        }
        if(s[i] == 32) {
            str += "   "
        }
        switch s[i] {
        case 48:
            str += ":zero: "
        case 49:
            str += ":one: "
        case 50:
            str += ":two: "
        case 51:
            str += ":three: "
        case 52:
            str += ":four: "
        case 53:
            str += ":five: "
        case 54:
            str += ":six: "
        case 55:
            str += ":seven: "
        case 56:
            str += ":eight: "
        case 57:
            str += ":nine: "
        }
    }
    return str
}

func main() {
    for {
        fmt.Println("Wpisz jakis tekst w ASCII")
        read := bufio.NewReader(os.Stdin)
        buf, _, _ := read.ReadLine()
        toLowerCase(buf)
        fmt.Println("Wklej to do Discord\n\n")
        printDiscordLetters(buf)
        clipboard.WriteAll(discordLettersToString(buf))
        fmt.Println("Skopiowano do schowka\n")
    }
}