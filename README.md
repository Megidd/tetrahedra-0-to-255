# Background

There is a source code that generates surface triangles. The isosurface is generated for the iso-value of `0`. The source code uses a table for `2^8=256` possible *inside/outside*, i.e. *negative/positive*, combinations of `8` scalar values at `8` cube corners. The table returns an array. Every `3` consecutive array items would correspond to a triangle. The array items could be from `0` to `11`, pointing to the `12` edges a cube has. Probably this table comes from a published paper in the field of mathematics or computer science:

https://github.com/deadsy/sdfx/blob/2d4e9502ec6fe898e8774020882cb8150f16a6a6/render/march3.go#L360

# Objective

I'm trying to adapt the above *marching cubes* source code, and its *tables*, to generate tetrahedra throughout the volume of a 3D model. The code would extract tetrahedra elements with all the non-positive, i.e. `<=0`, values. Non-positive means the 3D space *on* and *inside* the isosurface of the `0` value.

# Question

For some reason, I cannot find any publication for extracting a tetrahedral mesh *on* and *inside* the isosurface from a three-dimensional discrete scalar field. Maybe I'm not looking at the right places. Am I missing something? Or do I have to come up with the *tables* myself? It looks like a daunting task to me.
