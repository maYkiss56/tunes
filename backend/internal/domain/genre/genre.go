package genre

type Genre struct {
	ID       int
	Title    string
	ImageURL string
}

func NewGenre(title string, imageURL string) (*Genre, error) {
	return &Genre{
		Title:    title,
		ImageURL: imageURL,
	}, nil
}
