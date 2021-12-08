package main

import (
	"github.com/patrickmn/go-cache"
	"log"
	"strconv"
	"time"
)

func initCache(cacheTTL string, cachePRG string) *cache.Cache {
	ttl, err := strconv.Atoi(cacheTTL)
	if err != nil {
		log.Fatal(err)
	}

	prg, err := strconv.Atoi(cachePRG)
	if err != nil {
		log.Fatal(err)
	}

	return cache.New(time.Duration(ttl) * time.Minute, time.Duration(prg) * time.Minute)
}