package authentication

// TODO: should be changed to a better storage for production
var storage []string

func StoreAntiForgeryToken(token string) {
	storage = append(storage, token)
}

func ValidateAntiForgeryToken(token string) bool {
	// TODO: should also check that token was issued for this session
	// probably pass with the cookies or start tracking sessions in DB

	for _, t := range storage {
		if t == token {
			return true
		}
	}
	return false
}
