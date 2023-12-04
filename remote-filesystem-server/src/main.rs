use std::path::PathBuf;
use std::time::UNIX_EPOCH;
use std::env::current_dir;
use std::fs;
use serde::{Serialize,Deserialize};

fn main() -> std::io::Result<()>{
    let scanned_dir = DirScan::new(current_dir()?);
    Ok(())
}

#[derive(Serialize, Deserialize)]
struct Info {
    name: String,
    size: u64,
    last_modification: String,
}

#[derive(Serialize, Deserialize)]
struct DirScan {
    path: String,
    dir: Vec<Info>,
    files: Vec<Info>,
}

impl DirScan {
    fn new(path: PathBuf) -> DirScan {
        let scanned_dir = DirScan {
            path: current_dir().unwrap().display().to_string(),
            dir: Vec::new(),
            files: Vec::new()
        };
        let path = current_dir().unwrap().display();
        let metadata = fs::metadata(path);
    }

    fn parent() -> DirScan {

    }
}
