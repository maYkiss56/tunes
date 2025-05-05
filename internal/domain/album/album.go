package album

type Album struct {
	ID       int
	Title    string
	ImageURL string
	ArtistID int
}

func NewAlbum(title, imageURL string, artistID int) *Album {
	return &Album{
		Title:    title,
		ImageURL: imageURL,
		ArtistID: artistID,
	}
}
