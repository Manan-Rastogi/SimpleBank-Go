# Understanding SQLC Structure.

```
After running queries 1st file made is models.go which contains structures of our DB tables.
```

```
2nd file is db.go, which contains DBTX interface. It defines 4 common methods that both sql.DB and sql.Tx object has. This allows us to freely use either a db or a transaction to execute a query.


the New() function takes a DBTX as input and returns a Queries object. So we can pass in either a sql.DB or sql.Tx object

method WithTx() - allows a Queries instance to be associated with a transaction.
```

```
accounts.sql.go is the file made from our query.

At top there is create account SQL query.
CreateAccountParams struct, which contains all columns that we want to set when we create a new account.

The CreateAccount function is defined as a method of the Queries object.

Similar to this are other queries and files.
```

## DB Transaction

#### Successful Transaction
```
START;
    QUERIES
COMMIT;
```

#### Failed Transaction
```
START;
    QUERIES
ROLLBACK;
```

```
The file store.go contains transaction function. Then file store_test.go contains test for the same transactions.
```

## Try this -- For seeing how transactions works
```
Open 2 tabs for running sql and run this transaction line by line parallelly.

BEGIN;
select * from accounts where id = 1;
ROLLBACK;

-- the 2 transactions will proceed that we do not want.

```
```

-- Let's check this with FOR UPDATE ---          SPOILER - 2nd transaction will be blocked.

BEGIN;
select * from accounts where id = 1 for update;
update accounts set balance = 10001 where id = 1;
COMMIT;
```

## Lock Monitoring
#### https://wiki.postgresql.org/wiki/Lock_Monitoring
```
>> Get blocked Statement.

SELECT
    blocked_locks.pid AS blocked_pid,
    blocked_activity.usename AS blocked_user,
    blocking_locks.pid AS blocking_pid,
    blocking_activity.usename AS blocking_user,
    blocked_activity.query AS blocked_statement,
    blocking_activity.query AS current_statement_in_blocking_process
FROM
    pg_catalog.pg_locks blocked_locks
    JOIN pg_catalog.pg_stat_activity blocked_activity ON blocked_activity.pid = blocked_locks.pid
    JOIN pg_catalog.pg_locks blocking_locks ON blocking_locks.locktype = blocked_locks.locktype
    AND blocking_locks.database IS NOT DISTINCT
FROM
    blocked_locks.database
    AND blocking_locks.relation IS NOT DISTINCT
FROM
    blocked_locks.relation
    AND blocking_locks.page IS NOT DISTINCT
FROM
    blocked_locks.page
    AND blocking_locks.tuple IS NOT DISTINCT
FROM
    blocked_locks.tuple
    AND blocking_locks.virtualxid IS NOT DISTINCT
FROM
    blocked_locks.virtualxid
    AND blocking_locks.transactionid IS NOT DISTINCT
FROM
    blocked_locks.transactionid
    AND blocking_locks.classid IS NOT DISTINCT
FROM
    blocked_locks.classid
    AND blocking_locks.objid IS NOT DISTINCT
FROM
    blocked_locks.objid
    AND blocking_locks.objsubid IS NOT DISTINCT
FROM
    blocked_locks.objsubid
    AND blocking_locks.pid != blocked_locks.pid
    JOIN pg_catalog.pg_stat_activity blocking_activity ON blocking_activity.pid = blocking_locks.pid
WHERE
    NOT blocked_locks.granted;
```

```
>> Lock Monitoring

SELECT
    a.datname,
    a.application_name,
    l.relation :: regclass,
    l.transactionid,
    l.mode,
    l.locktype,
    l.GRANTED,
    a.usename,
    a.query,
    a.query_start,
    age(now(), a.query_start) AS "age",
    a.pid
FROM
    pg_stat_activity a
    JOIN pg_locks l ON l.pid = a.pid
WHERE
    a.application_name = 'psql'
ORDER BY
    a.pid;
```

## Where we are getting lock and how to solve it?
```
-> We are getting lock in these 2 queries when running parallely. This might be different for different users.

Insert transfer for tx2;
Select account for tx1;

-> If we remove referential integrity between these two tables (By removing foreign key constraint), we will not get the lock, but the solution is not efficient, as we loose the referential integrity between the table.


-> Another Observation is account id will never be edited/updated so if we can tell this to postgres that during update no key will be updated for account table we can avoid this deadlock situation.

Doing so is simple - update in account.sql
SELECT * FROM accounts
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

and then make sqlc

```
