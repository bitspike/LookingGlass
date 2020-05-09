# CloudKarafka Manager

## About
This project is based on the [CloudKarafka/cloudkarafka-manager](https://github.com/CloudKarafka/cloudkarafka-manager) and it is a tool which intends to become the #1 OpenSrouce Kafka ecosystems monitoring and management solution.

## Motivation
After many years of involvement with the Kafka ecosystem, painfull and expensive interactions with the viarious vendors offering the Kafka as a service, it became evident that the lack of a complete management GUI for the kafka ecosystem was missing.
This tool is hopping to create a complete web interface with the following features:
* Provide GUI for monitoring and management of multiple Kafka clusters
* Provide GUI for multiple installations of Kafka Connect
* Provide GUI for multiple installations of KSQLDB instances
* Provide GUI for multiple installations of Schema Registry instances
* Provide GUI for Kafka MirrorMaker and MM2 and orchestrate cluster replication
* Organize the various attached systesm (i.e. Kafka, Kafka Connect, etc.) into groups for easier management
* Provide complete user management including SSO, RBAC, Hashicorp Vault integration
* Provide and interface which could 

## Usage

Usage of this prodcut is described by the following documents:
* [User's Guide]()
* [Administrator's Guide]()

```
Original instructions are following (to be removed)
* Download the [latest version](https://github.com/CloudKarafka/cloudkarafka-manager/releases/latest) from the releases and extract the file
* Make sure all your brokers have the [Kafka HTTP Reporter](https://github.com/CloudKarafka/KafkaHttpReporter) installed
* Start the application: `./cloudkarafka-mgmt.linux`
* Open your web browser and go to [http://localhost:8080](http://localhost:8080)
```

## Development

* Clone this repo into $GOPATH/src/github.com/CloudKarafka/cloudkarafka-manager
* Run `go get -u` to get latest dependencies
* Install the metrics reporter [Kafka HTTP Reporter](https://github.com/CloudKarafka/KafkaHttpReporter) on your local kafka broker
* Run Management interface with `go run app.go --authentication=none-with-write`

## Contributing

Please read [CONTRIBUTING.md](CONTRIBUTING.md) for details on our code of conduct, and the process for submitting pull requests to us.

## Versioning

We use simple integer versioning for this prodcut. For the versions available, see the [tags on this repository](https://github.com/bitspike/cloudkarafka-manager/tags).

