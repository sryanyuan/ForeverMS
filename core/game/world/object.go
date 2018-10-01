package world

import "image"

type IMapObject interface {
	GetObjectID() int
	GetPosition() *image.Point
	GetType() int
}

type BaseMapObject struct {
	objectID int
	typ      int
	position image.Point
}

func (o *BaseMapObject) GetObjectID() int {
	return o.objectID
}

func (o *BaseMapObject) GetPosition() *image.Point {
	return &o.position
}

func (o *BaseMapObject) GetType() int {
	return o.typ
}
