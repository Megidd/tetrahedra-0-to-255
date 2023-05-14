include <shared.scad>;

triangle_table = [
    [ 0, 8, 3 ],
];

draw_triangles(triangle_table);

tetrahedron_table = [
    [ 12 + 0, 0, 3, 8 ],
];

draw_tetrahedra(tetrahedron_table);

color([ 1, 1, 1, 0.2 ]) cube(size, center = false);