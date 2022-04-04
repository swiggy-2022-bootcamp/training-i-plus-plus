package responses

type TrackStreamResponse struct {
    Status  int                    `json:"status"`
    Message string                 `json:"message"`
	Digital int64                     `json:"digital"`
	COD     int64                     `json:"cod"`
  
}