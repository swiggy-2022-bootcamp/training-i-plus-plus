# Monolithic VS Micro Service Architecture

## Monolithic 

The monolithic architecture is considered to be the traditional way of building an application.

A monolithic application is usually built as a single unit and indivisible unit, generally comprising of the following layers:

>A database, usually an RDBMS (relational database management system), that consists of many tables.

>A server-side application that handles and serves client-side requests, retrieves and stores data from/to the database and executes business logic.

>A client-side UI (User Interface) that generally consists of HTML and/or Javascript pages running on a browser.

The above layers group together to form a single logical executable. A monolithic application usually has one large codebase and lacks modularity.

If any updates are needed in the system, developers must build and deploy the entire stack at once.

## Micro Service

The microservices architecture breaks down an application into a collection of smaller independent units.

>These units carry out application processes as separate services, each of which perform specific functions and have their own business logic and database.

>The microservices architecture is an approach to developing a single application as a suite of small services.

>The entire functionality is split up into independent modules that communicate with each other through defined APIs.

>Each service has its own scope and can be deployed, updated and scaled independently. There is little to no centralized management of these services.

>In contrast to a monolithic architecture, the functionality of microservices are expressed formally with business oriented APIs.

>They encapsulate a core business requirement and the implementation of the service is hidden as the interface is defined in business terms.

>These services are adaptable for use in multiple contexts and can be reused in multiple business processes or over different channels.

Due to the application of the principle of loose coupling, dependencies between services and their consumers are minimized.

This allows service owners to roll out any updates to the implementation of the services without any impact to the users.

## Benefits

- Monolithic - Simple to develop, Simple to deploy, Easier to debug and test, Less cross-cutting, Performance

- Micro Service - Decoupled components, Understanding, Scalability, New technologies, Higher agility, Reuse