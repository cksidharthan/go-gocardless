# go-nordigen
Go Library for [Nordigen Openbanking API v2](https://nordigen.com/en/account_information_documenation/api-documention/overview/)

## Requirements
- Go 1.16+
- Nordigen API Secret ID & Key
- Taskfile (optional)

## Usage
Usage examples can be found in the respective package's `*_test.go` files.

# Implemented Endpoints
- Token 
- Accounts - Balances, Details, Transactions
- Agreements
- Institutions
- Requisitions

Please refer to [Nordigen API Documentation](https://nordigen.com/en/docs/account-information/integration/parameters-and-responses/) for more information on the endpoints.

# Pending Endpoints
- Payments - Since this needs some access to the API which is paid, I will not be implementing this since I can't test. If you would like to contribute, please feel free to open a PR.

# Issues
Please report any issues or bugs to the [Issues](https://github.com/weportfolio/go-nordigen/issues) page.

![pkg-coverage-img](./assets/cover-treemap.svg?raw=true "Unit Test Coverage Image")