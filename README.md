# Micro Blog

This is the base project for the SILIN and SUBSY archetype, also to start the internal study of Go

### Generate Coverage Test with Report

* #### Test Coverage
    This is the command that will run all the project tests
    ```
    go test -coverprofile=coverage.out ./... 
    ```
  
    This is the command to see the coverage percentage, shows breakdown coverage of function
    ```
    go tool cover -func=coverage.out
    ```

    This is the command that will generate the report
    ````
    go tool cover -html=coverage.out
    ````


* #### Test Count Coverage
    We run the test and write a coverage profile so that we can present the information in a pleasant way afterwards
    ````
    go test -covermode=count -coverprofile=count.out fmt
    ````
  
    Shows breakdown coverage of function
    ````
    go tool cover -func=count.out
    ````
  
    Generate an HTML output
    ````
    go tool cover -html=count.out
    ````

* #### Cover Mode
    There are three different cover modes:
     * **set:** did each statement run?
     * **count:** how many times did each statement run?
     * **atomic:** like count, but counts precisely in parallel programs
     
### Generate Mock Interface
This is an automatic mock generator using mockery, the first thing we must do is go to the path of the file that we want to autogenerate:
```
cd path
```
After entering the route we must execute the following command, Repository this is name the interface
```
mockery -name Repository
```

### Test commands for the project
These are the commands to run the unit and integration tests of the project

#### Test Repository
```
go test -v -coverprofile=coverage.out -coverpkg=./repository ./test/repository/

go tool cover -html=coverage.out
```

#### Test Handler
```
go test -v -coverprofile=coverage.out -coverpkg=./handler/server/v1 ./test/handler/server/v1/

go tool cover -html=coverage.out
```

#### Test Integration
```
go test -v -coverprofile=coverage_integration.out  ./test/

go tool cover -html=coverage_integration.out
```

### Test coverage commands for the project

#### Test Repository
```
go test -covermode=count -coverprofile=coverage.out -coverpkg=./repository ./test/repository/

go tool cover -html=coverage.out

go tool cover -func=coverage.out
```

#### Test Handler
```
go test -covermode=count -coverprofile=coverage.out -coverpkg=./handler/server/v1 ./test/handler/server/v1/

go tool cover -html=coverage.out

go tool cover -func=coverage.out
```
