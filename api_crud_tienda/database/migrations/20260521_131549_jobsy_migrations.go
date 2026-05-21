package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type JobsyMigrations_20260521_131549 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &JobsyMigrations_20260521_131549{}
	m.Created = "20260521_131549"

	migration.Register("JobsyMigrations_20260521_131549", m)
}

// Run the migrations
func (m *JobsyMigrations_20260521_131549) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *JobsyMigrations_20260521_131549) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
