// This package provides an implementation of the bowyer_watson algorithm to create a 
// Delaunay Triangulation.
// Given a set of points, it will output a set of Triangles
package bowyer_watson

import (
	"math"
	"container/list"
)

// Basic x,y coordinate 
type Point struct {
	X, Y float64
}

type Triangle struct {
	A, B, C Point
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
	var ab = math.Pow(t.A.X, 2) + math.Pow(t.A.Y, 2)
	var cd = math.Pow(t.B.X, 2) + math.Pow(t.B.Y, 2)
	var ef = math.Pow(t.C.X, 2) + math.Pow(t.C.Y, 2)

	var circum_x = (ab * (t.C.Y - t.B.Y) + cd * (t.A.Y - t.C.Y) + ef * (t.B.Y - t.A.Y)) / (t.A.X * (t.C.Y - t.B.Y) + t.B.X * (t.A.Y - t.C.Y) + t.C.X * (t.B.Y - t.A.Y)) / 2
	var circum_y = (ab * (t.C.X - t.B.X) + cd * (t.A.X - t.C.X) + ef * (t.B.X - t.A.X)) / (t.A.Y * (t.C.X - t.B.X) + t.B.Y * (t.A.X - t.C.X) + t.C.Y * (t.B.X - t.A.X)) / 2
	var circum_radius = math.Sqrt(math.Pow(t.A.X - circum_x, 2) + math.Pow(t.A.Y - circum_y, 2))

	var dist = math.Sqrt(math.Pow(p.X - circum_x, 2) + math.Pow(p.Y - circum_y, 2))
	return dist <= circum_radius
}

// Triangle method
// Determine if one of the Triangle's vertices is the Point p
// Return: True if the Point is equal to one of the vertices
func (t Triangle) ContainsPoint(p Point) bool {
	return t.A == p || t.B == p || t.C == p
}

// Given an array of points, return an array of triangles of the triangulation
// Super triangle is a triangle that contains all the points
// Source for algorithm: paulbourke.net/papers/triangulate
func DelaunayTriangulation(points []Point, super_triangle Triangle) []Triangle {
	triangle_list := list.New()
	triangle_list.PushBack(super_triangle)

	for _, p := range points {
		edge_list := list.New()
		remove_triangles := list.New()

		for itr := triangle_list.Front(); itr != nil; itr = itr.Next() {
			if itr.Value.(Triangle).CircumcircleContains(p) {
				triangle := itr.Value.(Triangle)

				var new_edge [3]Edge

				new_edge[0] = Edge{triangle.A, triangle.B}
				new_edge[1] = Edge{triangle.A, triangle.C}
				new_edge[2] = Edge{triangle.B, triangle.C}

				remove_triangles.PushBack(itr)

				for i := 0; i < 3; i++ {
					edge_list.PushBack(new_edge[i])
				}
			}
		}

		for itr := remove_triangles.Front(); itr != nil; itr = itr.Next() {
			// The iterator points to an element, so dereference and remove from list
			triangle_list.Remove(itr.Value.(*list.Element))
		}

		remove_edges := list.New()
		for itr := edge_list.Front(); itr != nil; itr = itr.Next() {

			left := itr
			if itr.Next() == nil {
				break
			}
			right := itr.Next()
			if left.Value.(Edge).isEqual(right.Value.(Edge)) {
				// Push the *Element onto the list
				remove_edges.PushBack(left)
				remove_edges.PushBack(right)
			}

		}

		for itr := remove_edges.Front(); itr != nil; itr = itr.Next() {
			// The iterator points to an element, so dereference and remove from list
			edge_list.Remove(itr.Value.(*list.Element))
		}

		for itr := edge_list.Front(); itr != nil; itr = itr.Next() {
			new_triangle := Triangle{itr.Value.(Edge).a, itr.Value.(Edge).b, p}
			triangle_list.PushBack(new_triangle)
		}
	}

	remove_triangles := list.New()

	//Remove any triangles using the Points of the supertriangle
	for itr := triangle_list.Front(); itr != nil; itr = itr.Next() {
		if itr.Value.(Triangle).ContainsPoint(super_triangle.A) ||
		   itr.Value.(Triangle).ContainsPoint(super_triangle.B) ||
		   itr.Value.(Triangle).ContainsPoint(super_triangle.C) {	
	   		
	   	    // Push the *Element onto the list
	   	    remove_triangles.PushBack(itr)
		} 
	}

	for itr := remove_triangles.Front(); itr != nil; itr = itr.Next() {
		// The iterator points to an element, so dereference and remove from list
		triangle_list.Remove(itr.Value.(*list.Element))
	}

	return_triangles := make([]Triangle, triangle_list.Len(), triangle_list.Len())

	i := 0
	for itr := triangle_list.Front(); itr != nil; itr = itr.Next() {
		return_triangles[i] = itr.Value.(Triangle)
		i++
	}

	return return_triangles
}
