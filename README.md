# Integration suite
https://barkardk.github.io/integration_suite/  
Small set of applications designed to run in kubernetes integration tests or e2e test. 

## Mariadb
 Run a set of table creation and a few simple queries to verify database response the following setups are supported
 - Mariadb standalone
 - Mariadb galera cluster

## Mysql 
Run a set of table creation and a few simple queries to verify database response the following setups are supported

## Mongodb 
Run a set of table creation and a few simple queries to verify database response the following setups are supported 

## Postgres 
Run a set of table creation and a few simple queries to verify database response the following setups are supported 
the following options are supported
- Plain database
- Timescaledb addon 

## Rabbitmq
RabbitMQ integration is a small test suite to test rabbitmq server installations.  
It works by connecting to a rabbitmq server via a provided AMQP string, it will then create a queue , post a message and consume the message.   


