use std::collections::HashMap;
use std::time::{SystemTime, UNIX_EPOCH};

const VERSION: &str = "1.0.0";

struct ThreatIntel {
    ioc: String,
    risk_level: u8,
    tags: Vec<String>,
}

impl ThreatIntel {
    fn new(ioc: &str) -> Self {
        let risk = if ioc.contains(".") { 75 } else { 25 };
        ThreatIntel {
            ioc: ioc.to_string(),
            risk_level: risk,
            tags: vec!["malicious".to_string(), "c2".to_string()],
        }
    }
    
    fn analyze(&self) -> HashMap<String, String> {
        let mut result = HashMap::new();
        result.insert("ioc".to_string(), self.ioc.clone());
        result.insert("risk".to_string(), format!("{}/100", self.risk_level));
        result.insert("verdict".to_string(), 
            if self.risk_level > 70 { "HIGH RISK".to_string() } 
            else { "LOW RISK".to_string() });
        result
    }
}

fn get_timestamp() -> u64 {
    SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap()
        .as_secs()
}

fn main() {
    println!("Cyberbro Rust CLI v{}", VERSION);
    println!("Started at: {}", get_timestamp());
    
    let test_iocs = vec!["8.8.8.8", "evil.com", "malware.exe"];
    
    for ioc in test_iocs {
        let intel = ThreatIntel::new(ioc);
        let analysis = intel.analyze();
        println!("📊 Analysis for {}: {:?}", ioc, analysis);
    }
}
