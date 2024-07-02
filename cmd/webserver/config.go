package main

func localAndServerDevelopment() {
	envVar := os.Getenv("siteonline")
	certDirectory := os.Getenv("certDir")
	if envVar == "site" {
		log.Fatal(app.Listen(":443", fiber.ListenConfig{CertFile: certDirectory + "fullchain.pem", CertKeyFile: certDirectory + "privkey.pem"}))
	} else {
		log.Fatal(app.Listen(":4050"))
	}
}
