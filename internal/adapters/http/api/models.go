package api

type calculatePacksRequest struct {
	Items int   `json:"items" binding:"required,number"`
	Packs []int `json:"packs" binding:"required,dive,number"`
}

type packDetails struct {
	Size       int `json:"size"`
	Quantity   int `json:"quantity"`
	TotalItems int `json:"totalItems"`
}

type calculatePacksResponse []packDetails
