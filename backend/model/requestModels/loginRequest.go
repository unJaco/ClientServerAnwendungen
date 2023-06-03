package requestmodel

// anfrage zum login muss folgende JSON-Struktur haben
type LoginRequest struct {
	Email  string `json:"email"`
	PwHash string `json:"pw_hash"`
}
