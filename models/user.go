package models

import "time"

// User struct'ı veritabanı tablosu ile JSON istekleri arasındaki köprüdür.
type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	// omitempty: JSON'da bu alan yoksa veya boşsa hata vermez, görmezden gelir.
	// Bu, POST isteklerinde tarih göndermediğimiz için çok kritiktir.
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}