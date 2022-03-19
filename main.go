package main

import (
	"fmt"

	internal "github.com/Mellywins/gokrb5-kdc-wrapper/internal"
)

func main() {
	status, err := internal.EnsureServiceIsRunning()

	fmt.Printf("Is service running: %v, status is : %s", err, status)
}
