package book

func validID(id ID) bool {
	if id > minID && id < maxID {
		return true
	}
	return false
}
