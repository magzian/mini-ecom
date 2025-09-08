package requests

import "backend/models"

type Signup struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AddPermission struct {
	Username   string             `json:"username"`
	Permission models.Permissions `json:"permission"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
