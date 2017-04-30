package types

type GetWallUploadServer struct {
	Response struct {
		UploadURL string `json:"upload_url"`
		AlbumID   int    `json:"album_id"`
		UserID    int    `json:"user_id"`
	} `json:"response"`
}

type UploadResponse struct {
	Server int    `json:"server"`
	Photo  string `json:"photo"`
	Hash   string `json:"hash"`
}

type SavedPhoto struct {
	Response []struct {
		ID      int    `json:"id"`
		AlbumID int    `json:"album_id"`
		OwnerID int    `json:"owner_id"`
		UserID  int    `json:"user_id"`
		Text    string `json:"text"`
		Date    int    `json:"date"`
	} `json:"response"`
}
