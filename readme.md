# BankServices

- [Soon](## Soon)
- [Services](## Services)
- [Account](## Account)
- [Transactions](## Transactions)
- [Installations] (## Installations)
- [Dependencies](## Dependencies)
- [Lauching](## Lauching)

**Banking services developed in Go based on a microservices architecture, including a microservice for account management, another for certain banking transactions (deposits and withdrawals).**

## Soon

A microservice for bank transfers.
## Services

## Account


## Transactions


## Installations

`install go v1.20` **[Link]**(https://go.dev/doc/install) to download.

After installing, make sure you have the latest version by typing `go version` in your terminal. Then add the gopath directory in your environment variables. On Windows your gopath is `%HOMEPATH%/go`
Check these tutorials for more information [Windows](https://www.youtube.com/watch?v=kjr3mOPv8Sk) and [Linux](https://fr.techtribune.net/linux/comment-installer-golang-langage-de-programmation-go-sous-linux/484285/)

In your `$GOPATH` directory, in `$GOPATH/src` clone the repository.

Install MongoDB and MongoDB compass or install mongodb and create a cluster MongoDB on mongodb Atlas. By default, i installed MongoDB and MongoDB compass to manage my data storage.

### Dependencies

You have to install the following dependencies. In repositories type `go get .`

or install :

* gin from github by typing `go get github.com/gin-gonic/gin v1.9.0`
* mongodb dirvers by typing `go get go.mongodb.org/mongo-driver/mongo`  and  `go get go.mongodb.org/mongo-driver/bson/primitive`
* gotoenv to load our project env variables `go get github.com/joho/godotenv`
* validator by typing `go get github.com/go-playground/validator`

### Launching

After those installation, you can run `go run main.go` in your project directory