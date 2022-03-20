package responses

type LoginResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Token   string                 `json:"token"`
	Data    map[string]interface{} `json:"data"`
}
