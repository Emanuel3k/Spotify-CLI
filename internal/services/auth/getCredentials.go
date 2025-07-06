package auth

import (
	"errors"
	"github.com/emanuel3k/Spotify-CLI/internal/services"
	"net/http"
	"os"
	"time"
)

func GetCredentials() error {
	clientId := os.Getenv(services.SpotifyClientId)
	urlCallback := os.Getenv(services.SpotifyURLCallback)
	clientSecret := os.Getenv(services.SpotifyClientSecret)

	ch := make(chan string)
	server := &http.Server{Addr: ":8080"}

	http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Query().Get("code")
		if code != "" {
			ch <- code
			w.Write([]byte("Autorização recebida. Você pode fechar esta janela."))
			go server.Close()
		} else {
			http.Error(w, "Código não encontrado", http.StatusBadRequest)
		}
	})

	go func() {
		if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			panic("failed to start HTTP server for Spotify callback: " + err.Error())
		}
	}()

	if err := openBrowser(&clientId, &urlCallback); err != nil {
		return err
	}

	select {
	case code := <-ch:
		if err := getToken(&clientId, &clientSecret, &urlCallback, &code); err != nil {
			return err
		}
	case <-time.After(2 * time.Minute):
		return errors.New("tempo esgotado para autorização do Spotify")
	}

	return nil
}
