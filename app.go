package main

import (
	"context"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Get zone time
func (a *App) Zonetime(name string) string {
	loc, _ := time.LoadLocation(name)
	return time.Now().In(loc).Format("Mon 02 Jan 2006 15 04 05 pm")
}

func (a *App) Addnums(nums []int64) int64 {
	var sum int64
	for _, n := range nums {
		sum += int64(n)
	}
	// fmt.Println("nums = ", nums, "sum:", sum)
	return sum
}
