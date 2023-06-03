package requestmodel

import "github.com/unJaco/UberClientServer/backend/model"

// anfrage um ride zu buchen muss folgende JSON-Struktur haben
type DriveRequest struct {
	CustomerId    uint           `json:"customer_id"`
	StartLocation model.Location `json:"start_location"`
	EndLocation   model.Location `json:"end_location"`

	// TODO: add vehicle category

}
