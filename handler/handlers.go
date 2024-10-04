package handler

import (
	"encoding/json"
	"log"
	"net/http"
	shortener "url-shortnener/shortnener"
	"url-shortnener/store"
)

// Request model definition
type UrlCreationRequest struct {
	LongUrl string `json:"long_url" binding:"required"`
	UserId  string `json:"user_id" binding:"required"`
}
type UrlCreationResponse struct {
	Message   string `json:"message"`
	Short_url string `json:"short_url"`
}

func CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	var creationRequest UrlCreationRequest
	if err := json.NewDecoder(r.Body).Decode(&creationRequest); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	shortUrl := shortener.GenerateShortLink(creationRequest.LongUrl, creationRequest.UserId)
	store.SaveUrlMapping(shortUrl, creationRequest.LongUrl, creationRequest.UserId)
	rsp := UrlCreationResponse{
		Short_url: r.Host + shortUrl,
		Message:   "short url created successfully",
	}
	byteData, err := json.Marshal(rsp)
	if err != nil {
		log.Print(err)
		return
	}
	w.Write(byteData)
}

func HandleShortUrlRedirect(w http.ResponseWriter, r *http.Request) {
	shortUrl := r.URL.Query().Get("shortUrl")
	initialUrl := store.RetrieveInitialUrl(shortUrl)
	http.Redirect(w, r, initialUrl, http.StatusSeeOther)
}
