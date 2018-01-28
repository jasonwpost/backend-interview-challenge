package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShapeDB(t *testing.T) {
	assert := assert.New(t)

	strShapes := []string{
		"1234|3,5;3,-2;4,2",
		"4321|5,3;7,-9;9,1;-4,3",
	}
	shapeMap := generateShapeDB(strShapes)
	triPoints := []Point{
		Point{3, 5},
		Point{3, -2},
		Point{4, 2},
	}
	squPoints := []Point{
		Point{5, 3},
		Point{7, -9},
		Point{9, 1},
		Point{-4, 3},
	}
	assert.Equal(shapeMap["1234"], Shape{ id: "1234", kind: "Triangle", points: triPoints})
	assert.Equal(shapeMap["4321"], Shape{ id: "4321", kind: "Square", points: squPoints})
}

func TestGenerateOperations(t *testing.T) {
	assert := assert.New(t)

	strOperations := []string{
		"create-shape|1243|4,3;3,4;5,4;4,5",
		"delete-shape|4321",
	}
	operations := generateOperations(strOperations)
	assert.Equal(operations[0], Operation{ kind: "create-shape", shapeId: "1243", operationMetadata: "4,3;3,4;5,4;4,5" })
	assert.Equal(operations[1], Operation{ kind: "delete-shape", shapeId: "4321", operationMetadata: "" })

}

func TestAddPointToShape(t *testing.T) {
	assert := assert.New(t)

	shapePoints := []Point{
		Point{3, 5},
		Point{3, -2},
		Point{4, 2},
	}
	shape := Shape{ id: "1432", kind: "Triangle", points: shapePoints }
	newShape, err := addPointToShape(shape, Point{7, 8})
	assert.Equal(err, nil)
	assert.Equal(newShape.kind, "Square")
	assert.Equal(len(newShape.points), 4)
}

func TestRemovePointFromShape(t *testing.T) {
	assert := assert.New(t)

	shapePoints := []Point{
		Point{3, 5},
		Point{3, -2},
		Point{4, 2},
		Point{2, 4},
	}
	shape := Shape{ id: "1432", kind: "Square", points: shapePoints }
	newShape, err := removePointFromShape(shape, 2)
	assert.Equal(err, nil)
	assert.Equal(newShape.kind, "Triangle")
	assert.Equal(len(newShape.points), 3)
	assert.NotEqual(newShape.points[2], Point{4, 2})
}

func TestPointsFactory(t *testing.T) {
	assert := assert.New(t)

	strSliceOfPoints := []string{
		"3,5",
		"6,4",
		"4,1",
	}
	points := pointsFactory(strSliceOfPoints)
	assert.Equal(points[0], Point{ 3, 5 })
	assert.Equal(points[1], Point{ 6, 4 })
	assert.Equal(points[2], Point{ 4, 1 })
}

func TestCreatePointFromString(t *testing.T) {
	assert := assert.New(t)

	point := createPointFromString("4,3")
	assert.Equal(point, Point{ 4, 3 })
}

func TestShapeFactory(t *testing.T) {
	assert := assert.New(t)

	shapePoints := []Point{
		Point{3, 5},
		Point{3, -2},
		Point{4, 2},
		Point{2, 4},
	}
	shape, err := shapeFactory("8765", shapePoints)
	assert.Equal(err, nil)
	assert.Equal(shape.kind, "Square")
	assert.Equal(shape.id, "8765")
	assert.Equal(shape.points, []Point{
		Point{3, 5},
		Point{3, -2},
		Point{4, 2},
		Point{2, 4},
	})
}

func TestShapeFactoryFailure(t *testing.T) {
	assert := assert.New(t)

	shapePoints := []Point{
		Point{3, 5},
		Point{3, -2},
	}
	shape, err := shapeFactory("111", shapePoints)
	assert.Equal(shape, Shape{})
	assert.Equal(err.Error(), "2 point shape not supported")
}