# Step & Command

## initialize go project assessment

```sh
go mod init github.com/BBBunnyDefi/assessment
go mod tidy
```

## create project structure

> ref: devops-go-example structure

```sh
$ tree
.
├── Dockerfile
├── Dockerfile.test
├── README.md
├── db
│   └── 01-init.sql
├── docker-composer.test.yml
├── expenses.postman_collection.json
├── go.mod
├── note.html
├── note.md
├── rest
│   └── expenses
│       ├── db.go
│       ├── expenses.go
│       ├── expenses_integration_test.go
│       └── expenses_test.go
├── server.go
└── three-way-merge.png
```

## implement graceful shutdown with middleware logger

```sh
$ PORT:2565 go run server.go
```

## create Makefile

```sh
$ make server
```

## update Dockerfile & docker-composer & db

- docker-composer.test.yml
- Dockerfile
- Dockerfile.test
- db/01-init.sql

## test postman-collection

> should show all error

```sh
# run server
$ make server
```

```sh
$ newman run expenses.postman_collection.json
newman

expenses

→ create expense
  POST http://localhost:2565/expenses [404 Not Found, 154B, 36ms]
  ┌
  │ { message: 'Not Found' }
  └
  1. should response success(201) and object of customer
  2. Status code is 201 or 202

→ get latest expense (expenses/:id)
  GET http://localhost:2565/expenses/null [404 Not Found, 154B, 6ms]
  3. Status code is 200
  4. should response object of latest expense

→ update latest expenses
  PUT http://localhost:2565/expenses/null [404 Not Found, 154B, 4ms]
  5. Status code is 200
  6. should response success(200) and object of customer

→ get all expenses
  GET http://localhost:2565/expenses [404 Not Found, 154B, 4ms]
  7. Status code is 200
  8. should response success(200) and object of latest expense

→ Bonus middleware check Autorization
  GET http://localhost:2565/expenses [404 Not Found, 154B, 5ms]
  9. Status code is 401 Unauthorized

┌─────────────────────────┬──────────────────┬──────────────────┐
│                         │         executed │           failed │
├─────────────────────────┼──────────────────┼──────────────────┤
│              iterations │                1 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│                requests │                5 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│            test-scripts │                5 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│      prerequest-scripts │                0 │                0 │
├─────────────────────────┼──────────────────┼──────────────────┤
│              assertions │                9 │                9 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 190ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 120B (approx)                            │
├───────────────────────────────────────────────────────────────┤
│ average response time: 11ms [min: 4ms, max: 36ms, s.d.: 12ms] │
└───────────────────────────────────────────────────────────────┘

  #  failure                    detail                                                                                                   
                                                                                                                                         
 1.  AssertionError             should response success(201) and object of customer                                                      
                                expected undefined to deeply equal 'strawberry smoothie'                                                 
                                at assertion:0 in test-script                                                                            
                                inside "create expense"                                                                                  
                                                                                                                                         
 2.  AssertionError             Status code is 201 or 202                                                                                
                                expected 404 to be one of [ 201, 202 ]                                                                   
                                at assertion:1 in test-script                                                                            
                                inside "create expense"                                                                                  
                                                                                                                                         
 3.  AssertionError             Status code is 200                                                                                       
                                expected response to have status code 200 but got 404                                                    
                                at assertion:0 in test-script                                                                            
                                inside "get latest expense (expenses/:id)"                                                               
                                                                                                                                         
 4.  TypeError                  should response object of latest expense                                                                 
                                Cannot read properties of undefined (reading 'toString')                                                 
                                at assertion:1 in test-script                                                                            
                                inside "get latest expense (expenses/:id)"                                                               
                                                                                                                                         
 5.  AssertionError             Status code is 200                                                                                       
                                expected response to have status code 200 but got 404                                                    
                                at assertion:0 in test-script                                                                            
                                inside "update latest expenses"                                                                          
                                                                                                                                         
 6.  TypeError                  should response success(200) and object of customer                                                      
                                Cannot read properties of undefined (reading 'toString')                                                 
                                at assertion:1 in test-script                                                                            
                                inside "update latest expenses"                                                                          
                                                                                                                                         
 7.  AssertionError             Status code is 200                                                                                       
                                expected response to have status code 200 but got 404                                                    
                                at assertion:0 in test-script                                                                            
                                inside "get all expenses"                                                                                
                                                                                                                                         
 8.  AssertionError             should response success(200) and object of latest expense                                                
                                expenses should not be empty: expected undefined to be a number or a date                                
                                at assertion:1 in test-script                                                                            
                                inside "get all expenses"                                                                                
                                                                                                                                         
 9.  AssertionError             Status code is 401 Unauthorized                                                                          
                                expected response to have status code 401 but got 404                                                    
                                at assertion:0 in test-script                                                                            
                                inside "Bonus middleware check Autorization"          
```

## create branch EXP01

> `expenses.go`, `expenses_test.go`
>
> run test => pass

```sh
$ go test -v ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestHealthHandler
--- PASS: TestHealthHandler (0.00s)
=== RUN   TestCreateExpensesHandler
    expenses_test.go:40: EXP01: POST /expenses - with json body  COMPLETED!!
--- PASS: TestCreateExpensesHandler (0.00s)
PASS
ok      github.com/BBBunnyDefi/assessment/rest/expenses (cached)
```

> switch to main branch and use 3 ways merge --no-ff

```sh
$ git switch main
$ git merge --no-ff EXP01
# :wq
Merge made by the 'ort' strategy.
 go.mod                         |  10 +++++++++-
 go.sum                         |  11 +++++++++++
 note.md                        |  18 +++++++++++++++++
 rest/expenses/expenses.go      |  28 ++++++++++++++++++++++++++
 rest/expenses/expenses_test.go | 120 +++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 server.go                      |   1 +
 6 files changed, 187 insertions(+), 1 deletion(-)
```

## create sample integration test

> `expenses_integration_test.go`
>
> no connection to database
>
> just test TestITHealthHandler

## first run docker compose test

> command

```sh
# run
$ docker-compose -f docker-composer.test.yml up --build --abort-on-container-exit --exit-code-from it_tests

# down
$ docker-compose -f docker-composer.test.yml down
```

> try => some container using it.

```sh
$ docker-compose -f docker-composer.test.yml up --build --abort-on-container-exit --exit-code-from it_tests
[+] Building 0.1s (6/6) FINISHED                                                                                              
 => [internal] load build definition from Dockerfile.test                                                                0.0s
 => => transferring dockerfile: 179B                                                                                     0.0s
 => [internal] load .dockerignore                                                                                        0.0s
 => => transferring context: 2B                                                                                          0.0s
 => [internal] load metadata for docker.io/library/golang:1.19-alpine                                                    0.0s
 => [1/2] FROM docker.io/library/golang:1.19-alpine                                                                      0.0s
 => CACHED [2/2] WORKDIR /go/src/target                                                                                  0.0s
 => exporting to image                                                                                                   0.0s
 => => exporting layers                                                                                                  0.0s
 => => writing image sha256:34a0f43c68f82c949b8e1de7ae13c38891bf0cf44f421c15f761a16deba98ed7                             0.0s
 => => naming to docker.io/library/assessment-it_tests                                                                   0.0s
[+] Running 3/2
 ⠿ Network assessment_integration-test-example  Created                                                                  0.1s
 ⠿ Container assessment-db-1                    Created                                                                  0.1s
 ⠿ Container assessment-it_tests-1              Created                                                                  0.0s
Attaching to assessment-db-1, assessment-it_tests-1
Error response from daemon: driver failed programming external connectivity on endpoint assessment-db-1 (06893cc1deac64f24af519f03dd479e6c74d871d764c224e0efbfcc2be0bbaae): Bind for 0.0.0.0:5432 failed: port is already allocated
```

> find and stop it

```sh
$ docker ps --all | grep 5432

$ docker stop <container-name>
```

> run docker-compose up again!!
>
> Stop: run down before run up

```sh
$ docker-compose -f docker-composer.test.yml down
$ docker-compose -f docker-composer.test.yml up --build --abort-on-container-exit --exit-code-from it_tests
# ...
assessment-it_tests-1  | go: downloading golang.org/x/sys v0.3.0
assessment-it_tests-1  | go: downloading golang.org/x/text v0.5.0
assessment-it_tests-1  | ?      github.com/BBBunnyDefi/assessment       [no test files]
assessment-it_tests-1  | ok     github.com/BBBunnyDefi/assessment/rest/expenses 0.011s
assessment-it_tests-1 exited with code 0
Aborting on container exit...
[+] Running 2/2
 ⠿ Container assessment-it_tests-1  Stopped                                                                                                 0.0s
 ⠿ Container assessment-db-1        Stopped       
```

> let's test func CreateExpensesHandler to database
>
> Step 1: start postgres:12.12 database
>
> Step 2: start server
>
> Step 3: use Thunder Client POST with Body (EXP01: POST /expenses - with json body)

```sh
$ docker start assessment-db-1
$ make server
# ... test it
```

## next to Story: EXP02 for getting data

> create branch EXP02

> update `server.go`, `expenses.go`, `expenses_test.go`
>
> try test

```sh
$ go test -v --tags=unit ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestHealthHandler
--- PASS: TestHealthHandler (0.00s)
=== RUN   TestCreateExpensesHandler
    expenses_test.go:43: EXP01: POST /expenses - with json body  COMPLETED!!
--- PASS: TestCreateExpensesHandler (0.00s)
=== RUN   TestGetExpensesHandler
    expenses_test.go:128: EXP02: GET /expenses/:id COMPLETED!!
--- PASS: TestGetExpensesHandler (0.00s)
PASS
ok      github.com/BBBunnyDefi/assessment/rest/expenses (cached)
```

> but try to fail

```sh
$ go test -v --tags=unit ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestHealthHandler
--- PASS: TestHealthHandler (0.00s)
=== RUN   TestCreateExpensesHandler
    expenses_test.go:43: EXP01: POST /expenses - with json body  COMPLETED!!
--- PASS: TestCreateExpensesHandler (0.00s)
=== RUN   TestGetExpensesHandler
    expenses_test.go:128: EXP02: GET /expenses/:id COMPLETED!!
    expenses_test.go:233: 
                Error Trace:    /Users/bunny/Learns/kbtg/production-assessment/assessment/rest/expenses/expenses_test.go:233
                Error:          Not equal: 
                                expected: "{\"id\":1,\"title\":\"apple smoothie1\",\"amount\":89,\"note\":\"no discount\",\"tags\":[\"beverage\"]}"
                                actual  : "{\"id\":1,\"title\":\"apple smoothie\",\"amount\":89,\"note\":\"no discount\",\"tags\":[\"beverage\"]}"
                            
                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -{"id":1,"title":"apple smoothie1","amount":89,"note":"no discount","tags":["beverage"]}
                                +{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}
                Test:           TestGetExpensesHandler
--- FAIL: TestGetExpensesHandler (0.00s)
FAIL
FAIL    github.com/BBBunnyDefi/assessment/rest/expenses 0.925s
FAIL
```

> ok, rollback

```sh
$ go test -v --tags=unit ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestHealthHandler
--- PASS: TestHealthHandler (0.00s)
=== RUN   TestCreateExpensesHandler
    expenses_test.go:43: EXP01: POST /expenses - with json body  COMPLETED!!
--- PASS: TestCreateExpensesHandler (0.00s)
=== RUN   TestGetExpensesHandler
    expenses_test.go:128: EXP02: GET /expenses/:id COMPLETED!!
--- PASS: TestGetExpensesHandler (0.00s)
PASS
ok      github.com/BBBunnyDefi/assessment/rest/expenses (cached)
```

> test with Thunder Client [EXP01 - EXP02]