package structs

type Tagger struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	RelAvatarURL string `json:"relAvatarURL"`
}
