package auth

import (
	"fmt"
	"github.com/emanuel3k/Spotify-CLI/internal/services"
	"net/url"
	"os/exec"
	"runtime"
)

func openBrowser(clientId, urlCallback *string) error {

	baseUrl := "https://accounts.spotify.com/authorize?"
	query := url.Values{
		"client_id":     {*clientId},
		"response_type": {"code"},
		"redirect_uri":  {*urlCallback},
		"scope":         {services.Scope},
	}

	openUrl := baseUrl
	for k, v := range query {
		openUrl += fmt.Sprintf("%s=%s&", k, v[0])
	}

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start", "", openUrl}
	case "darwin":
		cmd = "open"
		args = []string{openUrl}
	default:
		cmd = "xdg-open"
		args = []string{openUrl}
	}

	return exec.Command(cmd, args...).Start()
}
