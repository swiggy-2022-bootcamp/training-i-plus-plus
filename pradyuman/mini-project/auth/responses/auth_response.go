package responses

type AuthResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
	JWT     string                 `json:"jwt"`
}
