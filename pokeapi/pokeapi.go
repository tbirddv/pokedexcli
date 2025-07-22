package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func FetchStruct(url string, target interface{}) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, fmt.Errorf("failed to fetch data from %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf("received non-200 response code: %d", resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(target); err != nil {
		return 0, fmt.Errorf("failed to decode JSON response: %w", err)
	}

	return 200, nil
}

func GenerateURL(endpoint string) string {
	baseURL := "https://pokeapi.co/api/v2/"
	return fmt.Sprintf("%s%s", baseURL, endpoint)
}
