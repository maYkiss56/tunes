package genre

type Genre struct {
	ID    int
	Title string
}

func NewGenre(title string) (*Genre, error) {
	return &Genre{Title: title}, nil
}
