package goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello world")
}
func DisplayNumber(number int) {
	fmt.Println("Display", number)
}
func TestCreateGoroutine(t *testing.T) {
	go RunHelloWorld() // berjalan secara async
	fmt.Println("Ups")
	// pause selama 1 detik
	time.Sleep(1 * time.Second)
}
func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}
