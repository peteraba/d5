Persister
=========

Designed to persist JSON data in mongodb

Requires the following environment variables to be set:

 * D5_HOSTNAME: mongodb hostname
 * D5_DBNAME: mongodb database name


CLI
---

Persister provides a regular unix interface

Data to be stored is expected as standard input in JSON format

Requires the following flags:

 * **coll *{collectionName}* **

```bash
cat persister/fixture/gerdict.json | persister -coll german
```

