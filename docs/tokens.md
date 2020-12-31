# Necessary Tokens

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/) [![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](../NOTICE)

## GitHub

1. Visit [the new token page](https://github.com/settings/tokens/new)
2. Give it the `public_repo` scope

   ![`public_repo` scope](./images/pat-public_repo-scope.png)
3. Copy the token somewhere

   ![copy new PAT](./images/copy-new-pat.png)
4. Export it as `WOTW_DAILY_PROG_GITHUB_TOKEN` in your shell. That's verbose but it's good design and you're only copypasting it once so why do you care?

   ```shell
   $ env | grep WOTW_DAILY_PROG_GITHUB_TOKEN
   WOTW_DAILY_PROG_GITHUB_TOKEN=qqq123
   ```

## Reddit

TBD
