package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gookit/goutil/dump"
	"github.com/joho/godotenv"
	"github.com/simonbuckner/dattoapi"
	"github.com/simonbuckner/goquadac"
)

func main() {
	// Load .env files
	// .env.local takes precedence (if present)
	godotenv.Load(".env.local")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	baseUrl := os.Getenv("BASE_URL")
	apiKey := os.Getenv("API_KEY")
	apiSecret := os.Getenv("API_SECRET")

	customerId := goquadac.StringtoI64(os.Getenv("SAAS_CUSTOMER_ID"))
	datto := dattoapi.NewDattoApi(baseUrl)
	datto.Authenticate(apiKey, apiSecret)

	// -------------------------------------------------------------------------
	//   Endpoints starting /v1/saas
	// -------------------------------------------------------------------------

	printBanner("GET:/v1/saas/domains")
	allDomains, err := datto.GetSaasDomains().GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(allDomains)

	printBanner("GET:/v1/saas/{customerId}/seats")
	allSeats, err := datto.GetSaasSeatsForCustomerId(customerId).GetAll()
	if err != nil {
		panic(err)
	}
	dump.Print(allSeats)

	printBanner("GET:/v1/saas/{customerId}/applications")
	allApplications, err := datto.GetSaasApplicationsForCustomerId(customerId).Get()
	if err != nil {
		panic(err)
	}
	dump.Print(allApplications)
}

func printBanner(message string) {
	fmt.Println()
	fmt.Println("-------------------------------------------------------------")
	fmt.Println("    " + message)
	fmt.Println("-------------------------------------------------------------")
	fmt.Println()
}
