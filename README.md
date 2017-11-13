## JSON Object NoSQL Database Container

A file system database with an in-memory cache and a RESP API for table/item management

The database is capable of storning items of arbitrary format as long as they are valid JSON and include the string field "id" as part of the item. This field is intended to be the object's primary key within the table.
