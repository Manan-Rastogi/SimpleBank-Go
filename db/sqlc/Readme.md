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

