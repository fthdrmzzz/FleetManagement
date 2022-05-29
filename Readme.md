

This is a REST Api developed in Golang.

## Technical Details

- **Go programming language is used**
- **PostgreSQL Database is used**
- **Integration tests applied**: to test the integrity of the database, integration tests are implemented.
- **Dependency injection** principle is implemented. 
- **Dockerfile and docker-compose is implemented**
- **Logging mechanism is implemented:** Info and Errors are logged through implemented component. 
- **Feature Encapsulation is adopted:** Related files and components are grouped together.

## Folder Structure

| Folder  | Description |
| ------------- | ------------- |
| cmd  | contains the main function of the program.  |
| api  |  server related components  |
| logger | logger module|
| setup | http requests to initialize the data|
| util | utilities for test cases|

## How To Build and Run The Project

### Requirements
- [Docker](https://www.docker.com/) and [docker-compose](https://docs.docker.com/compose/) is needed to run the project.
- [Golang](https://go.dev/) is needed to run the tests.
- [Python](https://www.python.org/) is needed to test the case by sending requests. (This can be done by curl aswell. I wrote the requests in python because it was more readable)

### Running the project
If the requirements are satisfied, you should first run 
```
make run
```
This command makes the project run in port 8080.


After running the project, to send the initial data, you should run 
```
make setupdb
```
This command sends requests to fulfill the initial requirements asked in the code.

### Stopping the project
In order to stop the container, run `make stop`. 
In order to remove the container, run `make down`.

## Test the Program and Case
### Test Database
To run database integrity tests of the program, you can run `make test`
### Test The Case
After running the project, in order to test the case, you can simply run 
```
make delivery
```
This command sends the request given in the case, and prints the result in console.

Given Request
```json
{
  "plate": "34 TL 34",
  "route": [
    {
      "deliveryPoint": 1,
      "deliveries": [
        {"barcode": "P7988000121"},
        {"barcode": "P7988000122"},
        {"barcode": "P7988000123"},
        {"barcode": "P8988000121"},
        {"barcode": "C725799"}
      ]
    },
    {
      "deliveryPoint": 2,
      "deliveries": [
        {"barcode": "P8988000123"},
        {"barcode": "P8988000124"},
        {"barcode": "P8988000125"},
        {"barcode": "C725799"}
      ]
    },
    {
      "deliveryPoint": 3,
      "deliveries": [
        {"barcode": "P9988000126"},
        {"barcode": "P9988000127"},
        {"barcode": "P9988000128"},
        {"barcode": "P9988000129"},
        {"barcode": "P9988000130"}
      ]
    }
  ]
}
```

Acquired response

```json
{
    "plate": "34 TL 34",
    "route": [
        {
            "deliveries": [
                {
                    "barcode": "P7988000121",
                    "state": "unloaded"
                },
                {
                    "barcode": "P7988000122",
                    "state": "unloaded"
                },
                {
                    "barcode": "P7988000123",
                    "state": "unloaded"
                },
                {
                    "barcode": "P8988000121",
                    "state": "loaded"
                },
                {
                    "barcode": "C725799",
                    "state": "loaded"
                }
            ],
            "deliveryPoint": 1
        },
        {
            "deliveries": [
                {
                    "barcode": "P8988000123",
                    "state": "unloaded"
                },
                {
                    "barcode": "P8988000124",
                    "state": "unloaded"
                },
                {
                    "barcode": "P8988000125",
                    "state": "unloaded"
                },
                {
                    "barcode": "C725799",
                    "state": "unloaded"
                }
            ],
            "deliveryPoint": 2
        },
        {
            "deliveries": [
                {
                    "barcode": "P9988000126",
                    "state": "loaded"
                },
                {
                    "barcode": "P9988000127",
                    "state": "loaded"
                },
                {
                    "barcode": "P9988000128",
                    "state": "unloaded"
                },
                {
                    "barcode": "P9988000129",
                    "state": "unloaded"
                },
                {
                    "barcode": "P9988000130",
                    "state": "loaded"
                }
            ],
            "deliveryPoint": 3
        }
    ]
}
```

As you can see the result is sent as expected
