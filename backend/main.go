package main

import "api/db"

func main() {
	db.ConexaoBD()
	routes.HandleRequest()
}
