# Ideas that come to me during development of this Go HTTP Server from scratch 
 - to avoid rabbit holes not DIRECTLY pertaining to http server we are building

---
## Find out what C Go is 
 i had something to write here damnit lol -- see couldnt have been worth getting off topic of the project/topic
 - it was WHAT IS "CGO":

    Postgres (pure Go): https://github.com/jackc/pgx [*]
    Postgres (pure Go): https://github.com/lib/pq [*]
    Postgres (uses cgo): https://github.com/jbarham/gopgsqldriver

    guess: just code written with c and go

---
## Learn what a Context is and how to use it well in Go and what it is used for

## What is DSN and ODBC

## Connection Pools?
Connection pool sizing is a deep topic — there's a famous HikariCP (Java connection pool) post called "About Pool Sizing" that applies directly to any language. The tldr is that more connections is not always better, and the optimal number is often surprisingly small. Save it for later.