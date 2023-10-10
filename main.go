package main

import "os"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}

	Router().Run(":" + port)
}
