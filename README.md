# Go Transaction Processing Program

This is a Go transaction processing program that uses scalable goroutines to process large volumes of transaction data. The program opens a CSV file, processes the transaction data, calculates the credit total and average, saves the information to a PostgreSQL database, and sends an email with the results.

Requirements

* golang
* Docker
* Docker Compose

To make it easier to set these environment variables, it is recommended that you edit the * docker-compose.yml * file. In this file are all the configuration options needed to start and run the program in a Docker container. By editing this file, you can customize the environment variables used in the program to suit your specific needs.

### Run the following command to start the container:

`` docker-compose up -d ``


Once the processing is complete, an email will be sent with the results.