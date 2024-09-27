package main


import (
    "time"
    "github.com/go-vgo/robotgo"
)

// Function to press the right arrow key repeatedly
func pressRightArrow(repeats int, delay time.Duration) {
    for i := 0; i < repeats; i++ {
        robotgo.KeyTap("right") // Simulate pressing the right arrow key
        time.Sleep(delay)       // Wait for the specified delay
    }
}

func main(){
	// Press the right arrow key 10 times with a 500ms delay between each press
	pressRightArrow(10, 500*time.Millisecond)
}