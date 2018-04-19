package main
import "fmt"

func isWeird(x int) string {
    var weird   string = "Weird"
    var not     string = "Not Weird"

    if x %2 != 0 {
        return weird
    } else if (x < 5 || x > 20) {
        return not
    } else {
        return weird
    }
}

func main() {
    var x int

    fmt.Scan(&x)

    fmt.Printf("%s", isWeird(x))
}
