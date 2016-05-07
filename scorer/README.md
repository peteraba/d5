Scorer
======

Saves game play results.

Requires the following environment variables to be set:

 * D5_DBHOST: mongodb hostname
 * D5_DBNAME: mongodb database name


CLI
---

Scorer provides a regular unix interface

Requires the following flags:

 * **coll *{db collection name}* **
 * **debug false ** (optional, false by default)

```bash
scorer --coll=german --data='{"wordId":"aabadih29a", "score": 10}'
```


Server
------

Finder also provides a server

Requires the following flags:

 * **coll *{db collection name}* **
 * **debug false ** (optional, false by default)
 * **server true**
 * **port *{port number to listen to}* ** (optional, 17171 by default)

```bash
scorer --coll=german --server=true --port=20202 --coll=german
```

Scores must be sent via POST with wordId and score form fields.

