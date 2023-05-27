size = 80;

edges = [
    [ size / 2, 0, 0 ],
    [ size, size / 2, 0 ],
    [ size / 2, size, 0 ],
    [ 0, size / 2, 0 ],
    [ size / 2, 0, size ],
    [ size, size / 2, size ],
    [ size / 2, size, size ],
    [ 0, size / 2, size ],
    [ 0, 0, size / 2 ],
    [ size, 0, size / 2 ],
    [ size, size, size / 2 ],
    [ 0, size, size / 2 ],
];

corners = [
    [ 0, 0, 0 ],
    [ size, 0, 0 ],
    [ size, size, 0 ],
    [ 0, size, 0 ],
    [ 0, 0, size ],
    [ size, 0, size ],
    [ size, size, size ],
    [ 0, size, size ],
];

edges_and_corners = concat(edges, corners);

// label the edges
// We are adjusting the labels by custom offsets and rotations.
translate(edges[0] + [ 0, 0, -5 ]) rotate([ 90, 0, 0 ]) text("0", size = 5);
translate(edges[1] + [ 0, -3, -5 ]) rotate([ 90, 0, 90 ]) text("1", size = 5);
translate(edges[2] + [ 5, 0, -5 ]) rotate([ 90, 0, 180 ]) text("2", size = 5);
translate(edges[3] + [ 0, 0, -5 ]) rotate([ 90, 0, -90 ]) text("3", size = 5);
translate(edges[4] + [ 0, 0, 2 ]) rotate([ 90, 0, 0 ]) text("4", size = 5);
translate(edges[5] + [ 0, -3, 2 ]) rotate([ 90, 0, 90 ]) text("5", size = 5);
translate(edges[6] + [ 5, 0, 2 ]) rotate([ 90, 0, 180 ]) text("6", size = 5);
translate(edges[7] + [ 0, 3, 2 ]) rotate([ 90, 0, -90 ]) text("7", size = 5);
translate(edges[8] + [ -10, 0, 0 ]) rotate([ 90, 0, 0 ]) text("8", size = 5);
translate(edges[9] + [ 1, 0, 0 ]) rotate([ 90, 0, 0 ]) text("9", size = 5);
translate(edges[10] + [ 14, 0, 0 ]) rotate([ 90, 0, 180 ]) text("10", size = 5);
translate(edges[11] + [ -2, 0, 0 ]) rotate([ 90, 0, 180 ]) text("11", size = 5);

// label the corner points
// We are adjusting the labels by custom offsets and rotations.
translate(corners[0] + [ 0, 0, -5 ]) rotate([ 90, 0, 0 ]) text("12", size = 5);
translate(corners[1] + [ 0, 0, -5 ]) rotate([ 90, 0, 0 ]) text("13", size = 5);
translate(corners[2] + [ 0, 0, -5 ]) rotate([ 90, 0, 180 ]) text("14", size = 5);
translate(corners[3] + [ 0, 0, -5 ]) rotate([ 90, 0, 180 ]) text("15", size = 5);
translate(corners[4]) rotate([ 90, 0, 0 ]) text("16", size = 5);
translate(corners[5]) rotate([ 90, 0, 0 ]) text("17", size = 5);
translate(corners[6]) rotate([ 90, 0, 180 ]) text("18", size = 5);
translate(corners[7]) rotate([ 90, 0, 180 ]) text("19", size = 5);

module draw_triangles(triangle_table)
{
    for (i = [0:3:len(triangle_table) - 1])
    {
        indices = [triangle_table[i], triangle_table[i+1], triangle_table[i+2]];

        p0 = edges[indices[0]];
        p1 = edges[indices[1]];
        p2 = edges[indices[2]];

        color([ 1, 1, 1, 0.4 ]) polyhedron(points = [ p0, p1, p2 ], faces = [[ 0, 1, 2 ]]);

        echo("Draw triangle ", i, "indices: ", indices);
    }
}

function tetrahedra_points(tetrahedron_table) =
    [for (i = [0:len(tetrahedron_table) -
                 1])[edges_and_corners[tetrahedron_table[i][0]], edges_and_corners[tetrahedron_table[i][1]],
                     edges_and_corners[tetrahedron_table[i][2]], edges_and_corners[tetrahedron_table[i][3]]]];