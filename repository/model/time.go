package model

import "time"

type TTimeAttributes struct {
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	LastUpdatedAt time.Time `json:"last_updated_at" bson:"last_updated_at"`
}

func (t *TTimeAttributes) SetCreatedNow() {
	t.CreatedAt = time.Now()
	t.LastUpdatedAt = time.Now()
}

func (t *TTimeAttributes) SetLastUpdatedNow() {
	t.LastUpdatedAt = time.Now()
}