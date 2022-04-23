Introduction
============

Way before NoSql databases hit the scene, application developers/organizations used relational database management systems(RDBMS) to store and retrieve data. As their applications got bigger, data manipulation got complex and expensive. Some of the issues faced were:

-   Writing joins and aggregate functions to retrieve data from multiple tables greatly affected performance and was really expensive.
-   Changing the underlying data model for the application once it was deployed was really difficult, if not impossible because tables had fixed rows and columns.

-   Scaling an RDBMS was a pain. They can only be scaled vertically. (scale-up with a larger server). That's expensive and resources can easily be wasted.

With all these issues, development wasn't rapid, the cost of running the service was high and scaling was difficult. In the early 2000s, NoSQL was born.\
NoSQL is a non-relational DBMS, that does not require a fixed schema, avoids joins, and is easy to scale. It stores data as JSON documents instead of as columns and rows used by relational databases. NoSQL stands for "not only SQL" rather than "no SQL" at all. This means a NoSQL JSON database can store and retrieve data using literally "no SQL." Or you can combine the flexibility of JSON with the power of SQL for the best of both worlds.

Some of the technical challenges NoSql addresses are

1.  More customers are going online. So a system has to scale to support thousands if not millions at a time.
2.  Applications are moving to the cloud.NoSql databases minimize infrastructure costs, achieving a faster time to market.
3.  The world has gone mobile. Creating "offline first" apps are required.
4.  Many devices need to connect at once.Computers, mobile devices, IoT, etc. Seeing the advantages NoSQL DB's bring to the table caused a lot of organizations to migrate their data. Big companies such as Amazon, Microsoft, Google offer solid, highly performant, cost-effective, and scalable NoSQL databases to host Terrabytes of data with ease.\
    Today, we'll be focusing on dynamo DB from Amazon

AWS DynamoDB
------------

It's a fully managed NoSQL database service. It's fast, has predictable performance, and highly scalable. As a matter of fact, it has near-infinite scaling, enabling the developer not to worry about performance bottlenecks as their application grows.\
It's available over HTTP API or HTTPS endpoints, providing a simple, secure interaction model with your database.

### Core Features

-   Create tables that can store and retrieve any amount of data
-   Serve any request traffic.
-   Scale Up or scale down your tables throughput capacity without downtime or performance degradation.
-   On-demand Backup for long term retention

### Core Components

Similar to other database systems, DynamoDB stores data in tables.\
A table is a collection of items and each item is a collection of attributes.\
Tables: A collection of items(data)\
Items : A single item is made up of a group of attributes\
Items are similar to rows in other database systems\
Attributes : An attribute is a fundamental data element, something that does not need to be broken down any further. For Example,\
Let's assume we create a table called "products". Each item in the table represents a product. And the attributes for a product can be

-   ProductID,
-   Name,
-   Description,
-   Price,
-   Quantity etc.

Primary Key\
A unique identifier for each item in a table. When creating a table, you must specify the table name and primary key. In the Products table above, ProductID is the primary key. It should be unique to that item alone.

Dynamo DB supports 2 types of primary keys. A partition key and a partition and sort key. A partition key is a simple primary key composed of one attribute. Dynamo Db uses this partition key as input to an internal Hash function. The output determines the partition where the item would be stored. No 2 items can have the same partition key. In the "Products" table above ProductID is an example of a simple primary key.\
A Partition Key and Sort Key also known as composite primary keys has 2 attributes. In this scenario, two items can have the same partition key. All identical partition key values are stored together in sorted order by sort key value.

The following snippet shows a table named Products, with some examples of items and attributes

COPY

COPY

COPY

```
#Products table
{
"ProductID":001,
"Name":"coffee",
"Qty":605,
"price":500
},
{
"ProductID":002,
"Name":"Milk",
"Qty":200,
"Price":150,
"Category":"proteins",
"Img":"https://githubusercontent.com/41929050/61567049-13938600-aa33-11e9-9c69-a4184bf8e524.jpeg"
},
{
"ProductID":003,
"Name":"Oranges",
"Qty":1000,
"Price":50,
"Category":"Fruits",
"Compliments":{
"fruits":["bananas","pineapples","guavas"],
"processed":"Yoghurt"
},
"Img":"https://githubusercontent.com/41929050/61567049-aa33-11e9-9c69-a4184bf8e57856.jpeg"
}

```

Note the following about the Products table:

Each item in the table has a unique identifier, or primary key, that distinguishes the item from all of the others in the table. In the Products table, the primary key consists of one attribute (ProductID).

Other than the primary key, the Products table is schemaless, which means that neither the attributes nor their data types need to be defined beforehand. Each item can have its own distinct attributes.

Most of the attributes are scalar, which means that they can have only one value. Strings and numbers are common examples of scalars.

Some of the items have a nested attribute (Compliments). DynamoDB supports nested attributes up to 32 levels deep.

Secondary Indexes\
It's a way of efficiently accessing records by means of some piece of information other than the usual primary key. You can create one or more secondary indexes on a table.\
DynamoDB supports 2 kinds of indexes

Global secondary index

-   An index with a partition key and sort key that can be different from those on the table.\
    Local secondary index

-   An index that has the same partition key as the table, but a different sort key. Each dynamo DB table has a quota of 20 Global and 5 local secondary keys.

#### DynamoDB Streams

It's an optional feature. It lets you capture data modification events on your table in real-time. For example, in a chat application, you'll want to know in real-time when a new message has been added to the database so that you update the user chat screen accordingly. DynamoDB Streams writes a stream record whenever one of the following events occur

-   A new item is added to the table
-   An item is updated
-   An item is deleted from the table\
    DynamoDb streams are best used with lambda functions to create code that runs automatically when an event occurs.

You can read a ton more about dynamo DB straight from the AWS website [AWS DynamoDB](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Introduction.html)

### Setting Up DynamoDB

AWS provides 2 versions of Dynamo DB. One is a downloadable version that is great for development and testing applications locally, and another is a web service, which is great for development and testing online.\
We will be using the downloadable version for this tutorial since we are developing it locally. I'll go ahead to download it and install it on my computer. Access this [link](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html) to grab a copy for your OS and install it.

*To run DynamoDB on your computer, you must have the Java Runtime Environment (JRE) version 8.x or newer. The application doesn't run on earlier JRE versions.*

*Also make sure you install aws cli on your computer*

Once you've installed and configured dynamo DB locally as illustrated [here](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/DynamoDBLocal.DownloadingAndRunning.html), running this command

COPY

COPY

COPY

```
aws dynamodb list-tables --endpoint-url http://localhost:8000

```

yield this ouput.(assuming it's a fresh dynamoDB installation)

COPY

COPY

COPY

```
{
    "TableNames": []
}

```

The article is already getting very long. I'll end here and continue in the next chapter, which illustrates how to build CRUD functions using python to store and access data from a table in Dynamo DB.