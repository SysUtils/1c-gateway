1c have beautiful odata interface& This instrument have 3 options:
1. Work as library for 1C odata interface. Project generate methods and types by metadata with translating cyrillic with dictionary(types.dat, fields.dat)
2. Work as grpc middleware

### Install

git clone <repo>

go install

### Use as library

1. Go to dir with project

2. Fill types.dat and fileds.dat

3. Run generate lib

```bash
1cclientgenerator <1c-server> <1c-odata-instance> <odata-usesrname> <odata-password> ./types.dat ./fields.dat
```

For example
```bash
1cclientgenerator.git 127.0.0.1:80 JewelleryTrade odata 11111 ./types.dat ./fields.dat
```

### Use as middleware with graphql interface


### Use as middleware with grpc interface
(In feature versions)