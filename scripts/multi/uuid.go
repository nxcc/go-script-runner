package main

import (
	"fmt"

	"github.com/google/uuid"
)

func UUID() {
	fmt.Println("uuid:", uuid.New().String())
}
