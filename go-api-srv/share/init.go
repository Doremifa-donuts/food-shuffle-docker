package main

import (
	"fmt"
	"log"
	"/log"
	"github.com/joho/godotenv"

	"github.com/go-playground/locales/my"
)

func Init() {
	// ログの設定


	// envから読み込む
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	fmt.println ("jwtSecretKey", jwtSecretKey)

}