package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/rearmid/topup-game/internal/external/viptopup"
)

// mockStats is a simple stats collector.
// Replace with your actual stats implementation.
type mockStats struct{}

func (m *mockStats) Inc(name string, value int64) {
	log.Printf("[STATS] %s: %d", name, value)
}

func main() {

	apiID := "<API ID>"
	apiKey := "<API KEY>"

	log := slog.Default()
	log.Info("Starting VIPTopup example with API ID", "api_id", apiID)

	// Create a new VIPTopup client with default options
	// You can also use functional options: viptopup.WithEndpoint("..."), viptopup.WithHTTPClient(customClient)
	client := viptopup.New(
		apiID,
		apiKey,
		viptopup.WithHTTPClient(&http.Client{}), // Using default http client
		viptopup.WithLogger(log),
		viptopup.WithStats(&mockStats{}),
	)

	ctx := context.Background()

	// --- Service Game & Streaming ---
	fmt.Println("\n--- Service Game (no filters) ---")
	respServiceGame, err := client.ServiceGame(ctx, nil, nil, nil)
	if err != nil {
		log.Error("Error getting service game", "error", err)
		return
	}
	if respServiceGame.Status {
		for _, data := range respServiceGame.Data {
			fmt.Printf("Service Code: %s\n", data.Code)
			fmt.Printf("Service Game / Streaming: %s\n", data.Game)
			fmt.Printf("Service Name: %s\n", data.Name)
			fmt.Printf("Price Basic: %d\n", data.Price.Basic)
			fmt.Printf("Price Premium: %d\n", data.Price.Premium)
			fmt.Printf("Price Special: %d\n", data.Price.Special)
			fmt.Printf("Server: %s\n", data.Server)
			fmt.Printf("Status: %s\n", data.Status)
			fmt.Printf("Note: %s\n", data.Note)
			fmt.Println("---")
		}
		fmt.Printf("Message: %s\n", respServiceGame.Message)
	} else {
		fmt.Printf("Error: %s\n", respServiceGame.Message)
	}

	fmt.Println("\n--- Service Game (with filters) ---")
	filterType := "game"
	filterValue := "Mobile Legend"
	filterStatus := "available"

	respServiceGameFiltered, err := client.ServiceGame(ctx, &filterType, &filterValue, &filterStatus)
	if err != nil {
		log.Error("Error getting filtered service game", "error", err)
		return
	}

	if respServiceGameFiltered.Status {
		for _, data := range respServiceGameFiltered.Data {
			fmt.Printf("Service Code: %s\n", data.Code)
			fmt.Printf("Service Game / Streaming: %s\n", data.Game)
			fmt.Printf("Service Name: %s\n", data.Name)
			fmt.Printf("Price Basic: %d\n", data.Price.Basic)
			fmt.Printf("Price Premium: %d\n", data.Price.Premium)
			fmt.Printf("Price Special: %d\n", data.Price.Special)
			fmt.Printf("Server: %s\n", data.Server)
			fmt.Printf("Status: %s\n", data.Status)
			fmt.Printf("Note: %s\n", data.Note)
			fmt.Println("---")
		}
		fmt.Printf("Message: %s\n", respServiceGameFiltered.Message)
	} else {
		fmt.Printf("Error: %s\n", respServiceGameFiltered.Message)
	}

	// Example of Profile call (optional, as it was not in the PHP example snippet for game services)
	fmt.Println("\n--- Profile Information ---")

	profileResp, err := client.Profile(ctx)
	if err != nil {
		log.Error("Error fetching profile", "error", err)
		return
	} else {
		if profileResp.Status {
			fmt.Println("Profile data:", profileResp.Data)
			fmt.Println("Message:", profileResp.Message)
		} else {
			fmt.Println("Error fetching profile:", profileResp.Message)
		}
	}
}
