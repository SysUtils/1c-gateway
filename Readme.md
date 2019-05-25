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
