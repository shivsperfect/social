package main

import "log"

func main() {
	cfg := config{
		addr: ":3000",
	}

	app := &app{
		config: cfg,
	}

	log.Fatal(app.run(app.mount()))
}
