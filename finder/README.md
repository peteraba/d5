Finder
======

Designed to provide search utilitites

Requires the following environment variables to be set:

 * D5_HOSTNAME: mongodb hostname
 * D5_DBNAME: mongodb database name


CLI
---

Finder provides a regular unix interface

Search query is expected as standard input in JSON format

Requires the following flags:

 * **coll *{collectionName}* **

```bash
echo "{\"word.german\": \"solche\"}" | finder --coll german
```


Server
------

Finder also provides a server

Requires the following flags:

 * **coll *{collectionName}* **
 * **server true**
 * **port *{portnumber}* ** (optional, 17171 by default)

```bash
finder --coll german --server true --port 20202
```

Search query should be posted as JSON
