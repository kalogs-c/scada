package entity_test

import (
	"testing"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/kalogs-c/scada/entity"
)

func TestNewRectangle(t *testing.T) {
	e := entity.Entity{
		Position:   rl.Vector2{X: 10, Y: 20},
		Dimensions: rl.Vector2{X: 30, Y: 40},
	}

	rect := e.NewRectangle()

	if rect.X != e.Position.X || rect.Y != e.Position.Y {
		t.Errorf("Expected rectangle position (%f, %f), got (%f, %f)", e.Position.X, e.Position.Y, rect.X, rect.Y)
	}

	if rect.Width != e.Dimensions.X || rect.Height != e.Dimensions.Y {
		t.Errorf("Expected rectangle dimensions (%f, %f), got (%f, %f)", e.Dimensions.X, e.Dimensions.Y, rect.Width, rect.Height)
	}
}

func TestCheckCollisionWith(t *testing.T) {
	tests := []struct {
		name     string
		e1       entity.Entity
		e2       entity.Entity
		expected bool
	}{
		{
			name: "Colliding entities",
			e1: entity.Entity{
				Position:   rl.Vector2{X: 0, Y: 0},
				Dimensions: rl.Vector2{X: 10, Y: 10},
			},
			e2: entity.Entity{
				Position:   rl.Vector2{X: 5, Y: 5},
				Dimensions: rl.Vector2{X: 10, Y: 10},
			},
			expected: true,
		},
		{
			name: "Non-colliding entities",
			e1: entity.Entity{
				Position:   rl.Vector2{X: 0, Y: 0},
				Dimensions: rl.Vector2{X: 10, Y: 10},
			},
			e2: entity.Entity{
				Position:   rl.Vector2{X: 20, Y: 20},
				Dimensions: rl.Vector2{X: 5, Y: 5},
			},
			expected: false,
		},
		{
			name: "Touching edges but not overlapping",
			e1: entity.Entity{
				Position:   rl.Vector2{X: 0, Y: 0},
				Dimensions: rl.Vector2{X: 10, Y: 10},
			},
			e2: entity.Entity{
				Position:   rl.Vector2{X: 10, Y: 0},
				Dimensions: rl.Vector2{X: 5, Y: 5},
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.e1.CheckCollisionWith(&tt.e2)
			if result != tt.expected {
				t.Errorf("Expected collision = %v, got %v", tt.expected, result)
			}
		})
	}
}
