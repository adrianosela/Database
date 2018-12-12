## JSON Object NoSQL Database Container

A file system database with an in-memory cache and a REST API for table/item management

The database is capable of storing items of arbitrary format as long as they are valid JSON and include the string field "id" as part of the item. This field is intended to be the object's primary key within the table

The point of this project was to create a variable schema DB for testing. The DB is ephemeral and is not suitable for any real usecase other than perhaps a runtime cache.
