use clap::Parser;
use colored::*;

#[derive(Parser)]
#[command(name = "cyberbro-cli")]
struct Cli {
    ioc: String,
    #[arg(short, long, default_value = "http://localhost:5000")]
    server: String,
}

fn main() {
    let cli = Cli::parse();
    println!("{} Analyzing: {}", "[*]".green(), cli.ioc.cyan());
    
    let client = reqwest::blocking::Client::new();
    if let Ok(resp) = client.post(&format!("{}/api/analyze", cli.server))
        .json(&serde_json::json!({"text": cli.ioc}))
        .send() {
        println!("{}", resp.text().unwrap());
    }
}
