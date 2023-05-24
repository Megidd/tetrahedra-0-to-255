include <shared.scad>;

triangle_table = [ [ 10, 7, 5 ], [ 10, 11, 7 ], [ 9, 8, 1 ], [ 8, 3, 1 ] ];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 12, 1, 3, 8 ],
    [ 12, 13, 1, 8 ],
    [ 13, 1, 8, 9 ],
    [ 10, 11, 5, 18 ],
    [ 11, 7, 5, 18 ],
    [ 11, 7, 18, 19 ],
];

points = tetrahedra_points(tetrahedron_table = tetrahedron_table);

union()
{
    polyhedron(points = points[0], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[1], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[2], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[3], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[4], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[5], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
}

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);