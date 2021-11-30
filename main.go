package main

import (
	"os"
)

var DB = os.Getenv("DB")
var CLIENT_CERT = os.Getenv("CLIENT_CERT")
var CLIENT_KEY = os.Getenv("CLIENT_KEY")
var SERVER_CA = os.Getenv("SERVER_CA")
var SERVER_NAME = os.Getenv("SERVER_NAME")
var PORT = os.Getenv("PORT")
var FAVICON = os.Getenv("FAVICON")

func main() { api(initDB(DB, CLIENT_CERT, CLIENT_KEY, SERVER_CA, SERVER_NAME), PORT, FAVICON) }