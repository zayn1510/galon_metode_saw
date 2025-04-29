package resources

type DepotData struct {
	Alternatif          []*DepotAlternatif     `json:"alternatif,omitempty"`
	NormalisasiResource []*NormalisasiResource `json:"normalisasi,omitempty"`
	Hasil               []*HasilSawResource    `json:"hasil,omitempty"`
}
type NormalisasiResource struct {
	IDDepot  uint    `json:"id_depot"`
	Depot    string  `json:"depot"`
	Harga    float64 `json:"harga,omitempty"`
	Jarak    float64 `json:"jarak,omitempty"`
	Diskon   float64 `json:"diskon,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Distance float64 `json:"distance,omitempty"`
}
type HasilSawResource struct {
	IDDepot        uint    `json:"id_depot"`
	Depot          string  `json:"depot"`
	Harga          int     `json:"harga,omitempty"`
	Jarak          string  `json:"jarak,omitempty"`
	Diskon         int     `json:"diskon,omitempty"`
	Rating         float64 `json:"rating,omitempty"`
	Nilai          float64 `json:"nilai"`
	Image          string  `json:"image"`
	Alamat         string  `json:"alamat"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	NomorHandphone string  `json:"nomor_handphone"`
	UpdatedAt      string  `json:"updated_at"`
}
