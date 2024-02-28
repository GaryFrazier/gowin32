package test

import (
	"testing"
	"time"

	"github.com/GaryFrazier/gowin32/src/github.com/GaryFrazier/gowin32"
)

func TestCreateWindow(t *testing.T) {
	go gowin32.CreateWindow("Hello Window")

	for {
		time.Sleep(1000)
	}
}
