
# ziggy
RPC API service  to serve up Crypto price information.

## Prerequisites (MacOs)
1. Install go
2. Intall protobuf

```
$ brew install protobuf
$ protoc --version  # Ensure compiler version is 3+
```

3. Install the protocol compiler plugins for Go.

```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```

4. Update your PATH so that the protoc compiler can find the plugins.

```
$ export PATH="$PATH:$(go env GOPATH)/bin"

```

5. Build proto_go lib

```
bazel build //api:all
```
6. Create BUILD files

```
bazel run //:gazelle
```

7. Build 
```
bazel build //...
```

## Set up database

1. Install mysql 
2. Connect to mysql server 
```
mysql -u root -p
```
3. Create `coins` database
```
create database coins;
```
4. Switch to `coins` database

```
use coins;
```
5. Create `coins` table in `coins` database

```
source /path/to/ziggy/folder/database/create-table.sql
```

6. Check if table is created.
```
mysql> select * from coins;
```
should have schema of table `coins`.

## Send requests 
1. Install grpc curl to interact with server 

```
brew install grpcurl
```
2. Run GRPC server
```
bazel run //cmd/server:server
```

3. Send requests to grpc server.

### Method A
```
grpcurl -plaintext -format text \
  localhost:5000 crypto.Crypto.ListCoins
```
### Method C
```
 grpcurl \
  -d '{"coin_id": "2252",
       "enableTracking": true}' \
   -plaintext localhost:5000 crypto.Crypto.TrackCoin
```
## Known issues 
1. Initially more than 5 coins are added to database as CoinMarketPlace API gives 10 coins even when there is limit of 5.

## Dev

### Testing coinmarketplace API

1. API definitions are present [here](https://pro.coinmarketcap.com/api/v1#operation/getV1CryptocurrencyListingsLatest)

2. Curl to get 
```
curl -H "X-CMC_PRO_API_KEY: 0b424217-a625-470d-8d51-d1a344bac7d2" -H "Accept: application/json" -d "start=1&limit=5&convert=USD" -G https://sandbox-api.coinmarketcap.com/v1/cryptocurrency/listings/latest | jq
```

## Improvements 
1. Use two database tables, one for storing coins and other for storing history of prices of the coins.
2. Create unit tests for all go files.
3. Have e2e test to check if the API is working with mysql.
4. Use in-build  proto validators to check requests.  