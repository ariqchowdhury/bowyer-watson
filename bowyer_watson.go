// This package provides an implementation of the bowyer_watson algorithm to create a 
// Delaunay Triangulation.
// Given a set of points, it will output a set of Triangles
package bowyer_watson

import (
	"os"
	"bufio"
	"fmt"
	"math"
	"container/list"
)

// Basic x,y coordinate 
type Point struct {
	x, y float64
}

type Triangle struct {
	a, b, c Point
}

type Edge struct {
	a, b Point
}

// Edge method
// Determines if Edge, e2, is an equivalent edge
// Return: True if equal
func (e1 Edge) isEqual(e2 Edge) bool {
	return (e1.a == e2.a && e1.b == e2.b || e1.a == e2.b && e1.b == e2.a)
}

// Triangle method
// Determines if a given Point is contained within the circumcircle of the triangle
// A circumcircle is the circle whose circumference contains all 3 vertices of a triangle
// Return: True if point is contained
func (t Triangle) CircumcircleContains(p Point) bool {
	var ab = math.Pow(t.a.x, 2) + math.Pow(t.a.y, 2)
	var cd = math.Pow(t.b.x, 2) + math.Pow(t.b.y, 2)
	var ef = math.Pow(t.c.x, 2) + math.Pow(t.c.y, 2)

	var circum_x = (ab * (t.c.y - t.b.y) + cd * (t.a.y - t.c.y) + ef * (t.b.y - t.a.y)) / (t.a.x * (t.c.y - t.b.y) + t.b.x * (t.a.y - t.c.y) + t.c.x * (t.b.y - t.a.y)) / 2
	var circum_y = (ab * (t.c.x - t.b.x) + cd * (t.a.x - t.c.x) + ef * (t.b.x - t.a.x)) / (t.a.y * (t.c.x - t.b.x) + t.b.y * (t.a.x - t.c.x) + t.c.y * (t.b.x - t.a.x)) / 2
	var circum_radius = math.Sqrt(math.Pow(t.a.x - circum_x, 2) + math.Pow(t.a.y - circum_y, 2))

	var dist = math.Sqrt(math.Pow(p.x - circum_x, 2) + math.Pow(p.y - circum_y, 2))
	return dist <= circum_radius
}

// Triangle method
// Determine if one of the Triangle's vertices is the Point p
// Return: True if the Point is equal to one of the vertices
func (t Triangle) ContainsPoint(p Point) bool {
	return t.a == p || t.b == p || t.c == p
}

// Given an array of points, return an array of triangles of the triangulation
func DelaunayTriangulation(points []Point) []Triangle {

}