include <shared.scad>;

// Triangles come from here:
// https://github.com/deadsy/sdfx/blob/1a71e404e4b2aa00c59f53cffc219a9e83e62d85/render/march3.go#L360
triangle_table = [ 1, 8, 3, 9, 8, 1 ];
draw_triangles(indices = triangle_table, vertices = edges, triangle_color = [ 0, 1, 1, .4 ]);

tetrahedron_indices = [
    [ 1, 15, 3, 19 ],
    [ 1, 14, 15, 19 ],
    [ 1, 14, 19, 18 ],
    [ 1, 18, 19, 17 ],
    [ 1, 3, 9, 17 ],
    [ 17, 19, 3, 16 ],
    [ 9, 3, 8, 17 ],
    [ 8, 17, 3, 16 ],
];

create_tetrahedra(tetrahedron_indices);

// Just to be able to visualize the created tetrahedra with respect to the cube.
// Draw cube last.
// https://stackoverflow.com/a/76245980/3405291
color([ 1, 1, 1, 0.2 ]) cube(size, center = false);