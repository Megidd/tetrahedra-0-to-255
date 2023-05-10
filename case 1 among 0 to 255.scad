// Define the size of the cube
size = 50;

// Define the transparency of the tetrahedra
alpha = 0.5;

// Define the colors of the tetrahedra
colors = [
    [1, 0, 0, alpha],  // red
    [0, 1, 0, alpha],  // green
    [0, 0, 1, alpha],  // blue
    [1, 1, 0, alpha],  // yellow
    [1, 0, 1, alpha],  // magenta
    [0, 1, 1, alpha]   // cyan
];

// Define the vertices of the cube
cube_vertices = [
    [0, 0, 0],
    [size, 0, 0],
    [size, size, 0],
    [0, size, 0],
    [0, 0, size],
    [size, 0, size],
    [0, size, size],
    [size, size, size],
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
    
    echo("Tetrahedron ", i, ": ", vertex_indices);
    
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
color([1, 1, 1, 0]) cube(size);