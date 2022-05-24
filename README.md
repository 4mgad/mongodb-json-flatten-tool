## Build

- Make sure Go is installed on your machine (if not, please refer to https://go.dev/doc/install)
- Run the following command to generate the binary file "flatten":
```shell
  go build -o flatten
```

## Unit Test
- Run the following command to run unit test cases:
```shell
  go test -v ./...
```

## Run
- Pipe any json file to the tool by running the following command example:
```shell
  cat test.json | ./flatten
```
