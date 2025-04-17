package resources

type AlternativeResource struct {
	Userid uint               `json:"userid"`
	Depot  []*DepotAlternatif `json:"alternatives"`
}
