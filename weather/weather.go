package weather

import (
    "fmt"
    "io"
    "net/http"
    "time"
)

// GetWeather fetches weather information from wttr.in for the given country and city.
// Country can be empty. City must not be empty.
func GetWeather(country, city string) string {
    location := city
    if country != "" {
        location = fmt.Sprintf("%s,%s", city, country)
    }
    url := fmt.Sprintf("https://wttr.in/%s?format=3", location)

    client := &http.Client{Timeout: 10 * time.Second}
    resp, err := client.Get(url)
    if err != nil {
        return fmt.Sprintf("request error: %v", err)
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return fmt.Sprintf("read error: %v", err)
    }
    return string(body)
}

