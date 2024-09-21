# go-gocardless
Go Library for [Gocardless Bank Account Data API](https://developer.gocardless.com/bank-account-data/overview)

## Requirements
- Go 1.16+
- Gocardless API Secret ID & Key
- Taskfile (optional)

## Usage
Usage examples can be found in the respective package's `*_test.go` files.

# Implemented Endpoints
- Token 
- Accounts - Balances, Details, Transactions
- Agreements
- Institutions
- Requisitions

Please refer to [Gocardless Bank Account Data API](https://developer.gocardless.com/bank-account-data/overview) for more information on the endpoints.

# Pending Endpoints
- Payments - Since this needs some access to the API which is paid, I will not be implementing this since I can't test. If you would like to contribute, please feel free to open a PR.

# Issues
Please report any issues or bugs to the [Issues](https://github.com/cksidharthan/go-gocardless/issues) page.

![pkg-coverage-img](./assets/cover-treemap.svg?raw=true "Unit Test Coverage Image")