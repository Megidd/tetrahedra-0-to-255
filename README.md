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

![Screenshot_20230516_140235](https://github.com/Megidd/tetrahedron-table/assets/17475482/f6c9a066-64d8-4487-ac46-ea8bb22fceee)

## Case 2

![Screenshot_20230516_140256](https://github.com/Megidd/tetrahedron-table/assets/17475482/421181e7-177a-4ce3-822a-f8f0dc98e637)

## Case 3

![Screenshot_20230516_140318](https://github.com/Megidd/tetrahedron-table/assets/17475482/61b5bb12-403b-4e99-826e-af3ef25a7296)

## Case 4

![Screenshot_20230516_140344](https://github.com/Megidd/tetrahedron-table/assets/17475482/f4fb8af1-221b-414d-9f74-ca69d7d094a2)

## Case 5

![Screenshot_20230516_140617](https://github.com/Megidd/tetrahedron-table/assets/17475482/dd93b90d-0459-474f-980a-1dea795e8e17)

## Case 255

Case 255 result is below. All cube corners have zero/negative values.

![Screenshot_20230513_190339](https://github.com/Megidd/tetrahedron-table/assets/17475482/343201fa-76a4-4b83-bfbc-fb6323ac352a)
