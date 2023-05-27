include <shared.scad>;

triangle_table = [
    [ 0, 8, 3 ],
];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 12, 0, 3, 8 ],
];

points = tetrahedra_points(tetrahedron_table = tetrahedron_table);

union()
{
    polyhedron(points = points[0], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
}

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);