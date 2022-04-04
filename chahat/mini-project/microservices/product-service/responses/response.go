package responses

type ProductResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
    Data    map[string]interface{} `json:"data"`
}