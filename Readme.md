1c have beautiful odata interface& This instrument have 3 options:
1. Work as library for 1C odata interface. Project generate methods and types by metadata with translating cyrillic with dictionary(types.dat, fields.dat)
2. Work as grpc middleware
3. Work as graphql middleware

For using you must install protobuf and protoc-gen-go(https://github.com/golang/protobuf):
```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

### 1c-native

#### Install

```bash
git clone <repo>
cd cmd/1c-native/
go generate
go install
```

#### Use

1. Go to dir with project

2. Fill types.dat and fileds.dat

3. Run generate lib

```bash
1c-gateway <1c-server> <1c-odata-instance> <odata-usesrname> <odata-password>
```

For example
```bash
1c-gateway 127.0.0.1:80 JewelleryTrade odata 11111
```

### 1c-graphql-gateway
Use as middleware with graphql interface

#### Install

```bash
git clone <repo>
cd cmd/1c-graphql-gateway/
go generate
go install
```



### Use as middleware with grpc interface
(In feature versions)