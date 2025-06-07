#
# What is PostgreSQL
is an object-relational database management system(DBMS) 
that means it is a system form managing data stored in relations (table) 
# Architecture
uses client/server model.
- A server process: which manage the database files, accepts connections from the client application and performs actions on the database on the behalf of the clients. called postmaster.
can handle multiple concurrent connections from clients, for that purpose it starts ("forks") a new process for each connection.
- The user's client (FE) application that want to perform DB operations. the client could be a web server, a graphical application, a text-oriented tool.

[client] ----TCP/IP---- [server]

# psql cli

```sql
CREATE DATABASE 
DROP mydb

$ psql mydb
$ psql -s mydb -- s: for single step mode
select version();
select current_date;

\q : quit
\i basics.sql : import sql ifle to a DB
```
