# SIMPLE REST API GO-LANG WITH MSYQL

## HOW TO INSTALL GOLAND ON UBUNTU

[INSTALL GOLANG](https://golang.org/doc/install)

## HOW TO CREATE

**FIRST CREATE FOLDER AND MOD REST API**
```
- mkdir RestApiGo
- cd RestApiGo
- go mod init RestApiGo

```

## REQUIRMENT LIBRARY

``go get -u "github.com/gin-gonic/gin"``

``go get -u "github.com/jinzhu/gorm"``

``go get -u "github.com/jinzhu/gorm/dialects/mysql"``

``go get -u "github.com/go-sql-driver/mysql"``

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


# RESULT API

### Persons

``URL : http://localhost:3000/api/v1/perons/``

```
Request:
Method: GET 
    - Endpoint: /api/v1/persons
Response:
    {
        "count": 4,
        "result": [
            {
            "ID": 1,
            "CreatedAt": "2021-02-24T14:34:30+07:00",
            "UpdatedAt": "2021-02-24T14:34:30+07:00",
            "DeletedAt": null,
            "FirstName": "LISA",
            "LastName": "HIMURA"
            },
            {
            "ID": 2,
            "CreatedAt": "2021-02-24T14:50:46+07:00",
            "UpdatedAt": "2021-02-24T14:50:46+07:00",
            "DeletedAt": null,
            "FirstName": "",
            "LastName": ""
            },
            {
            "ID": 3,
            "CreatedAt": "2021-02-24T14:51:03+07:00",
            "UpdatedAt": "2021-02-24T14:51:03+07:00",
            "DeletedAt": null,
            "FirstName": "",
            "LastName": ""
            },
            {
            "ID": 4,
            "CreatedAt": "2021-02-24T14:51:25+07:00",
            "UpdatedAt": "2021-02-24T14:51:25+07:00",
            "DeletedAt": null,
            "FirstName": "asd",
            "LastName": "asda"
            }
        ]
    }
```
# REFERENCE

**BELAJAR GOLANG**

[BELAJAR GOLANG](https://dasarpemrogramangolang.novalagung.com/1-berkenalan-dengan-golang.html)


**TUTORIAL REST API**

[GOLANG REST API](https://medium.com/skyshidigital/golang-restapi-untuk-pemula-ef1c345b3ef5)
