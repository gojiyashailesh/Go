use std::io;
use rand::Rng;


fn main(){
    println!("Guess the Number:");
    let secret_number = rand::thread_rng().gen_range(1..=100);
    println!("Secret Nubmer is the {}",secret_number);
    println!("Please Enter the Number:");
    let mut guess = String::new();
    io::stdin().read_line(&mut guess).expect("failed to read line");
    println!("Your guessed {} ",guess);
    println!("You guessed: {guess}");

    match guess.cmp(&secret_number) {
        Ordering::Less => println!("Too small!"),
        Ordering::Greater => println!("Too big!"),
        Ordering::Equal => println!("You win!"),
    }
}