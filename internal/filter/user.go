package filter

// UserListFilter
func UserListFilter(username string) map[string]interface{} {
	filter := make(map[string]interface{})
	if username != "" {
		filter["username"] = username
	}
	return filter
}
