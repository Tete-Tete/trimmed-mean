package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Tete-Tete/trimmedmean"
)

func main() {
	intData := make([]float64, 100)
	floatData := make([]float64, 100)

	for i := 0; i < 100; i++ {
		intData[i] = float64(i + 1)
		floatData[i] = float64(i) + 0.5
	}

	trim := 0.05
	meanInt, err := trimmedmean.TrimmedMean(intData, trim, trim)
	if err != nil {
		log.Fatalf("Error computing trimmed mean for integers: %v", err)
	}
	fmt.Printf("Trimmed mean of integers: %.2f\n", meanInt)

	meanFloat, err := trimmedmean.TrimmedMean(floatData, trim, trim)
	if err != nil {
		log.Fatalf("Error computing trimmed mean for floats: %v", err)
	}
	fmt.Printf("Trimmed mean of floats: %.2f\n", meanFloat)

	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadString('\n')
}
