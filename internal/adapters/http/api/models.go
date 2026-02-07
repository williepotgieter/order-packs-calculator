package api

type calculatePacksRequest struct {
	Items uint   `json:"items" binding:"required,number"`
	Packs []uint `json:"packs" binding:"required,dive,number"`
}

type packDetails struct {
	Size       uint `json:"size"`
	Quantity   uint `json:"quantity"`
	TotalItems uint `json:"totalItems"`
}

type calculatePacksResponse []packDetails
