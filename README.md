# SIMPLE REST API GO-LANG WITH MSYQL

## HOW TO CREATE

**FIRST CREATE FOLDER AND MOD REST API**
```
- mkdir RestApiGo
- cd RestApiGo
- go mod init RestApiGo

```

## REQUIRMENT LIBRARY

```
go get -u "github.com/gin-gonic/gin"
go get -u "github.com/jinzhu/gorm"
go get -u "github.com/jinzhu/gorm/dialects/mysql"
go get -u "github.com/go-sql-driver/mysql"

```
## STRUCTURE FOLDER DATA


```
├── base.go
├── config
│   └── config.go
├── controllers
│   ├── gorm.go
│   └── person.go
├── go.mod
├── go.sum
└── structs
    └── structs.go

```
# REFERENCE

-[GOLANG REST API](https://medium.com/skyshidigital/golang-restapi-untuk-pemula-ef1c345b3ef5)