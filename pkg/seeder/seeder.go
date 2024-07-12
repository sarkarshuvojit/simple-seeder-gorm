// Package seeder provides a type to declare individual seeder methods and a function RunSeeder using which set of seeder methods could be run within a single transaction
package seeder

import (
	"log"

	"gorm.io/gorm"
)

// Seed is used to declare seeders, each contain a name and a method, additionally contains a Skip flag to skip certain seeders without removing them.
type Seed struct {
	Name string `json:"name"`
	Skip bool   `json:"skip"`
	Run  func(*gorm.DB) error
}

// RunSeeder takes an instance of gorm and a set of seeds. It then creates a transaction & runs all non-skipped seeds against the transaction created from gorm instance
func RunSeeder(db *gorm.DB, seeds []Seed) error {
	tx := db.Begin()
	for _, seed := range seeds {
		if seed.Skip {
			log.Printf("Skipping seed: %s\n", seed.Name)
		} else {
			if err := seed.Run(tx); err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
