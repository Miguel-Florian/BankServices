# BankServices

Banking services developed in Go based on a microservices architecture and Mongodb as Database server, including a microservice for account management, another for certain banking transactions (deposits and withdrawals).

| Rapids links | Description |
| -------- | ----------- |
|  [Soon](#Soon) | Next features |
|  [Services](#Services)| The existing services |
|  [Account](#Account) | Account Management microservice |
|  [Transactions](#Transactions) | Transaction Microservice |
|  [Installations](#Installations) | Steps and requirement for installation |
|  [Dependencies](#Dependencies) | List and version of all dependencies |
|  [Lauching](#Lauching) | Run the app |

## Soon

A microservice for bank transfers and tests of each fonctionnalities of each microservices

## Installations

Install `go v1.20.3` [**Link to download**](https://go.dev/doc/install).

After installing, make sure you have the appropriate version by typing `go version` in your terminal then, configure your workspace path. If you're new in go. See these tutorials for more information [**Windows**](https://www.youtube.com/watch?v=kjr3mOPv8Sk) and [**Linux**](https://fr.techtribune.net/linux/comment-installer-golang-langage-de-programmation-go-sous-linux/484285/)

Install [**MongoDB**](https://www.mongodb.com/try/download/community) and [**MongoDB compass**](https://www.mongodb.com/try/download/compass) or install mongodb and create a cluster MongoDB on mongodb Atlas. By default, i installed MongoDB and MongoDB compass to manage my data storage.

In your `$GOPATH` directory, opens `$GOPATH/src` clone this repository.

### Dependencies

You have to install the following dependencies

* gin from github by typing `go get github.com/gin-gonic/gin v1.9.0`
* mongodb dirvers by typing `go get go.mongodb.org/mongo-driver/mongo`  and  `go get go.mongodb.org/mongo-driver/bson/primitive`
* gotoenv to load our project env variables `go get github.com/joho/godotenv`
* validator by typing `go get github.com/go-playground/validator`

## Services

Below, we have the existing services

## Account

As account model, i used
```Go
type Account struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,onitempty"`
	Number_account string             `json:"account_number,onitempty" bson:"account_number,onitempty,unique"`
	Type_account   string             `json:"account_type,onitempty" bson:"account_type,onitempty"`
	Amount         int64              `json:"amount,onitempty" bson:"amount,onitempty"`
	DateCreation   time.Time          `json:"dateCreation,onitempty" bson:"dateCreation,onitempty"`
	DateUpdate     time.Time          `json:"dateUpdate,onitempty" bson:"dateUpdate,onitempty"`
	FirstName      string             `json:"firstname,onitempty" bson:"firstname,onitempty"`
	LastName       string             `json:"lastname,onitempty,unique" bson:"lastname,onitempty"`
	Email          string             `json:"email,onitempty" bson:"email,onitempty,unique"`
	Adresse        string             `json:"adresse,onitempty" bson:"adresse,onitempty,unique"`
	Phone          string             `json:"phone,omnitempty" bson:"phone,onitempty,unique"`
	Password       string             `json:"-" bson:"password,onitempty"`
}
```
This microservice is executed on `PORT = 3000` and has the routes, opens ` BankServices/AccountServices/routers/accountRouter.go`

## Transactions

As transaction model, i used
```Go
type Transaction struct {
	ID             primitive.ObjectID `json:"-" bson:"_id,onitempty"`
	Number_account string             `json:"account_number,onitempty" bson:"account_number,onitempty,unique"`
	Amount         int64              `json:"amountDeposit,onitempty" bson:"amountDeposit,onitempty"`
	Date           time.Time          `json:"date,onitempty" bson:"date,onitempty"`
}
```
This transcation microservice is executed on `PORT = 8080` and has the routes, opens `BankServices/TransactionServices/routers/Transactionrouter.go`

### Launching

After all of this, you can run this command in terminal of each belowing paths inside the `BankServices/AccountServices ` `BankServices/TransactionServices`

```Go 
  go run main.go 
``` 
