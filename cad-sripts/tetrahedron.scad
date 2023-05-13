module create_tetrahedra(tetrahedron_count, tetrahedron_indices) {
    for (i = [0:tetrahedron_count-1]) {
        vertex_indices = tetrahedron_indices[i];

        tetrahedron_vertices = [
            edges_and_corners[vertex_indices[0]],
            edges_and_corners[vertex_indices[1]],
            edges_and_corners[vertex_indices[2]],
            edges_and_corners[vertex_indices[3]]
        ];

        tetrahedron(tetrahedron_vertices, tetrahedron_colors[i]);
    }
}

module tetrahedron(tetrahedron_vertices, tetrahedron_color) {
    color(tetrahedron_color) polyhedron(points=tetrahedron_vertices, faces=[[0, 1, 2], [0, 2, 3], [0, 3, 1], [1, 2, 3]]);
}