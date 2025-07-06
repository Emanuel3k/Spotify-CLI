package user

import (
	"encoding/json"
	"errors"
	"github.com/emanuel3k/Spotify-CLI/internal/domain"
	"github.com/emanuel3k/Spotify-CLI/internal/services"
	"net/http"
)

func GetUserData() (*domain.UserData, error) {

	url := services.SpotifyAPIURL + "/me"

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+services.SpotifyAccessToken)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("failed to get user data, status code: " + res.Status)
	}

	var userData domain.UserData
	if err = json.NewDecoder(res.Body).Decode(&userData); err != nil {
		return nil, err
	}

	return &userData, nil
}
