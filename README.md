# Go API client for openapi

Mastercard Send Payment Transfer API

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.0.1
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://developer.mastercard.com/support](https://developer.mastercard.com/support)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://sandbox.api.mastercard.com/send*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*PaymentTransferApi* | [**CreatePayment**](docs/PaymentTransferApi.md#createpayment) | **Post** /v1/partners/{partnerId}/transfers/payment | Create a payment transfer (Payment Transaction) to send funds to a recipient
*PaymentTransferApi* | [**GetTransferById**](docs/PaymentTransferApi.md#gettransferbyid) | **Get** /v1/partners/{partnerId}/transfers/{transferId} | Retrieve details of a payment transfer, including latest status, by using the transfer ID that was returned when you created the transfer
*PaymentTransferApi* | [**GetTransferByRef**](docs/PaymentTransferApi.md#gettransferbyref) | **Get** /v1/partners/{partnerId}/transfers | Retrieve details of a payment transfer, including latest status, by using the transfer_reference value you provided when you created the transfer


## Documentation For Models

 - [Authentication](docs/Authentication.md)
 - [CustomField](docs/CustomField.md)
 - [Data](docs/Data.md)
 - [DataGet](docs/DataGet.md)
 - [FundingTransactionReference](docs/FundingTransactionReference.md)
 - [GovernmentIds](docs/GovernmentIds.md)
 - [GovernmentIdsResponse](docs/GovernmentIdsResponse.md)
 - [P2pCustomField](docs/P2pCustomField.md)
 - [P2pData](docs/P2pData.md)
 - [P2pRecipient](docs/P2pRecipient.md)
 - [P2pReconciliationData](docs/P2pReconciliationData.md)
 - [P2pResponseRecipientAddress](docs/P2pResponseRecipientAddress.md)
 - [P2pResponseSenderAddress](docs/P2pResponseSenderAddress.md)
 - [P2pSender](docs/P2pSender.md)
 - [PaymentTransfer](docs/PaymentTransfer.md)
 - [PaymentTransferWrapper](docs/PaymentTransferWrapper.md)
 - [Recipient](docs/Recipient.md)
 - [ReconciliationData](docs/ReconciliationData.md)
 - [RequestParticipant](docs/RequestParticipant.md)
 - [RequestRecipientAddress](docs/RequestRecipientAddress.md)
 - [RequestSenderAddress](docs/RequestSenderAddress.md)
 - [RequestTransferAcceptorAddress](docs/RequestTransferAcceptorAddress.md)
 - [ResponseParticipant](docs/ResponseParticipant.md)
 - [ResponseTransferAcceptorAddress](docs/ResponseTransferAcceptorAddress.md)
 - [Sender](docs/Sender.md)
 - [TokenCryptogram](docs/TokenCryptogram.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionAmount](docs/TransactionAmount.md)
 - [TransactionGet](docs/TransactionGet.md)
 - [TransactionHistory](docs/TransactionHistory.md)
 - [TransactionHistoryGet](docs/TransactionHistoryGet.md)
 - [Transfer](docs/Transfer.md)
 - [TransferAmount](docs/TransferAmount.md)
 - [TransferGet](docs/TransferGet.md)
 - [TransferWrapper](docs/TransferWrapper.md)
 - [TransferWrapperGet](docs/TransferWrapperGet.md)
 - [Transfers](docs/Transfers.md)
 - [TransfersWrapper](docs/TransfersWrapper.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

apisupport@mastercard.com

