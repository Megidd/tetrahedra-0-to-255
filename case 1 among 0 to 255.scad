module divide_cube_into_tetrahedra(cube_size, alpha=0.5, colors=[
    [1, 0, 0, alpha],  // red
    [0, 1, 0, alpha],  // green
    [0, 0, 1, alpha],  // blue
    [1, 1, 0, alpha],  // yellow
    [1, 0, 1, alpha],  // magenta
    [0, 1, 1, alpha]   // cyan
]) {
    // Define the vertices of the cube
    cube_vertices = [
        [0, 0, 0],
        [cube_size, 0, 0],
        [cube_size, cube_size, 0],
        [0, cube_size, 0],
        [0, 0, cube_size],
        [cube_size, 0, cube_size],
        [0, cube_size, cube_size],
        [cube_size, cube_size, cube_size],
    ];

    tetrahedron_indices = [
        [0, 4, 6, 7],
        [0, 3, 6, 7],
        [0, 4, 5, 7],
        [0, 1, 5, 7],
        [0, 3, 2, 7],
        [0, 1, 2, 7],
    ];

    // Define the tetrahedra by selecting vertices from the cube
    for (i = [0:5]) {
        // Define the indices of the vertices for the tetrahedron
        vertex_indices = tetrahedron_indices[i];

        // Define the vertices of the tetrahedron
        tetrahedron_vertices = [
            cube_vertices[vertex_indices[0]],
            cube_vertices[vertex_indices[1]],
            cube_vertices[vertex_indices[2]],
            cube_vertices[vertex_indices[3]]
        ];

        // Color the tetrahedron with the specified color
        color(colors[i]) polyhedron(points=tetrahedron_vertices, faces=[[0, 1, 2], [0, 2, 3], [0, 3, 1], [1, 2, 3]]);
    }

    // Draw the cube
    color([1, 1, 1, 0]) cube(cube_size);
}

// Example usage of the module
cube_size = 50;
alpha = 0.5;
colors = [
    [1, 0, 0, alpha],  // red
    [0, 1, 0, alpha],  // green
    [0, 0, 1, alpha],  // blue
    [1, 1, 0, alpha],  // yellow
    [1, 0, 1, alpha],  // magenta
    [0, 1, 1, alpha]   // cyan
];

divide_cube_into_tetrahedra(cube_size, alpha, colors);
