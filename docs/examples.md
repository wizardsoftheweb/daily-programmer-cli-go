# Examples

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)
[![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/)
[![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](../NOTICE)

## Single Post (first deliverable?)

Suppose you wanted to do [[2019-01-14] Challenge #372 [Easy] Perfectly balanced](https://old.reddit.com/r/dailyprogrammer/comments/afxxca/20190114_challenge_372_easy_perfectly_balanced/). You would run

```console
TBD
```

which would create the following structure in your repo

```
r-daily-programmer/
    easy/
        372/
            README.md
```
and its `README` would look like so (WIP)

```markdown
# [2019-01-14] Challenge #372 [Easy] Perfectly balanced

## Source

[Original post](https://old.reddit.com/r/dailyprogrammer/comments/afxxca/20190114_challenge_372_easy_perfectly_balanced/) by [u/Cosmologicon](https://old.reddit.com/user/Cosmologicon)

# Prompt

Given a string containing only the characters `x` and `y`, find whether there are the same number of `x`s and `y`s.

    balanced("xxxyyy") => true
    balanced("yyyxxx") => true
    balanced("xxxyyyy") => false
    balanced("yyxyxxyxxyyyyxxxyxyx") => true
    balanced("xyxxxxyyyxyxxyxxyy") => false
    balanced("") => true
    balanced("x") => false

## Optional bonus

Given a string containing only lowercase letters, find whether every letter that appears in the string appears the same number of times. Don't forget to handle the empty string (`""`) correctly!

    balanced_bonus("xxxyyyzzz") => true
    balanced_bonus("abccbaabccba") => true
    balanced_bonus("xxxyyyzzzz") => false
    balanced_bonus("abcdefghijklmnopqrstuvwxyz") => true
    balanced_bonus("pqq") => false
    balanced_bonus("fdedfdeffeddefeeeefddf") => false
    balanced_bonus("www") => true
    balanced_bonus("x") => true
    balanced_bonus("") => true

Note that `balanced_bonus` behaves differently than `balanced` for a few inputs, e.g. `"x"`.
```
