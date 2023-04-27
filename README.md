# Go Transaction Processing Program

This Golang program reads a CSV file, processes the data, saves it to a PostgreSQL database, and sends an email with a summary of the file's transactions.
Requirements

* golang
* Docker
* Docker Compose

To make it easier to set these environment variables, it is recommended that you edit the * docker-compose.yml * file. In this file are all the configuration options needed to start and run the program in a Docker container. By editing this file, you can customize the environment variables used in the program to suit your specific needs.

### Run the following command to start the container:

``` sh 
docker-compose up -d 
```

The destination email and the file name can be modified in the main file

``` go
var destinationEmail = "testemailchallenge@hotmail.com"
var fileName = "transactions.csv"

```


Once the processing is complete, an email will be sent with the results.