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

## Golang Stuff

### Collecting Challenges

Reddit and its APIs work newest to oldest. To do the challenges in order, we need to collect all the posts, from oldest to newest. Since we'll (possibly) need to walk the tree multiple times, we should store that info. The data is small enough we could probably hold it all in memory but that won't persist. We might end up storing this in the MD but, at least for development, it needs to be accessible.

We're not really talking anything massive here. To me, a JSON file would be rad. [Scribble](https://github.com/sdomino/scribble) seems to be the one Go package that does this. It's got a few forks. Because none of them are well-maintained, I'm not sure they're a great idea. Plus the first TODO, `Support for windows`, kinda kills the deliverable. While I find it impossible to reliably develop on Windows, I know people that like it just fine.

SQLite is next best thing and that means [`go-sqlite3`](https://github.com/mattn/go-sqlite3). But if we're getting into running an actual database, my first thought is portability. I don't remember how compatible SQLite queries are with anything else. That leads me to wonder about an ORM but that's just layers more complexity than I really want to add. In a web app, ORMs slow down response time because they're adding extra code to go through before returning. I'd assume the same applies for compiled languages but I wonder how much, if any, that truly affects things.

Since we're stuck with SQL (NoSQL doesn't make sense here because we have structured data and absolutely zero need to scale), we need to define what we want to collect for each post. "ERD" seems like overkill because we really only have one table, `challenges`, that contains everything. We know we want to search on challenge number, difficulty, and (because people are weird)  date. Those give us our index columns. Note a challenge's primary key, say, `id`, cannot be its challenge number as many challenge numbers have multiple difficulties. I suppose we could split `challenges` into three tables, `{easy,intermediate,hard}_challenges`, but then we'd have hole-y data as not all challenge numbers have all difficulties and then we'd have to `JOIN` for some searches. Again, that's complicated.

GRAW gives us [a nice `Post` interface](https://godoc.org/github.com/turnage/graw/reddit#Post). It gives us `Post.ID`, `Post.Name`, and `Post.Permalink` that are probably useful GRAW/reddit metadata. Additionally `Post.ID` gives us a UUID that we can use as a primary key which is rad (even though the name doesn't use idiomatic Pascal/camelCase). For attribution, we definitely want `Post.Author`; `Post.URL` looks useful but it's easy to reconstruct from `Post.Permalink` (plus it doesn't follow idiomatic Pascal/camelCase). From the post itself we really need its title and contents. `Post.Title` and `Post.SelfText` will give us the title and reddit-markdown (which we can spin to normal MD). Finally, `Post.CreatedUTC` gives us that date we (kinda don't) want as well. The rest of the info doesn't look like anything we really care about if we're going to solve the challenge ourselves and socialize our solutions.

SQLite3 has [only 5 data types](https://sqlite.org/datatype3.html). For our purposes, we really only need `TEXT` (how weird is that?). Were we using, say, MySQL, we'd probably also want `VARCHAR` and `DATETIME`. I don't know; I'm not a DBA. We don't want to insert a challenge if any of the necessary info is missing ergo we can use `NOT NULL` everywhere. Technically we don't need it for our `PRIMARY KEY` but readable code is better than technical code.

```sqlite
CREATE TABLE IF NOT EXISTS challenges(
    id TEXT NOT NULL PRIMARY KEY,
    name TEXT NOT NULL,
    permalink TEXT NOT NULL,
    author TEXT NOT NULL,
    title TEXT NOT NULL,
    self_text TEXT NOT NULL,
    created_utc TEXT NOT NULL
);
```

On second thought, this excludes some the metadata we want. We're missing challenge number and difficulty. That will have to be parsed from the post data. However, until we collect all the `Post`s, we won't know if we can parse that info from everything. This means we need a default value. Redundancy is good; while we could do that in Go (and should) we can also do that in SQLite. Challenge numbers are `INTEGER`s. If we use `DEFAULT 0`, we'll be able to quickly find all challenges whose challenge number could not be parsed. Difficulties are strings but we're going to be lazy and use `INTEGER` instead; we could use a mapping table but for now we'll just infer `1 == easy`, `2 == intermediate`, and `3 == hard`. If we use `DEFAULT 0` we'll then be able to quickly find all challenges whose difficulty could not be parsed. Having a `difficulties` table seems like overkill. I like it because it moves responsibility to the data instead of some hardcoded strings in the Go.

```sqlite
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

CREATE TABLE IF NOT EXISTS difficulties(
    id INTEGER NOT NULL PRIMARY KEY,
    difficulty TEXT NOT NULL
)
;

INSERT INTO difficulties(id, difficulty)
VALUES
    (0, "unknown"),
    (1, "easy"),
    (2, "intermediate"),
    (3, "hard")
;
```


## Ignore me

This is all stuff @thecjharries wrote when he was being dumb. You can ignore it or, if you read it, laugh at his stupidity.

### Finding the earliest challenge

GRAW really wants [some starting point](https://godoc.org/github.com/turnage/graw/reddit#Scanner). Using [redditsearch.io](https://redditsearch.io) and the creation date of the subreddit (2012/02/09), we can poke around for [the earliest posts](https://redditsearch.io/?term=&dataviz=false&aggs=false&subreddits=dailyprogrammer&searchtype=posts&search=true&start=1328767200&end=1328853600&size=100). It looks like the first challenges were posted in order of difficulty, so [[easy] challenge #1](https://www.reddit.com/r/dailyprogrammer/comments/pih8x/easy_challenge_1/) is most likely our first challenge we can use to seed GRAW.

Reddit search surprisingly [returns similar results](https://old.reddit.com/r/dailyprogrammer/search/?q=challenge+%231&restrict_sr=on&t=all&sort=relevance), something I didn't expect.

This yields a `Post.Name == t3_pih8x`.
