// Only the first corner of the cube has zero/negative value.

include <cube.scad>;
include <tetrahedron.scad>

tetrahedron_count = 6;

tetrahedron_indices = [
        [12+0, 12+4, 12+7, 12+6],
        [12+0, 12+3, 12+7, 12+6],
        [12+0, 12+4, 12+5, 12+6],
        [12+0, 12+1, 12+5, 12+6],
        [12+0, 12+3, 12+2, 12+6],
        [12+0, 12+1, 12+2, 12+6],
    ];

tetrahedron_colors = [
    [1, 0, 0, alpha],  // red
    [0, 1, 0, alpha],  // green
    [0, 0, 1, alpha],  // blue
    [1, 1, 0, alpha],  // yellow
    [1, 0, 1, alpha],  // magenta
    [0, 1, 1, alpha]   // cyan
];

create_tetrahedra(
    tetrahedron_count = tetrahedron_count,
    tetrahedron_indices = tetrahedron_indices,
    tetrahedron_colors = tetrahedron_colors
);

// Just to be able to visualize the created tetrahedra with respect to the cube.
labeled_cube();
