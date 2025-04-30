fn main() {
    let mut x = 5;   // make it mutable
    x = x + 1;       // change the same x
    println!("{}", x);
    {
        let x = x*2;
        println!("{}",x);
    }
    println!("{}",x);

    let spaces = "   ";
    let spaces = spaces.len();
    println!("{}",spaces);
}
