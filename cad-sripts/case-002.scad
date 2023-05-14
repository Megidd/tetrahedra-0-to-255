include <shared.scad>;

triangle_table = [
    [ 0, 1, 9 ],
];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 13, 1, 0, 9 ],
];

draw_tetrahedra(tetrahedron_table);

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);