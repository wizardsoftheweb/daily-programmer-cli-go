// Copyright 2021 CJ Harries
// Licensed under http://www.apache.org/licenses/LICENSE-2.0

package cmd

type Challenge struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Permalink  string `json:"permalink"`
	Author     string `json:"author"`
	Title      string `json:"title"`
	SelfText   string `json:"self_text"`
	CreatedUtc string `json:"created_utc"`
	Challenge  int    `json:"challenge"`
	Difficulty int    `json:"difficulty"`
}
