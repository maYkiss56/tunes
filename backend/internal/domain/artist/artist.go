package artist

type Artist struct {
	ID       int
	Nickname string
	BIO      string
	Country  string
}

func NewArtist(nickname, bio, country string) (*Artist, error) {
	return &Artist{
		Nickname: nickname,
		BIO:      bio,
		Country:  country,
	}, nil
}
