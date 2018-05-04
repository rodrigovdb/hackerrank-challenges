package main
import (
    "fmt"
    "strconv"
)

func count_consecutives(str string) int {
    count       := 1
    response    := 1
    prev        := "0"

    for _, r := range str {
        c := string(r)

        if c == "1" {
            if prev == "1" {
                count += 1
            } else {
                count = 1
            }
        }

        if count > response {
            response = count
        }

        prev = c
    }

    return response
}

func to_binary(number int) string {
    var remainder int
    var response  string

    for number > 0 {
        remainder = number % 2
        number    = number / 2

        response = strconv.Itoa(remainder) + response
    }

    return response
}

func main() {
    var input int
    fmt.Scan(&input)

    bin     := to_binary(input)
    count   := count_consecutives(bin)

    fmt.Println(count)
}
