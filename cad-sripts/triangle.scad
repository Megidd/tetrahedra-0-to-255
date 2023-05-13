module draw_triangles(indices, vertices) {
  assert(len(indices) % 3 == 0, "indices must have a length that is a multiple of 3");

  for (i = [0:len(indices)-1:3]) {
    triangle(
      vertices[indices[i]],
      vertices[indices[i+1]],
      vertices[indices[i+2]]
    );
  }
}

module triangle(p1, p2, p3) {
  polyhedron(points=[p1, p2, p3], faces=[[0, 1, 2]]);
}
