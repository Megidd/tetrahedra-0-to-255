include <shared.scad>;

triangle_table = [ [ 8, 4, 7 ], [ 3, 11, 2 ] ];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 8, 4, 7, 16 ],
    [ 2, 15, 3, 11 ],
];

points = tetrahedra_points(tetrahedron_table = tetrahedron_table);

union()
{
    polyhedron(points = points[0], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    polyhedron(points = points[1], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
}

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);