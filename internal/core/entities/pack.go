package entities

import "errors"

type Pack struct {
	Size     uint
	Items    uint
	Capacity uint
}

func NewPack(size uint) (*Pack, error) {
	if size == 0 {
		return nil, errors.New("pack size cannot be less than one item")
	}

	return &Pack{
		Size:     size,
		Items:    0,
		Capacity: size,
	}, nil
}

func (p *Pack) Fill(items uint) (overflow uint) {
	if p.Capacity == 0 {
		overflow = items
		return
	}

	if p.Capacity <= items {
		p.Items = p.Size
		overflow = items - p.Capacity
		p.Capacity = 0
		return
	}

	p.Items += items
	p.Capacity -= items
	overflow = 0

	return
}
