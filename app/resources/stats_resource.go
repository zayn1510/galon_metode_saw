package resources

type StatsResource struct {
	Kriteria  int64 `json:"kriteria"`
	Users     int64 `json:"users"`
	Kecamatan int64 `json:"kecamatan"`
	Depot     int64 `json:"depot"`
}
