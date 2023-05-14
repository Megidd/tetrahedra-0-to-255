// Only the first corner of the cube has zero/negative value.

include <shared.scad>;

// Triangles come from here:
// https://github.com/deadsy/sdfx/blob/1a71e404e4b2aa00c59f53cffc219a9e83e62d85/render/march3.go#L360
triangle_table = [
    [ 0, 8, 3 ],
];

draw_triangles(triangle_table);

// Generate tetrahedra by playing around with these:

tetrahedron_table = [
    [ 12 + 0, 0, 3, 8 ],
];

create_tetrahedra(tetrahedron_table);

// Just to be able to visualize the created tetrahedra with respect to the cube.
// Draw cube last.
// https://stackoverflow.com/a/76245980/3405291
color([ 1, 1, 1, 0.2 ]) cube(size, center = false);