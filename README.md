# cybercert-blockchain-api

CyberCert Hyperledger Fabric API wrapper for invoke private Fabric.

## Installation

Before procced the installation steps please complete [cybercert-fabric-setup](https://github.com/handysama/cybercert-fabric-setup#how-to-setup-on-debianubuntu) beforehand.

### Prerequisite software for API

- go 1.18
- rabbitmq

### Installation for Ubuntu

Network config file is generated after deploying Fabric setup, for example `org1.yaml`.

Modify `org_hlf_config` in `config/config/yaml` to your local network config file path.

```bash
# install rabbitmq (or follow official guide depend on your platform)
./scripts/setup-rabbitmq.sh

# install dependencies
make vendor
```

## How to run locally

```bash
# Start background worker to handle issue certificate batch request in separate terminal
go run worker/main.go

# Start blockchain api server
make run
```

### Examples invoke endpoints

```bash
# Issue Certificate
curl -v -X POST \
    -H "Content-Type: application/json" \
    -d @samples/certinfo-issue-cert.json \
    127.0.0.1:10788/api/certificate/info/issue

# Revoke Certificate
curl -X DELETE \
    -H "Content-Type: application/json" \
    -d @samples/certinfo-revoke-cert.json \
    127.0.0.1:10788/api/certificate/info/revoke

# Query Records (search certificates by query string, for example by CourseId)
curl -X POST \
    -H "Content-Type: application/json" \
    -d @samples/certinfo-query-records.json \
    127.0.0.1:10788/api/certificate/info/query-records

# Query Certificate
curl -X POST \
    -H "Content-Type: application/json" \
    -d @samples/certinfo-query-cert.json \
    127.0.0.1:10788/api/certificate/info/query-certificate

# Get Modification History of Certificate
curl -X POST \
    -H "Content-Type: application/json" \
    -d @samples/certinfo-get-history.json \
    127.0.0.1:10788/api/certificate/info/history
```
