# fluffy-doodle

Protoype for a payment API

## Installation

```linux
git clone git@github.com:Syntax753/fluffy-doodle.git
```

## Usage

The primary part of this is a chi router which can be lanched with:

```go
go run .
```

Alternatively can build and run the binary

## Project Layout

router.go

    This is the primary runtime

/schema

    Contains the full data set of payments as well as a test data set

/model/db

    This is an in memory database which can load the json formatted payments under schea as a map. This would be db.sql implementation but have not created an actual db schema for the purposes of this demo

    Contains the Datastore interface through which the CRUD operations are performed

/model/payment

    Respresents the actual payment records and the implementation of the Datastore interface

/model/errors

    Custom errors for the payment implementation

/api/payments

    Contains the payments subset of routes for the chi router. This delegates the actual logic to /model/payment and is just responsible for handling the http side of things
