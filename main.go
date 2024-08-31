package main

import app "github.com/arun-murali0/go-lang-backend/app"

func main() {
	err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}

}
