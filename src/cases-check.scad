include <shared.scad>;
include <tetrahedron-table.scad>
include <triangle-table.scad>

// Case number can be from 0 to 255.
// i = 254;

for (i = [0:255])
{

    triangle_table = mcTriangleTable[i];

    draw_triangles(triangle_table);

    tetrahedron_table = mcTetrahedronTable[i];

    for (l = [0:len(triangle_table) - 1])
    {
        matches = find_matches(triangle_table[l], tetrahedron_table);
        found = len(matches) > 0;
        if (!found)
        {
            echo("*** Case with un-used edge with a triangle on that edge.");
            echo("*** Case:", i, "edge:", triangle_table[l]);
            echo("*** trianggle table", triangle_table);
            echo("*** tetrahedron table", tetrahedron_table);
        }
    }

    for (k = [0:len(tetrahedron_table) - 1])
    {

        matches = find_matches(tetrahedron_table[k], triangle_table);

        found = len(matches) > 0;

        // Any cube corner can be used for tetrahedron creation.
        // But no cube edge can be used for tetrahedron creation, unless that edge has a triangle on it.
        // Below 12 means an edge.
        if (!found && tetrahedron_table[k] < 12)
        {
            echo("*** Case with wrong edge.");
            echo("*** Case:", i, "tethedron:", k, "edge:", tetrahedron_table[k]);
            echo("*** trianggle table", triangle_table);
            echo("*** tetrahedron table", tetrahedron_table);
        }
    }

    points = tetrahedra_points(tetrahedron_table = tetrahedron_table);

    for (j = [0:1:len(points) - 1])
    {
        jac = jacobian(points[j]);
        det = determinant(jac);
        if (det <= 0)
        {
            echo("*** Case with wrong points order.");
            echo("*** Case:", i);
            echo("*** Tetrahedron:", j, " Points: ", points[j]);
            echo("*** Jacobian matrix:\n", jac);
            echo("*** Determinant of Jacobian matrix:", det);
        }
    }

    union()
    {
        // No case will ever have more than this number of tetrahedra.
        // So, we are on the safe side.
        polyhedron(points = points[0], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[1], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[2], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[3], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[4], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[5], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[6], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[7], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[8], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[9], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[10], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[11], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[12], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[13], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[14], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[15], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
        polyhedron(points = points[16], faces = [ [ 0, 1, 2 ], [ 0, 2, 3 ], [ 0, 3, 1 ], [ 1, 2, 3 ] ]);
    }

    color([ 1, 1, 1, 0.2 ]) cube(size, center = false);
}