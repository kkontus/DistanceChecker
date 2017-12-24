package entity

type Customer struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	UserId    int    `json:"user_id"`
	Name      string `json:"name"`
	Distance  GeoDist
}
