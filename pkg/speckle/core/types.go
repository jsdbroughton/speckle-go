package core

type UserInfo struct {
	Name    string
	Email   string
	Company string
	ID      string
}

type ServerInfo struct {
	URL string
	// Add other necessary fields
}

type Account struct {
	IsDefault    bool
	Token        string
	RefreshToken string
	ServerInfo   ServerInfo
	UserInfo     UserInfo
	ID           string
}
