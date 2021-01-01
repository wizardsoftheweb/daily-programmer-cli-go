// Copyright 2021 CJ Harries
// Licensed under http://www.apache.org/licenses/LICENSE-2.0

package cmd

const queryCreateTableChallenges string = `
CREATE TABLE IF NOT EXISTS challenges(
id TEXT NOT NULL PRIMARY KEY,
name TEXT NOT NULL,
permalink TEXT NOT NULL,
author TEXT NOT NULL,
title TEXT NOT NULL,
self_text TEXT NOT NULL,
created_utc TEXT NOT NULL,
challenge INTEGER DEFAULT 0,
difficulty INTEGER DEFAULT 0
)
;
`

const queryCreateTableDifficulties string = `
CREATE TABLE IF NOT EXISTS difficulties(
id INTEGER NOT NULL PRIMARY KEY,
difficulty TEXT NOT NULL
)
;
`

const querySeedDifficulties string = `
INSERT INTO difficulties(id, difficulty)
VALUES
    (0, "unknown"),
    (1, "easy"),
    (2, "intermediate"),
    (3, "hard")
;
`

const insertChallengeRows string = `
INSERT INTO
    challenges(
       id,
       name,
       permalink,
       author,
       title,
       self_text,
       created_utc,
       challenge,
       difficulty
    )
VALUES
`

const insertChallengeRowParams string = " (?, ?, ?, ?, ?, ?, ?, ?, ?)"

var bootstrapQueries = []string{
	queryCreateTableChallenges,
	queryCreateTableDifficulties,
	querySeedDifficulties,
}
