package entities

// Order is the domain object for storing pack order
// information. The structure is map[<pack size>]<quantity>
type Order map[int]int
