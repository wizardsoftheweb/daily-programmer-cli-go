// Copyright 2021 CJ Harries
// Licensed under http://www.apache.org/licenses/LICENSE-2.0

package cmd

import (
	"database/sql"
	"fmt"
	"os"
	"path"
)

const defaultDatabaseRoot string = "."
const defaultDatabaseBasename string = "daily_prog.db"
const defaultDatabaseType string = "sqlite3"

type ChallengeDatabase struct {
	database *sql.DB
	source   string
	driver   string
}

func (c *ChallengeDatabase) constructSource(args ...string) {
	c.source = path.Join(args...)
}

func (c *ChallengeDatabase) createDatabase() {
	_ = os.Remove(c.source)
	handle, err := os.Create(c.source)
	whereErrorsGoToDie(err)
	_ = handle.Close()
}

func (c *ChallengeDatabase) bootstrap(queries []string) {
	for _, query := range queries {
		statement := c.prepare(query)
		_, err := statement.Exec()
		whereErrorsGoToDie(err)
	}
}

func (c *ChallengeDatabase) open() {
	var err error
	c.database, err = sql.Open(c.driver, c.source)
	whereErrorsGoToDie(err)
}

func (c *ChallengeDatabase) close() {
	err := c.database.Close()
	whereErrorsGoToDie(err)
}

func (c *ChallengeDatabase) prepare(query string) *sql.Stmt {
	statement, err := c.database.Prepare(query)
	whereErrorsGoToDie(err)
	return statement
}

func (c *ChallengeDatabase) insertChallenges(challenges ...Challenge) {
	query := insertChallengeRows
	var values []interface{}
	for _, challenge := range challenges {
		query = fmt.Sprintf("%s %s,", query, insertChallengeRowParams)
		values = append(
			values,
			challenge.Id,
			challenge.Name,
			challenge.Permalink,
			challenge.Author,
			challenge.Title,
			challenge.SelfText,
			challenge.CreatedUtc,
			challenge.Challenge,
			challenge.Difficulty,
		)
	}
	query = fmt.Sprintf("%s %s;", query, insertChallengeRowParams)
	statement := c.prepare(query)
	_, err := statement.Exec(values...)
	whereErrorsGoToDie(err)
}

func NewChallengeDatabase(sourceArgs ...string) *ChallengeDatabase {
	db := ChallengeDatabase{
		driver: defaultDatabaseType,
	}
	db.constructSource(sourceArgs...)
	db.createDatabase()
	db.bootstrap(bootstrapQueries)
	return &db
}
