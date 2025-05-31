package viptopup_test

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

func ExampleVIPTopup() {
	// API credentials (replace with your actual credentials)
	apiID := "YOUR_API_ID"
	apiKey := "YOUR_API_KEY"

	log := slog.Default()

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

	// --- Order Game Feature & Streaming ---
	fmt.Println("\n--- Ordering Game (with data_zone) ---")
	orderGameZone := "1111"
	respOrderGame, err := client.OrderGame(ctx, "ML5_0-S10", "111111", &orderGameZone)
	if err != nil {
		log.Error("Error ordering game", "error", err)
		return
	}

	if respOrderGame.Status {
		fmt.Printf("TRX ID: %s\n", respOrderGame.Data.TrxID)
		fmt.Printf("Service Code: %s\n", respOrderGame.Data.Code)
		fmt.Printf("Data: %s\n", respOrderGame.Data.DataNo)
		fmt.Printf("Data Zone: %s\n", respOrderGame.Data.DataZone)
		fmt.Printf("Price: %d\n", respOrderGame.Data.Price)
		fmt.Printf("Note: %s\n", respOrderGame.Data.Note)
		fmt.Printf("Last Balance: %d\n", respOrderGame.Data.Balance)
		fmt.Printf("Message: %s\n", respOrderGame.Message)
	} else {
		fmt.Printf("Error: %s\n", respOrderGame.Message)
	}

	fmt.Println("\n--- Ordering Game (without data_zone) ---")
	respOrderGameNoZone, err := client.OrderGame(ctx, "FF5-S24", "111111", nil)
	if err != nil {
		log.Error("Error ordering game", "error", err)
		return
	}
	if respOrderGameNoZone.Status {
		fmt.Printf("TRX ID: %s\n", respOrderGameNoZone.Data.TrxID)
		fmt.Printf("Service Code: %s\n", respOrderGameNoZone.Data.Code)
		fmt.Printf("Data: %s\n", respOrderGameNoZone.Data.DataNo)
		fmt.Printf("Price: %d\n", respOrderGameNoZone.Data.Price)
		fmt.Printf("Note: %s\n", respOrderGameNoZone.Data.Note)
		fmt.Printf("Last Balance: %d\n", respOrderGameNoZone.Data.Balance)
		fmt.Printf("Message: %s\n", respOrderGameNoZone.Message)
	} else {
		fmt.Printf("Error: %s\n", respOrderGameNoZone.Message)
	}

	// --- Status Order Game & Streaming ---
	fmt.Println("\n--- Status Order Game (single TRX_ID) ---")

	// Replace "TRX_ID_FROM_ORDER_GAME" with an actual transaction ID from a previous order
	trxIDForStatus := "TRX_ID_FROM_ORDER_GAME" // Example: respOrderGame.Data.TrxID
	respStatusOrder, err := client.StatusOrderGame(ctx, trxIDForStatus, nil)
	if err != nil {
		log.Error("Error getting status order game", "error", err)
		return
	}

	if respStatusOrder.Status {
		for _, data := range respStatusOrder.Data {
			fmt.Printf("TRX ID: %s\n", data.TrxID)
			fmt.Printf("Service Code: %s\n", data.Code)
			fmt.Printf("Data: %s\n", data.DataNo)
			fmt.Printf("Price: %d\n", data.Price)
			fmt.Printf("Note: %s\n", data.Note)
		}
		fmt.Printf("Message: %s\n", respStatusOrder.Message)
	} else {
		fmt.Printf("Error: %s\n", respStatusOrder.Message)
	}

	fmt.Println("\n--- Status Order Game (TRX_ID with limit) ---")

	limit := 5
	respStatusOrderLimit, err := client.StatusOrderGame(ctx, trxIDForStatus, &limit)
	if err != nil {
		log.Error("Error getting status order game with limit", "error", err)
		return
	}

	if respStatusOrderLimit.Status {
		for _, data := range respStatusOrderLimit.Data {
			fmt.Printf("TRX ID: %s\n", data.TrxID)
			fmt.Printf("Service Code: %s\n", data.Code)
			fmt.Printf("Data: %s\n", data.DataNo)
			fmt.Printf("Price: %d\n", data.Price)
			fmt.Printf("Note: %s\n", data.Note)
		}
		fmt.Printf("Message: %s\n", respStatusOrderLimit.Message)
	} else {
		fmt.Printf("Error: %s\n", respStatusOrderLimit.Message)
	}

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
