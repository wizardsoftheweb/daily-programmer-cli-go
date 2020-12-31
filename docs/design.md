# Design Overview

[![License: Apache 2.0](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)  [![License: CC BY 4.0](https://img.shields.io/badge/License-CC%20BY%204.0-lightgrey.svg)](https://creativecommons.org/licenses/by/4.0/) [![Check the NOTICE](https://img.shields.io/badge/Check%20the-NOTICE-420C3B.svg)](../NOTICE)

## Deliverable

As with most of my projects, my goal is to provide a simple binary that can be downloaded and run without dependencies. That means my choice of language is either Go or Rust (Python is rad but it's a serious pain to build a static binary). I don't know Rust (yet) so I'm starting with Go. Might eventually refactor to Rust but, as you'll see below, that might mean taking on more work to either update an API client or just use an HTTP library.

## Reddit API

[API docs](https://www.reddit.com/dev/api/)

### Library

[`graw`](https://github.com/turnage/graw) seems to be a (mostly?) functional Go wrapper for the Reddit API. Looks neat and does what it says on the tin.

#### Alternatives

As I mentioned above, this is (currently) written in Go. Here's what else I researched during language selection.

* Python: [PRAW](https://praw.readthedocs.io/en/latest/) is the definitive Python tool. I've used it before and I really like it.
* Rust: There doesn't appear to be a complete or active Rust API wrapper. [Orca](https://github.com/IntrepidPig/orca) and [RAWR](https://github.com/Aurora0001/rawr) are mentioned as decent clients but aren't receiving active updates.
