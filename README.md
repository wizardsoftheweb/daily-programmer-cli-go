# `daily-prog`

[![CircleCI](https://img.shields.io/circleci/build/github/wizardsoftheweb/daily-programmer-cli-go/dev)](https://circleci.com/gh/wizardsoftheweb/daily-programmer-cli-go/tree/dev)
[![Coverage Status](https://img.shields.io/coveralls/github/wizardsoftheweb/daily-programmer-cli-go/dev)](https://coveralls.io/github/wizardsoftheweb/daily-programmer-cli-go?branch=dev)
[![GoDoc](https://godoc.org/github.com/wizardsoftheweb/daily-programmer-cli-go?status.svg)](https://godoc.org/github.com/wizardsoftheweb/daily-programmer-cli-go)

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/) [![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](./NOTICE)

## Overview

This project aims to set up some organization and flow around using [r/DailyProgrammer](https://old.reddit.com/r/dailyprogrammer).

The problems from r/DailyProgrammer are useful demonstrations of coding skills and, over time, lend themselves to (semi?) useful code samples. Socializing solutions and tracking progress is annoying through reddit but much easier through VCS. For now, the project only supports GitHub.

This repo (at V1 or later) provides a tool that does the following things:

1. Creates a local directory structure for solution output
2. Links the local structure to GitHub
3. Seeds each directory with a `README.md` that
    * links to the given challenge
    * pulls in the challenge

See [the examples](./docs/examples.md) for more info.

## Features

See [Roadmap](#Roadmap) for future features and [Ideas](#Ideas) for things that might be features eventually

## Roadmap

1. Document necessary tokens and setup steps for usage
    1. Using opinionated structure
    2. Using config structure
2. Create a new repo
    1. Using opinionated structure
    2. Using config structure
    3. With...
        * `.gitignore`
        * skeleton directories?
        * `README` boilerplates
3. Set up and document file structure for the repo
    1. Using opinionated structure
    2. Using format string for structure
4. Set up and document how to connect to Reddit API
5. Search r/DailyProgrammer prompts
    1. Search by date
    2. Search by number
    3. Search by difficulty
        1. with date
        2. with number
6. Download specific r/DailyProgrammer prompts
    1. Download specific date
    2. Download specific number
    3. Download specific difficulty
        1. from date
        2. from number
7. Download all r/DailyProgrammer prompts

## Ideas

* Integrated testing
* Keyword challenge search
    * In titles
    * In descriptions
* Other VCS
    * GH Cloud
    * BitBucket + cloud
    * GitLab + cloud
    * Others?
