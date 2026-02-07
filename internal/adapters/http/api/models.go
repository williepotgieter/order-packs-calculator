package api

type calculatePacksRequest struct {
	Items uint   `json:"items"`
	Packs []uint `json:"packs"`
}

type packDetails struct {
	Size       uint `json:"size"`
	Quantity   uint `json:"quantity"`
	TotalItems uint `json:"totalItems"`
}

type calculatePacksResponse []packDetails
