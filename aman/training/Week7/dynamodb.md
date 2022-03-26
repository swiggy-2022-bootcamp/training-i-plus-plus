# DynamoDB

DynamoDB is a hosted NoSQL database offered by Amazon Web Services which offers reliable performance even when it scales and gives a managed experience and a small, simple API allowing for simple key-value access as well as more advanced query patterns.

## Advantages

- Scalable- Virtual unlimited storage, users can store infinity amount of data according to their need
- Cost Effective- It seems to be cutting costs, while a big part of data is able to migrate from SQL to NOSQL. Basically it charges for reading, writing, and storing data along with any optional features you choose to enable in DynamoDB
- Data Replication- All data items are stored on SSDs and replication is managed internally across multiple availability zones in a region or can be made available across multiple regions
- Serverless- DynamoDB scales horizontally by expanding a single table over multiple servers
- Easy Administration- Amazon DynamoDB is a fully managed service, you don't need to worry about hardware or software provisioning, setup & configuration, software patching, distributed database cluster or partitioning data over multiple instances as you scale
- Schemaless DB
- Secure- Customizable traffic filtering, Regulatory Compliance Automation, Comprehensive Database Threat Detection, Advanced System of Notification and Reporting

## Issues

- Weak querying model, querying data is extremely limited.
- Lack of server-side scripts. T
- Table Joins - Joins are impossible
- Hard to predict cost when that usage might spike
- No client-controlled transactions