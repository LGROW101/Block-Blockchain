package main

import (
	"net/http"

	"github.com/LGROW101/Block-Blockchain/blockchain"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const DIFFICULTY = 2

var bc *blockchain.BlockChain

func main() {
	bc = blockchain.New(DIFFICULTY)

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/blockchain", getBlockChain)
	e.POST("/wallet", createWallet)
	e.POST("/transaction", createTransaction)
	e.GET("/validate", validateChain)

	e.Logger.Fatal(e.Start(":8080"))
}

func getBlockChain(c echo.Context) error {
	return c.JSON(http.StatusOK, bc.Blocks())
}

func createWallet(c echo.Context) error {
	address := bc.CreateWallet()
	return c.JSON(http.StatusCreated, map[string]string{"address": address})
}

func createTransaction(c echo.Context) error {
	type Request struct {
		From   string  `json:"from"`
		To     string  `json:"to"`
		Amount float64 `json:"amount"`
	}

	var r Request
	if err := c.Bind(&r); err != nil {
		return err
	}

	if err := bc.CreateTransaction(r.From, r.To, r.Amount); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]string{"message": "transaction created"})
}

func validateChain(c echo.Context) error {
	isValid := bc.IsValid()
	return c.JSON(http.StatusOK, map[string]bool{"isValid": isValid})
}
