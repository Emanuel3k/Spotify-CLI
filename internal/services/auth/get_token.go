package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/emanuel3k/Spotify-CLI/internal/domain"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func getToken(clientId, clientSecret, urlCallback, code *string) error {
	baseUrl := "https://accounts.spotify.com/api/token"

	body := url.Values{}
	body.Set("code", *code)
	body.Set("redirect_uri", *urlCallback)
	body.Set("grant_type", "authorization_code")

	r, err := http.NewRequest(http.MethodPost, baseUrl, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(*clientId + ":" + *clientSecret))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Set("Authorization", "Basic "+auth)

	client := &http.Client{}
	res, err := client.Do(r)
	if err != nil {
		return fmt.Errorf("failed to get token: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to get token, status code: %d", res.StatusCode)
	}

	var tokenRes domain.SpotifyAccessToken
	err = json.NewDecoder(res.Body).Decode(&tokenRes)
	if err != nil {
		return fmt.Errorf("failed to decode token response: %w", err)
	}

	if err = os.Setenv(spotifyAccessToken, tokenRes.AccessToken); err != nil {
		return fmt.Errorf("failed to set access token in environment variable: %w", err)
	}

	return nil
}
