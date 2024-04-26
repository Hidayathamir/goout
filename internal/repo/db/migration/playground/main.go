// Package main -.
package main

import (
	"gorm.io/gorm"
)

func main() {
	gormPlayground(func(_ *gorm.DB) {})
}
