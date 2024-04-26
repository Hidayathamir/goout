// Package main -.
package main

import (
	"gorm.io/gorm"
)

//nolint:gomnd
func main() {
	gormPlayground(func(pg *gorm.DB) {})
}
