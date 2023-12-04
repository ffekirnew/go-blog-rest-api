package responses

type GenericResponse struct {
	Message string      `json:"message"`
	Success bool        `json:"success"`
	Value   interface{} `json:"value"`
	Error   string      `json:"error"`
}
