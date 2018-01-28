package main

import (
    "fmt"
    "io/ioutil"
    "strings"
    "strconv"
    "sort"
    "errors"
)

type Point struct {
	x int
	y int
}

type Shape struct {
	id string
	kind string
	points [] Point
}

type Operation struct {
	kind string
	shapeId string
	operationMetadata string
}

func main() {
   shapesList := importShapes()
   shapeDB := generateShapeDB(shapesList)
   fmt.Println("----Shapes as loaded----")
   printTotalByShapeKind(shapeDB)
   operations := generateOperations(importOperations())
   shapeDB = performOperations(shapeDB, operations)
   fmt.Println("----Shapes after operations----")
   printTotalByShapeKind(shapeDB)
   exportShapes(shapeDB)
}

func importShapes() []string {
	shapesBinary, err := ioutil.ReadFile("shapes.db")
    if err != nil {
        panic(err)
    }
    return strings.Split(string(shapesBinary), "\n")
}

func exportShapes(shapes map[string]Shape) {
	var shapesStrings []string

	for k := range shapes {
		var strPoints []string
		for j := 0; j < len(shapes[k].points); j++ {
			pointAsArray := []string{ strconv.Itoa(shapes[k].points[j].x), strconv.Itoa(shapes[k].points[j].y) }
			pointAsString := strings.Join(pointAsArray, ",")
			strPoints = append(strPoints, pointAsString)
		}
		strShape := fmt.Sprintf("%s|%s\n", k, strings.Join(strPoints, ";"))
		shapesStrings = append(shapesStrings, strShape)
	}
	sort.Strings(shapesStrings)
	shapeBytes := []byte(strings.Join(shapesStrings, ""))
	err := ioutil.WriteFile("updatedShapes.db", shapeBytes, 0644)
	if err != nil {
		panic(err)
	}
	return
}

func importOperations() []string {
	operationsBinary, err := ioutil.ReadFile("operations.data")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(operationsBinary), "\n")
}

func generateShapeDB(strShapes []string) map[string]Shape {
	shapeMap := make(map[string]Shape)
   	for i := 0; i < len(strShapes); i++ {
   		toProcess := strings.Split(strShapes[i], "|")
   		id := toProcess[0]
   		strPoints := strings.Split(toProcess[1], ";")
   		pointsList := pointsFactory(strPoints)
   		shape, err := shapeFactory(id, pointsList)
   		if err == nil {
   			shapeMap[id] = shape
   		} else {
   			fmt.Println("Not a valid operation: %s", err)
   		}
   }
   return shapeMap
}

func generateOperations(strOperations []string) []Operation {
	var operationDB []Operation
	for i := 0; i < len(strOperations); i++ {
   		toProcess := strings.Split(strOperations[i], "|")
   		if (len(toProcess) == 2) {
   				operationDB = append(operationDB, Operation{ toProcess[0], toProcess[1], "" })
   			} else {
   				operationDB = append(operationDB, Operation{ toProcess[0], toProcess[1], toProcess[2] })
   			}
   }
   return operationDB
}

func performOperations(shapes map[string]Shape, operations []Operation) map[string]Shape {
	shapeDB := shapes;
	for i := 0; i < len(operations); i++ {
		switch(operations[i].kind){
			case "create-shape":
				newShape, err := shapeFactory(operations[i].shapeId, pointsFactory(strings.Split(operations[i].operationMetadata, ";")))
				if err == nil {
					shapeDB[operations[i].shapeId] = newShape
				} else {
					fmt.Println("Not a valid operation:", err)
				}
			case "delete-shape":
				delete(shapeDB, operations[i].shapeId)
			case "add-point":
				updatedShape, err := addPointToShape(shapeDB[operations[i].shapeId], createPointFromString(operations[i].operationMetadata))
				if err == nil {
					shapeDB[operations[i].shapeId] = updatedShape
				} else {
					fmt.Println("Not a valid operation:", err)
				}
			case "delete-point":
				pointToRemove, err := strconv.Atoi(operations[i].operationMetadata)
				if err != nil {
					panic(err)
				}
				updatedShape, err := removePointFromShape(shapeDB[operations[i].shapeId], pointToRemove)
				if err == nil {
					shapeDB[operations[i].shapeId] = updatedShape
				} else {
					fmt.Println("Not a valid operation:", err)
				}
			default:
				panic(fmt.Sprintf("Operation %s not supported", operations[i].kind))
		}
	}
	return shapeDB
}

func addPointToShape(shape Shape, newPoint Point) (Shape, error) {
	newPoints := append(shape.points, newPoint)
	return shapeFactory(shape.id, newPoints)
}

func removePointFromShape(shape Shape, indexToRemove int) (Shape, error) {
	newPoints := append(shape.points[:indexToRemove], shape.points[indexToRemove+1:]...)
	return shapeFactory(shape.id, newPoints)
}
 
func pointsFactory(strPoints []string) []Point {
	var points []Point
	for i := 0; i < len(strPoints); i++ {
		point := createPointFromString(strPoints[i])
		points = append(points, point)
	}
	return points
}

func createPointFromString(strPoints string) Point {
	xy := strings.Split(strPoints, ",")
		x, xErr := strconv.Atoi(xy[0])
		y, yErr := strconv.Atoi(xy[1])
		if xErr != nil {
			panic(xErr)
		}
		if yErr != nil {
			panic(yErr)
		}
		return Point{ x, y }
}

func shapeFactory(id string, points []Point) (shape Shape, err error) {
	switch len(points) {
		case 3:
			shape = Shape{id, "Triangle", points}
		case 4:
			shape = Shape{id, "Square", points}
		case 5:
			shape = Shape{id, "Pentagon", points}
		case 6:
			shape = Shape{id, "Hexagon", points}
		case 7:
			shape = Shape{id, "Heptagon", points}
		default:
			err = errors.New(fmt.Sprintf("%d point shape not supported", len(points)))
	}
	return
}

func printTotalByShapeKind(shapeDB map[string]Shape) {
	shapeMap := make(map[string][]Shape)

	for k, v := range shapeDB {
		shapeMap[shapeDB[k].kind] = append(shapeMap[shapeDB[k].kind], v)
	}
	
	for k := range shapeMap {
		fmt.Printf("Shape:%s, Count: %d \n", k, len(shapeMap[k]))
	}
	return
}