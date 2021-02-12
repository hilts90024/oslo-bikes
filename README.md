# Oslo Bikes Client
https://oslobysykkel.no/apne-data/sanntid

## Install go
https://golang.org/doc/install

## Run in terminal
Required environment variables
- `BASE_URL` e.g. `https://gbfs.urbansharing.com/oslobysykkel.no/`
- `CLIENT_IDENTIFIER` e.g. `mittfirma-reiseplanlegger`
```sh
BASE_URL=https://gbfs.urbansharing.com/oslobysykkel.no/ CLIENT_IDENTIFIER=mittfirma-reiseplanlegger go run cmd/main.go
```


## GET
Visit http://localhost:8080 to see result in `json`.


## Run tests
```sh
go test -race -v ./...
```