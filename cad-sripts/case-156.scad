include <shared.scad>;

triangle_table = [ [ 8, 1, 3 ], [ 8, 6, 1 ], [ 8, 4, 6 ], [ 6, 10, 1 ] ];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 14, 3, 1, 8 ],
    [ 14, 8, 1, 10 ],
    [ 14, 15, 3, 8 ],
    [ 14, 15, 8, 10 ],
    [ 10, 15, 8, 6 ],
    [ 15, 8, 6, 19 ],
    [ 6, 19, 8, 4 ],
    [ 19, 8, 4, 16 ],
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
    polyhedron(points = points[6], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[7], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
}

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);