# Mini Project - ShopKart - Online Shopping Store

[In Progress] ShopKart Diagram

![In Progress Architecture](./static/images/Shop_kart%402x.png)

## Services

- **Auth**
- **Product**
- **Search**
- **Cart**
- **Order**
- **User**

## Tech Stack

- Golang
- Sonarqube
- JWT
- ElasticSearch
- Logstash
- Mongodb

## Auth Service

It uses private/public key for JWT creation. It uses private key to sign the JWT token and public key can be used to authenticated the JWT token by other services.
To generate public and private keys, run the `mkdir keys` command inside `auth/cmd/server`. Then, run the following commands inside `auth/cmd/server/keys` to generate the public and private keys for local configuration/testing of the service. <br>

1. Command to generate RSA private key:
   `openssl genrsa -out app.rsa 2048`

2. Command to generate RSA public key:
   `openssl rsa -in app.rsa -pubout > app.rsa.pub`

This service mainly does sign up of new user and sign in of the user by returning back the JWT token signed by private key.

## Product Service

This service allows the following endpoints

1. POST `product/v1/product` - User with Role `SELLER` can add the product to the mongo database. It accepts product information in the request body and JWT token in authorization header for authentication as well as authorisation purpose.

## Search Service

Elastic search and `Search` service working is documented inside the `elasticsearch` directory. It also contains info/related documents for data syncing with logstash between mongodb and elastic search.

## Cart Service

This service allows user of type `BUYER` to add a product to his/her cart, delete a product from the cart and then, can fetch all the products of the cart using GetCart API.

## Steps to run Sonarqube

- Run the following command in root of the project/microservices to generate coverage file.<br>
  > `go test -short -coverprofile=./cov.out ./...`
- Run the following command in root of the project/microservices to generate gosec report.<br>
  Install gosec from [Gosec Github link](https://github.com/securego/gosec)<br>
  > `gosec -fmt=sonarqube -out report.json ./...`
- SonarQube properties file is already present in the root of each service directory.
- Set the project in SonarQube and provide the same name in `sonar.projectKey` field and running address of SonarQube in `sonar.host.url` field of `sonar-project.properties`.
- Provide the generated authentication token by SonarQube in `sonar.login` field of `sonar-project.properties`.
- Run the `sonar-scanner` command to run sonar scanner and visit to dashboard for the analysis of the codebase.

## References

- Requirement Document - [Online Shopping Store](https://docs.google.com/document/d/1cnCHEVkOgFDYSmZmSbxcDlZiLjCZXr1W9jHf62id7T8/edit?usp=sharing)
- Combined Fields in Elastic Search - [Official Documentation](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-combined-fields-query.html)
