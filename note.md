# Step & Commands & Journey

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

## commit & merge EXP02

> switch to main branch and use 3 ways merge --no-ff

```sh
$ git switch main
$ git merge --no-ff EXP02
# :wq
Merge made by the 'ort' strategy.
 note.md                        |  69 ++++++++++++++++++++++++++++++++++++
 rest/expenses/expenses.go      |  49 +++++++++++++++++++++-----
 rest/expenses/expenses_test.go | 117 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++--
 server.go                      |   2 ++
 4 files changed, 225 insertions(+), 12 deletions(-)
```

![01 Git Graph](img/01-git-graph.png)

## test postman collection again

> run server & run newman

```sh
$ make server
```

```sh
$ newman run expenses.postman_collection.json
newman

expenses

→ create expense
  POST http://localhost:2565/expenses [201 Created, 256B, 56ms]
  ┌
  │ {
  │   id: 4,
  │   title: 'strawberry smoothie',
  │   amount: 79,
  │   note: 'night market promotion discount 10 bath',
  │   tags: [ 'food', 'beverage' ]
  │ }
  └
  ✓  should response success(201) and object of customer
  ✓  Status code is 201 or 202

→ get latest expense (expenses/:id)
  GET http://localhost:2565/expenses/4 [200 OK, 251B, 10ms]
  ✓  Status code is 200
  ✓  should response object of latest expense

→ update latest expenses
  PUT http://localhost:2565/expenses/4 [405 Method Not Allowed, 193B, 5ms]
  1. Status code is 200
  2. should response success(200) and object of customer

→ get all expenses
  GET http://localhost:2565/expenses [405 Method Not Allowed, 194B, 4ms]
  3. Status code is 200
  4. should response success(200) and object of latest expense

→ Bonus middleware check Autorization
  GET http://localhost:2565/expenses [405 Method Not Allowed, 194B, 3ms]
  5. Status code is 401 Unauthorized

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
│              assertions │                9 │                5 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 210ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 353B (approx)                            │
├───────────────────────────────────────────────────────────────┤
│ average response time: 15ms [min: 3ms, max: 56ms, s.d.: 20ms] │
└───────────────────────────────────────────────────────────────┘

  #  failure           detail                                                              
                                                                                           
 1.  AssertionError    Status code is 200                                                  
                       expected response to have status code 200 but got 405               
                       at assertion:0 in test-script                                       
                       inside "update latest expenses"                                     
                                                                                           
 2.  TypeError         should response success(200) and object of customer                 
                       Cannot read properties of undefined (reading 'toString')            
                       at assertion:1 in test-script                                       
                       inside "update latest expenses"                                     
                                                                                           
 3.  AssertionError    Status code is 200                                                  
                       expected response to have status code 200 but got 405               
                       at assertion:0 in test-script                                       
                       inside "get all expenses"                                           
                                                                                           
 4.  AssertionError    should response success(200) and object of latest expense           
                       expenses should not be empty: expected undefined to be a number or  
                       a date                                                              
                       at assertion:1 in test-script                                       
                       inside "get all expenses"                                           
                                                                                           
 5.  AssertionError    Status code is 401 Unauthorized                                     
                       expected response to have status code 401 but got 405               
                       at assertion:0 in test-script                                       
                       inside "Bonus middleware check Autorization"   
```

> ok, not bad

## 3-way merge ?

```sh
* 72c38f6 (HEAD -> main, origin/main, origin/HEAD) update note.md
* 9690434 update note.md
*   1c9ee3f Merge branch 'EXP02'
|\  
| * 84c955c (EXP02) update note.md
| * 63e1be3 add router GetExpensesHandler
| * 18abf27 create func TestGetExpensesHandler
| * 49d8252 create func GetExpensesHandler
|/  
* 7565c00 update progress to note.md
* 1b56b19 delete note.html
* 89f071b create sample integration test but no connection to database
* a069abd update db connection
* 1009720 update note.md
*   e5caf1c Merge branch 'EXP01'
|\  
| * ada111c (EXP01) update note.md
| * d8c505d create func CreateExpensesHandler & TestCreateExpensesHandler
| * 023d119 create test TestHealthHandler
|/  
* 88bacf1 update Makefile for run env PORT & DATABASE_URL
* dd553d9 update important files for build system
* dc3ba56 create Makefile
* b7e195e implement graceful shutdown
* 38c09ac create simple server
* 4cf9280 create expenses struct & handler
* 44a1d61 create project structure
* e873bfb just empty note.md
```

## next to Story: EXP03 for update data

> create branch EXP03

> update `server.go`, `expenses.go`, `expenses_test.go`
>
> unittest failures, but functionl update working with database
>
> TODO: Fix later, now use t.Skip()

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
=== RUN   TestUpdateExpensesHandler
    expenses_test.go:384: 
                Error Trace:    /Users/bunny/Learns/KBTG/production-assessment/assessment/rest/expenses/expenses_test.go:384
                Error:          Not equal: 
                                expected: 200
                                actual  : 400
                Test:           TestUpdateExpensesHandler
    expenses_test.go:385: 
                Error Trace:    /Users/bunny/Learns/KBTG/production-assessment/assessment/rest/expenses/expenses_test.go:385
                Error:          Not equal: 
                                expected: "{\"id\":1,\"title\":\"apple smoothie\",\"amount\":89,\"note\":\"no discount\",\"tags\":[\"beverage\"]}"
                                actual  : "{\"message\":\"can't execute expenses: call to ExecQuery 'UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1' with args [{Name: Ordinal:1 Value:1} {Name: Ordinal:2 Value:apple smoothie} {Name: Ordinal:3 Value:89} {Name: Ordinal:4 Value:no discount} {Name: Ordinal:5 Value:{\\\"beverage\\\"}}], was not expected, next expectation is: ExpectedQuery =\\u003e expecting Query, QueryContext or QueryRow which:\\n  - matches sql: 'UPDATE expenses'\\n  - is with arguments:\\n    0 - 1\\n    1 - apple smoothie\\n    2 - 89\\n    3 - no discount\\n    4 - \\u0026[beverage]\\n  - should return rows:\\n    row 0 - [1 apple smoothie 89 no discount {\\\"beverage\\\"}]\"}"
                            
                                Diff:
                                --- Expected
                                +++ Actual
                                @@ -1 +1 @@
                                -{"id":1,"title":"apple smoothie","amount":89,"note":"no discount","tags":["beverage"]}
                                +{"message":"can't execute expenses: call to ExecQuery 'UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id=$1' with args [{Name: Ordinal:1 Value:1} {Name: Ordinal:2 Value:apple smoothie} {Name: Ordinal:3 Value:89} {Name: Ordinal:4 Value:no discount} {Name: Ordinal:5 Value:{\"beverage\"}}], was not expected, next expectation is: ExpectedQuery =\u003e expecting Query, QueryContext or QueryRow which:\n  - matches sql: 'UPDATE expenses'\n  - is with arguments:\n    0 - 1\n    1 - apple smoothie\n    2 - 89\n    3 - no discount\n    4 - \u0026[beverage]\n  - should return rows:\n    row 0 - [1 apple smoothie 89 no discount {\"beverage\"}]"}
                Test:           TestUpdateExpensesHandler
--- FAIL: TestUpdateExpensesHandler (0.00s)
FAIL
FAIL    github.com/BBBunnyDefi/assessment/rest/expenses 1.214s
FAIL
```

## commit & merge EXP03

> switch to main branch and use 3 ways merge --no-ff

```sh
$ git switch main
Switched to branch 'main'
Your branch is up to date with 'origin/main'.

$ git merge --no-ff EXP03
Merge made by the 'ort' strategy.
 note.md                        |  48 +++++++++++++++++++++++++++++++++++++++-
 rest/expenses/expenses.go      |  30 +++++++++++++++++++++++++
 rest/expenses/expenses_test.go | 152 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 server.go                      |   2 ++
 4 files changed, 231 insertions(+), 1 deletion(-)
```

> check status & push to repo

```sh
$ git status
On branch main
Your branch is ahead of 'origin/main' by 7 commits.
  (use "git push" to publish your local commits)

nothing to commit, working tree clean

$ git push
Enumerating objects: 33, done.
Counting objects: 100% (33/33), done.
Delta compression using up to 4 threads
Compressing objects: 100% (23/23), done.
Writing objects: 100% (26/26), 3.98 KiB | 1.99 MiB/s, done.
Total 26 (delta 17), reused 0 (delta 0), pack-reused 0
remote: Resolving deltas: 100% (17/17), completed with 6 local objects.
To https://github.com/BBBunnyDefi/assessment.git
   e0fdd4e..d38c9c6  main -> main
```

## next to Story: EXP04 for get all data

> create branch EXP04

```sh
$ git checkout -b EXP04
Switched to a new branch 'EXP04'

$ git branch
  EXP01
  EXP02
  EXP03
* EXP04
  main
(END)
# q
```

> update  `server.go`, `expenses.go`, `expenses_test.go`
> 
> setup commend `Makefile`
>
> test all: before merge to main branch 
> - story by Thunder Client
> - unittest(EXP03-Skip)
> - integration tests(no connect database)
> - postman collection tests

> unittest

```sh
make unittest
go test -v --tags=unit ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestHealthHandler
--- PASS: TestHealthHandler (0.00s)
=== RUN   TestCreateExpensesHandler
    expenses_test.go:43: EXP01: POST /expenses - with json body  COMPLETED!!
--- PASS: TestCreateExpensesHandler (0.00s)
=== RUN   TestGetExpensesHandler
    expenses_test.go:128: EXP02: GET /expenses/:id COMPLETED!!
--- PASS: TestGetExpensesHandler (0.00s)
=== RUN   TestUpdateExpensesHandler
    expenses_test.go:238: TODO: EXP03: PUT /expenses/:id - with json body FAILED!!
--- SKIP: TestUpdateExpensesHandler (0.00s)
=== RUN   TestGetAllExpensesHandler
    expenses_test.go:391: EXP04: GET /expenses COMPLETED!!
--- PASS: TestGetAllExpensesHandler (0.00s)
PASS
ok      github.com/BBBunnyDefi/assessment/rest/expenses (cached)
```

> integration test

```sh
$ make ittest
go test -v --tags=integration ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestITHealthHandler
    expenses_integration_test.go:30: TODO: implement integration Health GET /

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:80
--- PASS: TestITHealthHandler (0.00s)
PASS
ok      github.com/BBBunnyDefi/assessment/rest/expenses (cached)
```

> run postman collection

```sh
$ newman run expenses.postman_collection.json
newman

expenses

→ create expense
  POST http://localhost:2565/expenses [201 Created, 257B, 57ms]
  ┌
  │ {
  │   id: 10,
  │   title: 'strawberry smoothie',
  │   amount: 79,
  │   note: 'night market promotion discount 10 bath',
  │   tags: [ 'food', 'beverage' ]
  │ }
  └
  ✓  should response success(201) and object of customer
  ✓  Status code is 201 or 202

→ get latest expense (expenses/:id)
  GET http://localhost:2565/expenses/10 [200 OK, 252B, 9ms]
  ✓  Status code is 200
  ✓  should response object of latest expense

→ update latest expenses
  PUT http://localhost:2565/expenses/10 [200 OK, 211B, 9ms]
  ✓  Status code is 200
  ✓  should response success(200) and object of customer

→ get all expenses
  GET http://localhost:2565/expenses [200 OK, 1.21kB, 9ms]
  ✓  Status code is 200
  ✓  should response success(200) and object of latest expense

→ Bonus middleware check Autorization
  GET http://localhost:2565/expenses [200 OK, 1.21kB, 7ms]
  1. Status code is 401 Unauthorized

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
│              assertions │                9 │                1 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 219ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 2.51kB (approx)                          │
├───────────────────────────────────────────────────────────────┤
│ average response time: 18ms [min: 7ms, max: 57ms, s.d.: 19ms] │
└───────────────────────────────────────────────────────────────┘

  #  failure               detail                                                                               
                                                                                                                
 1.  AssertionError        Status code is 401 Unauthorized                                                      
                           expected response to have status code 401 but got 200                                
                           at assertion:0 in test-script                                                        
                           inside "Bonus middleware check Autorization"  
```

> merge EXP04 to main

```sh
$ git switch main
Switched to branch 'main'
Your branch is up to date with 'origin/main'.

$ git merge --no-ff EXP04
Merge made by the 'ort' strategy.
 Makefile                       |   8 ++++-
 note.md                        | 144 ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
 rest/expenses/expenses.go      |  31 +++++++++++++++++
 rest/expenses/expenses_test.go |  42 +++++++++++++++++++++++
 server.go                      |   9 +++++
 5 files changed, 233 insertions(+), 1 deletion(-)
```

> after merge to main 
> 
> test all again and ok

## implement integration tests

> create branch for dev integration tests (DevIT)

> first: simple test for pass

```sql
-- 01-init.sql
-- add data for test
INSERT INTO "expenses" ("id", "title", "amount", "note", "tags") VALUES (1, 'strawberry smoothie', 79, 'night market promotion discount 10 bath', '{"food","beverage"}');
```

```go
// create func TestITGetExpensesHandler for check equal
// ...
expected := "{\"id\":1,\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"
// ...
```

> run pass

```sh
# pass
# ...
assessment-db-1        | PostgreSQL init process complete; ready for start up.
assessment-db-1        | 
assessment-db-1        | 2023-01-03 16:14:10.622 UTC [1] LOG:  starting PostgreSQL 12.12 (Debian 12.12-1.pgdg110+1) on x86_64-pc-linux-gnu, compiled by gcc (Debian 10.2.1-6) 10.2.1 20210110, 64-bit
assessment-db-1        | 2023-01-03 16:14:10.622 UTC [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
assessment-db-1        | 2023-01-03 16:14:10.622 UTC [1] LOG:  listening on IPv6 address "::", port 5432
assessment-db-1        | 2023-01-03 16:14:10.629 UTC [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
assessment-db-1        | 2023-01-03 16:14:10.648 UTC [84] LOG:  database system was shut down at 2023-01-03 16:14:10 UTC
assessment-db-1        | 2023-01-03 16:14:10.654 UTC [1] LOG:  database system is ready to accept connections
assessment-it_tests-1  | go: downloading golang.org/x/net v0.4.0
assessment-it_tests-1  | go: downloading github.com/golang-jwt/jwt v3.2.2+incompatible
assessment-it_tests-1  | go: downloading github.com/valyala/fasttemplate v1.2.2
assessment-it_tests-1  | go: downloading golang.org/x/time v0.2.0
assessment-it_tests-1  | go: downloading github.com/mattn/go-colorable v0.1.13
assessment-it_tests-1  | go: downloading github.com/mattn/go-isatty v0.0.16
assessment-it_tests-1  | go: downloading github.com/valyala/bytebufferpool v1.0.0
assessment-it_tests-1  | go: downloading golang.org/x/text v0.5.0
assessment-it_tests-1  | go: downloading golang.org/x/sys v0.3.0
assessment-it_tests-1  | ?      github.com/BBBunnyDefi/assessment       [no test files]
assessment-it_tests-1  | ok     github.com/BBBunnyDefi/assessment/rest/expenses 0.020s
assessment-it_tests-1 exited with code 0
Aborting on container exit...
[+] Running 2/2
 ⠿ Container assessment-it_tests-1  Stopped                                                                                         0.0s
 ⠿ Container assessment-db-1        Stopped   
```

> test error

```go
// edit title: strawberry smoothie1
// ...
expected := "{\"id\":1,\"title\":\"strawberry smoothie1\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"
// ...
```

```sh
# fail
# ...
assessment-it_tests-1  |    ____    __
assessment-it_tests-1  |   / __/___/ /  ___
assessment-it_tests-1  |  / _// __/ _ \/ _ \
assessment-it_tests-1  | /___/\__/_//_/\___/ v4.10.0
assessment-it_tests-1  | High performance, minimalist Go web framework
assessment-it_tests-1  | https://echo.labstack.com
assessment-it_tests-1  | ____________________________________O/_______
assessment-it_tests-1  |                                     O\
assessment-it_tests-1  | ⇨ http server started on [::]:80
assessment-it_tests-1  | 2023/01/03 16:12:12 dial tcp 127.0.0.1:80: connect: connection refused
assessment-it_tests-1  | --- FAIL: TestITGetExpensesHandler (0.01s)
assessment-it_tests-1  |     expenses_integration_test.go:79: TODO: implement integration EXP02 GET /expenses/:id
assessment-it_tests-1  |     expenses_integration_test.go:143: 
assessment-it_tests-1  |                Error Trace:    /go/src/target/rest/expenses/expenses_integration_test.go:143
assessment-it_tests-1  |                Error:          Not equal: 
assessment-it_tests-1  |                                expected: "{\"id\":1,\"title\":\"strawberry smoothie1\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"
assessment-it_tests-1  |                                actual  : "{\"id\":1,\"title\":\"strawberry smoothie\",\"amount\":79,\"note\":\"night market promotion discount 10 bath\",\"tags\":[\"food\",\"beverage\"]}"
assessment-it_tests-1  |                            
assessment-it_tests-1  |                                Diff:
assessment-it_tests-1  |                                --- Expected
assessment-it_tests-1  |                                +++ Actual
assessment-it_tests-1  |                                @@ -1 +1 @@
assessment-it_tests-1  |                                -{"id":1,"title":"strawberry smoothie1","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}
assessment-it_tests-1  |                                +{"id":1,"title":"strawberry smoothie","amount":79,"note":"night market promotion discount 10 bath","tags":["food","beverage"]}
assessment-it_tests-1  |                Test:           TestITGetExpensesHandler
assessment-it_tests-1  | FAIL
assessment-it_tests-1  | FAIL   github.com/BBBunnyDefi/assessment/rest/expenses 0.025s
assessment-it_tests-1  | FAIL
assessment-it_tests-1 exited with code 1
Aborting on container exit...
[+] Running 2/2
 ⠿ Container assessment-it_tests-1  Stopped                                                                                         0.0s
 ⠿ Container assessment-db-1        Stopped                                                                                         0.3s
make: *** [cup] Error 1
```

> !!! 4 days

## Fixed unittest EXP03

> update command
> 
> from db.Prepare() & db.Exec()
>
> to db.QueryRow().Scan()

## implement integration tests v2

> create new branch (DevITv2)
>
> create func TestITGetAllExpensesHandler() this check data in database more than once

> Note: 
> 
> sandbox mode use `const serverPort = 80`
>
> terminal mode use `const serverPort = 2565` and run database and server
>

> Example: terminal mode

```sh
$ make dbstart
docker start assessment-db-1
assessment-db-1

$ make server
DATABASE_URL=postgres://root:root@localhost:5432/assessment?sslmode=disable PORT=:2565 go run server.go
2023/01/04 22:45:49 Database connection OK!!

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:2565

```

> terminal 2 run integration tests

```sh
make ittest
go test -v --tags=integration ./...
?       github.com/BBBunnyDefi/assessment       [no test files]
=== RUN   TestITHealthHandler

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
--- PASS: TestITHealthHandler (0.00s)
=== RUN   TestITCreateExpensesHandler
    expenses_integration_test.go:81: TODO: implement integration EXP01: POST /expenses - with json body
--- SKIP: TestITCreateExpensesHandler (0.00s)
=== RUN   TestITGetExpensesHandler

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
--- PASS: TestITGetExpensesHandler (0.00s)
=== RUN   TestITUpdateExpensesHandler
    expenses_integration_test.go:140: TODO: implement integration EXP03: PUT /expenses/:id - with json body
--- SKIP: TestITUpdateExpensesHandler (0.00s)
=== RUN   TestITGetAllExpensesHandler

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.10.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
--- PASS: TestITGetAllExpensesHandler (0.00s)
PASS
ok      github.com/BBBunnyDefi/assessment/rest/expenses (cached)
```

> test new man collection but error [500] ? at (first time)
>
> test 3 times pass and check data with Thunder client 

```sh
make newman
newman run expenses.postman_collection.json
newman

expenses

→ create expense
  POST http://localhost:2565/expenses [201 Created, 256B, 40ms]
  ┌
  │ {
  │   id: 4,
  │   title: 'strawberry smoothie',
  │   amount: 79,
  │   note: 'night market promotion discount 10 bath',
  │   tags: [ 'food', 'beverage' ]
  │ }
  └
  ✓  should response success(201) and object of customer
  ✓  Status code is 201 or 202

→ get latest expense (expenses/:id)
  GET http://localhost:2565/expenses/4 [200 OK, 251B, 8ms]
  ✓  Status code is 200
  ✓  should response object of latest expense

→ update latest expenses
  PUT http://localhost:2565/expenses/4 [200 OK, 210B, 9ms]
  ✓  Status code is 200
  ✓  should response success(200) and object of customer

→ get all expenses
  GET http://localhost:2565/expenses [200 OK, 514B, 8ms]
  ✓  Status code is 200
  ✓  should response success(200) and object of latest expense

→ Bonus middleware check Autorization
  GET http://localhost:2565/expenses [200 OK, 514B, 6ms]
  1. Status code is 401 Unauthorized

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
│              assertions │                9 │                1 │
├─────────────────────────┴──────────────────┴──────────────────┤
│ total run duration: 188ms                                     │
├───────────────────────────────────────────────────────────────┤
│ total data received: 1.12kB (approx)                          │
├───────────────────────────────────────────────────────────────┤
│ average response time: 14ms [min: 6ms, max: 40ms, s.d.: 12ms] │
└───────────────────────────────────────────────────────────────┘

  #  failure                     detail                                                                                                     
                                                                                                                                            
 1.  AssertionError              Status code is 401 Unauthorized                                                                            
                                 expected response to have status code 401 but got 200                                                      
                                 at assertion:0 in test-script                                                                              
                                 inside "Bonus middleware check Autorization"                                                               
make: *** [newman] Error 1
```

## 3 Days [04/01/2023 23:09]

**TODO**:
- Bonus middleware check Autorization ?
- create postman environment ?
- 2 func integration tests (TestITCreateExpensesHandler, TestITUpdateExpensesHandler)
- other...

> This is the first time to use Git seriously

```sh
*   f9da7c9 Merge branch 'DevITv2'
|\  
| * 6e26d68 create func TestITGetAllExpensesHandler
| * 1323ba1 update command run newman collection
| * 16d55a1 update note.md
| * b67a200 create TestITGetAllExpensesHandler test pass outside sandbox
|/  
*   24c6f95 Merge branch 'RefacExpenses'
|\  
| * 3e0dfb6 update note.md
| * 5a9ee15 try to manage something, but found solution to fixed error EXP03 for unittest
| * 4672807 add ping database server for check connecting
| * 7c13e83 noting update
|/  
* 63da823 clean up comment
* d2c4978 update note.md
*   37898e0 Merge branch 'DevIT'
|\  
| * 9bc7ffa create first step for integration test
| * 4abdf81 update command
* | 5d5fe73 update note.md
|/  
* 85f49b0 try compose up and test integration fail
* 1c6f1c0 update command compose up & down
* 679d971 update command to Makefile start&stop docker postgres database
* 80443ec update note.md
*   4154e85 Merge branch 'EXP04'
|\  
| * 40ce236 update note.md
| * 2be301e add router GetAllExpensesHandler
| * 137a372 create func TestGetAllExpensesHandler
| * 166f5be create func GetAllExpensesHandler
| * 972f597 update command to Makefile
|/  
* f70ce68 update note.md
* d38c9c6 update note.md
*   4f11833 Merge branch 'EXP03' - unittest error
|\  
| * dbf7a3a update unittest fail note.md
| * b5377c3 try to fix unittest but still not pass use t.Skip()
| * cf02862 add router UpdateExpensesHandler
| * f936f22 create func TestUpdateExpensesHandler (test not pass) TODO: Fix it's
| * beb0226 create func UpdateExpensesHandler
|/  
* e0fdd4e update note.md
* eb4cdfa update note.md
* 72c38f6 update note.md
* 9690434 update note.md
*   1c9ee3f Merge branch 'EXP02'
|\  
| * 84c955c update note.md
| * 63e1be3 add router GetExpensesHandler
| * 18abf27 create func TestGetExpensesHandler
| * 49d8252 create func GetExpensesHandler
|/  
* 7565c00 update progress to note.md
* 1b56b19 delete note.html
* 89f071b create sample integration test but no connection to database
* a069abd update db connection
* 1009720 update note.md
*   e5caf1c Merge branch 'EXP01'
|\  
| * ada111c update note.md
| * d8c505d create func CreateExpensesHandler & TestCreateExpensesHandler
| * 023d119 create test TestHealthHandler
|/  
* 88bacf1 update Makefile for run env PORT & DATABASE_URL
* dd553d9 update important files for build system
* dc3ba56 create Makefile
* b7e195e implement graceful shutdown
* 38c09ac create simple server
* 4cf9280 create expenses struct & handler
* 44a1d61 create project structure
* e873bfb just empty note.md
```

## implement integration tests v3

> integration pass all, create 2 new func
>
> forgot: create branch
>
> func TestITUpdateExpensesHandler, TestITCreateExpensesHandler

**TODO**:
- Bonus middleware check Autorization ?
- create postman environment ?
- ~~2 func integration tests (TestITCreateExpensesHandler, TestITUpdateExpensesHandler)~~
- other...

> end of today 05/01/2023 23:17 :sleepy:

## still not complete

> 06/01/2023
> - Bonus middleware check Autorization ?
> - create postman environment ?
> - ...
>
> :panda_face: 
>
> tomorrow night, deadline :point_right: [07/01/2023 23:59:59]