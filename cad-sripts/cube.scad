module labeled_cube(corner, size) {
  translate(corner) {
    color([1, 1, 1, 0.5]) cube(size); // create a transparent cube with the given size, centered at the given corner coordinates

    // label the corner points
    translate([0, 0, -5]) rotate([90, 0, 0]) text("C0", size=5);
    translate([size[0], 0, -5]) rotate([90, 0, 0]) text("C1", size=5);
    translate([size[0], size[1], -5]) rotate([90, 0, 0]) text("C2", size=5);
    translate([0, size[1], -5]) rotate([90, 0, 0]) text("C3", size=5);
    translate([0, 0, size[2]]) rotate([90, 0, 0]) text("C4", size=5);
    translate([size[0], 0, size[2]]) rotate([90, 0, 0]) text("C5", size=5);
    translate([size[0], size[1], size[2]]) rotate([90, 0, 0]) text("C6", size=5);
    translate([0, size[1], size[2]]) rotate([90, 0, 0]) text("C7", size=5);
  }
}
