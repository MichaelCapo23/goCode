package models

type Reward struct {
	ID          int
	PointCount  int
	Claimed     int
	ClaimedDate string
	CreatedAt   string
}
