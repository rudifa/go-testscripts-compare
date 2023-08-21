# Compare different test methods

## 1 Normal go test

```
.
├── pkg
│   └── goeval
│       ├── testdata
│       │   ├── combo.cue
│       │   ├── raw.json
│       │   └── stdout.json
│       ├── goeval.go
│       └── goeval_test.go
```

```
testscripts-compare % go test $1 ./...
?       github.com/rudifa/testscripts-compare   [no test files]
ok      github.com/rudifa/testscripts-compare/pkg/goeval        (cached)
```

Adavantages:

- a well known test method
- close to the code under test
- assisted by the vscode go extension (run test, debug test, ...)

## 2 `testscript` with a standalone txtar script

```
.
├── pkg
│   └── goeval
│       └── testdata
│           └── standalone.txtar
```

```
testscripts-compare % testscript pkg/goeval/testdata/standalone.txtar
PASS
```

Advangates:

- all-in-one script, suitable for presenting test cases in discussons and reports
