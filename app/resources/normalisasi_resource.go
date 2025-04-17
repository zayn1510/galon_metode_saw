package resources

type DepotData struct {
	Alternatif          []*DepotAlternatif     `json:"alternatif"`
	NormalisasiResource []*NormalisasiResource `json:"normalisasi"`
	Hasil               []*HasilSawResource    `json:"hasil"`
}
type NormalisasiResource struct {
	IDDepot  uint    `json:"id_depot"`
	Depot    string  `json:"depot"`
	Harga    float64 `json:"harga,omitempty"`
	Jarak    float64 `json:"jarak"`
	Diskon   float64 `json:"diskon,omitempty"`
	Rating   float64 `json:"rating,omitempty"`
	Distance float64 `json:"distance,omitempty"`
}
type HasilSawResource struct {
	IDDepot uint    `json:"id_depot"`
	Depot   string  `json:"depot"`
	Harga   int     `json:"harga,omitempty"`
	Jarak   string  `json:"jarak,omitempty"`
	Diskon  int     `json:"diskon,omitempty"`
	Rating  float64 `json:"rating,omitempty"`
	Nilai   float64 `json:"nilai"`
}
