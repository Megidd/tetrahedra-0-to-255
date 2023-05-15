include <shared.scad>;

triangle_table = [
    [ 1, 8, 3 ],
    [ 9, 8, 1 ],
];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 1, 15, 3, 19 ],
    [ 1, 14, 15, 19 ],
    [ 1, 14, 19, 18 ],
    [ 1, 18, 19, 17 ],
    [ 1, 3, 9, 17 ],
    [ 17, 19, 3, 16 ],
    [ 9, 3, 8, 17 ],
    [ 8, 17, 3, 16 ],
];

points = tetrahedra_points(tetrahedron_table = tetrahedron_table);

color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[0], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[1], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[2], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[3], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[4], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[5], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[6], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
color([ 1, 0, 0, 0.5 ]) polyhedron(points = points[7], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);