package types

type Stop struct {
    Latitude  float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
}

type StopStore interface {
    CreateStop(stop Stop) error
}