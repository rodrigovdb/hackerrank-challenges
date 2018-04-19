package main

import "fmt"

func percent(value float32, percent int) float32 {
    return value * (float32(percent) / 100)
}

func totalCost(mealCost float32, tipPercent int, taxPercent int) float32 {
    return mealCost + percent(mealCost, tipPercent) + percent(mealCost, taxPercent)
}

func main() {
    var mealCost    float32
    var tipPercent  int
    var taxPercent  int

    fmt.Scan(&mealCost)
    fmt.Scan(&tipPercent)
    fmt.Scan(&taxPercent)

    fmt.Printf("The total meal cost is %.0f dollars.", totalCost(
        mealCost,
        tipPercent,
        taxPercent))
}
