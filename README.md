## JSON Object NoSQL database container with a REST API

### Features:

* Create tables by specifiying a name for the table and the name of the primary key as it appears on the JSON object
* Get a list of the database's tables
* Delete tables by name
* Get a table's contents by table name
* Get an item within a table given the table name and primary key

We take advantage of Golang's recursive datatype the map[string]interface{} which is the struct type of all JSON objects 
