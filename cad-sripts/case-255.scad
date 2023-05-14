// When all 8 corners of the cube have zero or negative values.
// It means the whole cube is on/inside the 3D model.
// In this case, 6 tetrahedra are generated for the whole cube.

include <shared.scad>;

tetrahedron_table = [
    [ 12 + 0, 12 + 4, 12 + 7, 12 + 6 ],
    [ 12 + 0, 12 + 3, 12 + 7, 12 + 6 ],
    [ 12 + 0, 12 + 4, 12 + 5, 12 + 6 ],
    [ 12 + 0, 12 + 1, 12 + 5, 12 + 6 ],
    [ 12 + 0, 12 + 3, 12 + 2, 12 + 6 ],
    [ 12 + 0, 12 + 1, 12 + 2, 12 + 6 ],
];

create_tetrahedra(tetrahedron_table);

// Just to be able to visualize the created tetrahedra with respect to the cube.
// Draw cube last.
// https://stackoverflow.com/a/76245980/3405291
color([ 1, 1, 1, 0.2 ]) cube(size, center = false);