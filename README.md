# Golang-Serverless Application with DynamoDB and AWS Lambda

This repository contains a simple serverless application built with GoLang, using AWS Lambda for compute and DynamoDB for storage.

## Features

- AWS Lambda for serverless compute
- DynamoDB for NoSQL database storage
- API Gateway for HTTP endpoints
- Validation of incoming requests' user info
- Modular structure for easy maintenance and scalability

## Prerequisites
 - GoLang 1.16+
 - AWS CLI configured with appropriate permissions
 - AWS SAM CLI for local testing and deployment
 - DynamoDB table set up in AWS

## Installation

Cone the repository:

```bash
git clone https://github.com/codeabuu/AWS-Serverless.git
cd AWS-Serverless

```
### Install Dependecies

```bash
go mod tidy
```

## Deployment

Deploy Go Lambda functions with .zip file archives using this detailed [documentation](https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html)
