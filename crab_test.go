package crab_test

import (
	"crab"
	"fmt"
	"testing"
)

func TestCrab(t *testing.T) {
	蟹 := crab.New蟹("Crab55e")
	鍋 := 蟹.Make鍋()
	fmt.Println(鍋)
}
