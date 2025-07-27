package scada

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	Position   rl.Vector2
	Dimensions rl.Vector2
}

func (e *Entity) NewRectangle() rl.Rectangle {
	return rl.NewRectangle(e.Position.X, e.Position.Y, e.Dimensions.X, e.Dimensions.Y)
}

func (e *Entity) CheckCollisionWith(other *Entity) bool {
	return e.Position.X < other.Position.X+other.Dimensions.X &&
		e.Position.X+e.Dimensions.X > other.Position.X &&
		e.Position.Y < other.Position.Y+other.Dimensions.Y &&
		e.Position.Y+e.Dimensions.Y > other.Position.Y
}
