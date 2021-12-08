package main

import (
	"os"
)

var DB = os.Getenv("DB")
var CLIENT_CERT = os.Getenv("CLIENT_CERT")
var CLIENT_KEY = os.Getenv("CLIENT_KEY")
var SERVER_CA = os.Getenv("SERVER_CA")
var SERVER_NAME = os.Getenv("SERVER_NAME")
var CACHE_TTL = os.Getenv("CACHE_TTL")
var CACHE_PRG = os.Getenv("CACHE_PRG")
var PORT = os.Getenv("PORT")
var FAVICON = os.Getenv("FAVICON")

func main() {
	database := initDB(DB, CLIENT_CERT, CLIENT_KEY, SERVER_CA, SERVER_NAME)
	cache := initCache(CACHE_TTL, CACHE_PRG)
	api(database, cache, PORT, FAVICON)
}