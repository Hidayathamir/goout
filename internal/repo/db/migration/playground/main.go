// Package main is the entry point for debug gorm.
package main

import (
	"gorm.io/gorm"
)

func main() {
	gormPlayground(func(_ *gorm.DB) {})
}
