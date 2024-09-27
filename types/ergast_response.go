package types

type ErgastResponse[T any] struct {
	Data T `json:"MRData,omitempty"`
}
