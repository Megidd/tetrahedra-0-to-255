# Background

There is a source code that generates surface triangles. The isosurface is generated for the iso-value of `0`. The source code uses a table for `2^8=256` possible *inside/outside*, i.e. *negative/positive*, combinations of `8` scalar values at `8` cube corners. The table returns an array. Every `3` consecutive array items would correspond to a triangle. The array items could be from `0` to `11`, pointing to the `12` edges a cube has. Probably this table comes from a published paper in the field of mathematics or computer science:

https://github.com/deadsy/sdfx/blob/2d4e9502ec6fe898e8774020882cb8150f16a6a6/render/march3.go#L360

# Objective

I'm trying to adapt the above *marching cubes* source code, and its *tables*, to generate tetrahedra throughout the volume of a 3D model. The code would extract tetrahedra elements with all the non-positive, i.e. `<=0`, values. Non-positive means the 3D space *on* and *inside* the isosurface of the `0` value.

# Question

For some reason, I cannot find any publication for extracting a tetrahedral mesh *on* and *inside* the isosurface from a three-dimensional discrete scalar field. Maybe I'm not looking at the right places. Am I missing something? Or do I have to come up with the *tables* myself? It looks like a daunting task to me.

# Reference

https://github.com/deadsy/sdfx/pull/68#issuecomment-1447714965

Let's do it...

# Example

## Case 0

Case 0 is trivial. No cube corner has zero/negative value. No tetrahedron is generated.

## Case 1

Case 1 result is below. A cube corner has zero/negative value.

![Screenshot_20230513_190216](https://github.com/Megidd/tetrahedron-table/assets/17475482/94b2acbf-4784-4c36-8e6e-35f8126ed004)

## Case 255

Case 255 result is below. All cube corners have zero/negative values.

![Screenshot_20230513_190339](https://github.com/Megidd/tetrahedron-table/assets/17475482/343201fa-76a4-4b83-bfbc-fb6323ac352a)
