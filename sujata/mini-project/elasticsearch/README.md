## Elastic Search

Elastic Search is used for fuzzy searching of the products by user of Role -> BUYER. The following steps are involved for implementing this feature in `Search` service.

1. User (BUYER) will search for the product which will be passed as the query string along with JWT token issued for that user to `Search` service. <br>
   > URL: <search_service_host:port>/v1/search?product=<searched_product>
2. `Search` service will check the do the authentication using the public key for JWT token of the user and will authorize its ROLE of type `BUYER`.
3. Elastic Search will do fuzzy matching on name and description of the documents and will return the related documents back to `Search` service.
4. `Search` service will then send these documents and related info to the user.

## Data syncing from Mongodb to elastic search

### Prerequisties

- ElasticSearch should be set up locally and running.
- Logstash configured locally.

Logstash is used for the purpose of data syncing from mongo db to elastic search. Below are the steps to setup logstash locally and run `conf` file for data syncing. Conf file used for this project is present in this directory only under folder logstash with name `mongodata.conf`.

1. Install the tar file of logstash from their site - [Download Logstash](https://www.elastic.co/downloads/logstash)
2. Move the above file to suitable place like Documents and do tar -xvf tarfilename. That will decompress it.
3. Provide the above extracted folder permission to access certain folders that will be specified in conf file by using chmod 777 <dir_name>
4. Create the logstash conf file by running the command `sudo su` and then, `mkdir /etc/logstash/conf.d/mongodata.conf`
5. Now create a new directory using `sudo su` and `mkdir /opt/logstash-mongodb` for `placeholder_db_dir`.
6. Starting logstash file command - `bin/logstash --debug -f /etc/logstash/conf.d/mongodata.conf` <br>
   `--debug` flag helps in debugging by providing extra information.
