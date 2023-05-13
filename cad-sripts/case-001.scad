// Only the first corner of the cube has zero/negative value.

include <cube.scad>;
include <triangle.scad>;
include <tetrahedron.scad>;

triangle_table = [0, 8, 3];

draw_triangles(indices = triangle_table, vertices = edges);

tetrahedron_count = 1;

tetrahedron_indices = [
    [12+0, 0, 3, 8],
];

tetrahedron_colors = [
    [1, 0, 0, 0.5],  // red
];

create_tetrahedra(
    tetrahedron_count = tetrahedron_count,
    tetrahedron_indices = tetrahedron_indices,
    tetrahedron_colors = tetrahedron_colors
);

// Just to be able to visualize the created tetrahedra with respect to the cube.
labeled_cube();
