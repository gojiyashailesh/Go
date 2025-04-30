fn main(){
    let x = 32;//by defualt immutable
    println!("The value of the  X is : {x}");
    let mut y = 20;
    println!("{y}");
    y = 30;
    println!("{y}");

    //declare the constant which may not changed the leter on

    const ALWAYS_IN_CAPITAL_LETTER:u32= 60*60;
    println!("{}",ALWAYS_IN_CAPITAL_LETTER);
}