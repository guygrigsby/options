use std::fmt;
// START OMIT
#[derive(Debug)]
struct Struct {
    a: u8,
    b: u32,
    c: bool,
}

fn main() {
    let s = Struct {a: 7, b: 3333, c: true, }; // HL
    println!("{}", s);
}
// END OMIT
impl fmt::Display for Struct {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "({}, {}, {})", self.a, self.b, self.c)
    }
}
