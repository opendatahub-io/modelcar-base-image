use std::env;
use std::os::unix::fs::symlink;
use std::process;
use std::thread;

fn main() {
    println!("Hello, world! And thanks Jason for the idea, Daniele and Roland");
    let args: Vec<String> = env::args().collect();
    println!("Invoked with and will check later is consistent with KServe deployment {}", args.join(" "));
    do_the_thing();
    thread::park();
    println!("Main terminated.")
}

fn do_the_thing() {
    let pid = process::id(); 
    let source = format!("/proc/{}/root/models", pid);
    let target = "/mnt/models";

    if let Err(e) = symlink(&source, target) {
        eprintln!("Failed to create symbolic link: {}", e);
        process::exit(1);
    }

    println!("Symbolic link created from '{}' to '{}'", source, target);
}
