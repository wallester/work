# Errors package

## Overview
This package contains is meant to be used for any errors related functionality.
It contains all of our constant errors, both RPC (status errors)
and constant errors that are used across the microservices.

The structure of this package is the following:

`business_errors.go` - RPC status errors

`business_errors_test.go` - Tests for RPC status errors

`error_codes.go` - Internal status codes that are used for RPC status errors

`custom_data.go` - Custom error type that allows to add additional metadata to error which will be added to the logs & rollbar.

`custom_data_error.go` - Implementation of methods used to work with custom data.

`custom_data_test.go` - Tests for custom data related functionality.

`errors.go` - Constant errors that are used in our microservices.

`fields.go` - Constant key names used in logging and custom data.

## Custom data

### Overview

Custom data is a new type, which is needed to be able to send additional data to our logger and rollbar. It works by adding values from key/value map to the log entry. It must be used to enrich messages returned in log messages and rollbar, e.g adding additional account-id field so that it can be easily filtered/found in the error without the need to add it inside the error itself because that does not allow to mute errors in rollbar as every error will be unique due to UUIDs or any other data being passed in the error text.

### Usage example

Custom data is very easy to use and has similar interface to the existing juju/errors package. There are several key methods that must be used in order to work the with custom data.

TO BE COMPLETED
