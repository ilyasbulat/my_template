package entity

type Image struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Size  uint64 `json:"size"`
	Bytes []byte `json:"bytes"`
}
