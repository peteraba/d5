Finder
======

Designed to provide search utilitites

Requires the following environment variables to be set:

 * D5_HOSTNAME: mongodb hostname
 * D5_DBNAME: mongodb db name
 * D5_COLL_WORDS: mongodb collection for words
 * FINDER_DEBUG: sets debug mode


CLI
---

Finder provides a regular unix interface

Search query is expected as standard input in JSON format

    echo "{\"word.german\": \"solche\"}" | finder


Server
------

Finder also provides a server

Requires to flags:
 * **server true**
 * **port *{portnumber}* ** (optional, 17171 by default)

    finder -server true -port 20202

Search query should be posted as JSON
