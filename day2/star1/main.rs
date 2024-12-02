use std::fs::File;
use std::path::Path;
use std::io::{self, BufRead};

fn parse_data(filename: String) -> Vec<Vec<i32>> {
    let path = Path::new(&filename);
    let file = match File::open(path) {
        Ok(file) => file,
        Err(e) => {
            eprintln!("{}", e);
            return vec![];
        }
    };

    let reader = io::BufReader::new(file);
    let mut stuff: Vec<Vec<i32>> = Vec::new();

    for lines in reader.lines() {
        match lines {
            Ok(content) => {
                let parts: Vec<&str> = content.split_whitespace().collect();
                let mut cont: Vec<i32> = Vec::new();
                for part in parts {
                    let s: i32 = part.parse().unwrap();
                    cont.push(s);
                }

                stuff.push(cont);
            }
            Err(e) => {
                eprintln!("{}", e);
            }
        };
    }

    return stuff;
}

fn num_of_safe(data: &Vec<Vec<i32>>) -> i32 {
    let safe: i32 = 0;
    for subarr in data {
        for in
    }
    return safe;
}

fn main() {
    // matrix
    let data = parse_data(String::from("../testdata.txt"));
    num_of_safe(&data);
    println!("{:?}", data);
}