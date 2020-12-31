// Copyright 2020 CJ Harries
// Licensed under http://www.apache.org/licenses/LICENSE-2.0

package main

import (
	"github.com/wizardsoftheweb/daily-programmer-cli/cmd"
)

func main() {
	err := cmd.Execute()
	whereErrorsGoToDie(err)
}

func whereErrorsGoToDie(err error) {
	if nil != err {
		panic(err)
	}
}
