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

