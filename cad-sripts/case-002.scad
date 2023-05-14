include <shared.scad>;

// Triangles come from here:
// https://github.com/deadsy/sdfx/blob/1a71e404e4b2aa00c59f53cffc219a9e83e62d85/render/march3.go#L360
triangle_table = [0, 1, 9];
draw_triangles(indices = triangle_table, vertices = edges, triangle_color=[1, 1, 1, .4]);

// Generate tetrahedra by playing around with these:

tetrahedron_count = 1;

tetrahedron_indices = [
    [13, 1, 0, 9],
];

tetrahedron_colors = [
    [1, 0, 0, 0.5],  // red
];

create_tetrahedra(tetrahedron_count,tetrahedron_indices,tetrahedron_colors);

// Just to be able to visualize the created tetrahedra with respect to the cube.
// Draw cube last.
// https://stackoverflow.com/a/76245980/3405291
labeled_cube();
