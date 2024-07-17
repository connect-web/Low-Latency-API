package main

import (
	"fmt"
	"github.com/connect-web/Low-Latency-API/internal/db/user"
)

func main() {
	r, err := user.FetchOrCreateProfile("letmein")
	fmt.Println(err)
	fmt.Println(r)
}
