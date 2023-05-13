module labeled_cube(corner, size) {
    edges = [
        [size/2, 0, 0],
        [size, size/2, 0],
        [size/2, size, 0],
        [0, size/2, 0],
        [size/2, 0, size],
        [size, size/2, size],
        [size/2, size, size[2]],
        [0, size/2, size],
        [0, 0, size/2],
        [size, 0, size/2],
        [size, size, size/2],
        [0, size, size/2],
    ];

    corners = [
        [0, 0, 0],
        [size, 0, 0],
        [size, size, 0],
        [0, size, 0],
        [0, 0, size],
        [size, 0, size],
        [size, size, size],
        [0, size, size],
    ];

  translate(corner) {
    color([1, 1, 1, 0.5]) cube(size, center=false); // create a transparent cube with the given size, centered at the given corner coordinates

    // label the corner points
    translate([0, 0, -5]) rotate([90, 0, 0]) text("C12", size=5);
    translate([size[0], 0, -5]) rotate([90, 0, 0]) text("C13", size=5);
    translate([size[0], size[1], -5]) rotate([90, 0, 0]) text("C14", size=5);
    translate([0, size[1], -5]) rotate([90, 0, 0]) text("C15", size=5);
    translate([0, 0, size[2]]) rotate([90, 0, 0]) text("C16", size=5);
    translate([size[0], 0, size[2]]) rotate([90, 0, 0]) text("C17", size=5);
    translate([size[0], size[1], size[2]]) rotate([90, 0, 0]) text("C18", size=5);
    translate([0, size[1], size[2]]) rotate([90, 0, 0]) text("C19", size=5);

    // label the edges
    translate([size[0]/2, 0, -5]) rotate([90, 0, 0]) text("E0", size=5);
    translate([size[0], size[1]/2, -5]) rotate([90, 0, 90]) text("E1", size=5);
    translate([size[0]/2+5, size[1], -5]) rotate([90, 0, 180]) text("E2", size=5);
    translate([0, size[1]/2, -5]) rotate([90, 0, -90]) text("E3", size=5);
    translate([size[0]/2, 0, size[2]+2]) rotate([90, 0, 0]) text("E4", size=5);
    translate([size[0], size[1]/2-3, size[2]+2]) rotate([90, 0, 90]) text("E5", size=5);
    translate([size[0]/2, size[1], size[2]+2]) rotate([90, 0, 180]) text("E6", size=5);
    translate([0, size[1]/2+3, size[2]+2]) rotate([90, 0, -90]) text("E7", size=5);
    translate([-1, 0, size[2]/2]) rotate([0, -90, 90]) text("E8", size=5);
    translate([size[0]+5, 0, size[2]/2]) rotate([90, -90, 0]) text("E9", size=5);
    translate([size[0]+1, size[1], size[2]/2]) rotate([180, -90, 90]) text("E10", size=5);
    translate([-1, size[1], size[2]/2]) rotate([0, -90, 0]) text("E11", size=5);
  }
}
