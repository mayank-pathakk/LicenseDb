package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/fossology/licenseDb/backend/internal/system"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("/Users/mayankpathak/Desktop/LicenseDb/.env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	flag.Parse()

	if os.Getenv("TOKEN_HOUR_LIFESPAN") == "" || os.Getenv("API_SECRET") == "" || os.Getenv("DEFAULT_ISSUER") == "" {
		log.Fatal("Mandatory environment variables not configured")
	}

	// if os.Getenv("JWKS_URI") != "" {
	// 	cache, err := jwk.NewCache(context.Background(), httprc.NewClient())
	// 	if err != nil {
	// 		log.Fatalf("Failed to create a jwk.Cache from the oidc provider's URL: %s", err)
	// 	}

	// 	if err := cache.Register(context.Background(), os.Getenv("JWKS_URI")); err != nil {
	// 		log.Fatalf("Failed to create a jwk.Cache from the oidc provider's URL: %s", err)
	// 	}

	// 	auth.Jwks = cache
	// }

	system.InitDependencies()
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
