const expect = @import("std").testing.expect;

test "if statement" {
    const a = true;
    var value: u16 = 0;
    if (a) {
        value += 1;
    } else {
        value += 2;
    }
    try expect(value == 1);
}

test "if expression" {
    const a = true;
    var value: i16 = 0;
    value += if (a) 1 else 2;
    try expect(value == 1);
}
