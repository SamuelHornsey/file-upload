# File Upload

Golang file upload example. Essentially this is a mini web browser based file drop. I may add user logins later.

Database DDL

```sql
CREATE TABLE files
(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    fileName TEXT,
    hash TEXT
);
```