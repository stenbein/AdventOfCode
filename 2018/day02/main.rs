//use std::env; //env variable functions, split paths, args, etc
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;
use std::path::Path;



fn lines_from_file<P>(filename: P) -> Vec<String>
where
    P: AsRef<Path>,
{
    let file = File::open(filename).expect("no such file");
    let buf = BufReader::new(file);
    buf.lines()
        .map(|l| l.expect("Could not parse line"))
        .collect()
}


fn count_in_str(c: char, s: &str) -> i32 {
    let mut i: i32 = 0;
    for ch in s.chars() {
        if ch == c {
            i += 1
        }
    }
    i
} 


fn is_threes(s: &str) -> bool {
    for c in s.chars() {
        if count_in_str(c, &s) == 3 {
            return true;
        }
    }
    return false;
}

fn is_twos(s: &str) -> bool {
    for c in s.chars() {
        if count_in_str(c, &s) == 2 {
            return true;
        }
    }
    return false;
}

//assume both strings same length
fn hamming(s1: &str, s2: &str) -> i32 {

    let s1_vec:Vec<char> = s1.chars().collect();
    let s2_vec:Vec<char> = s2.chars().collect();

    let mut dist: i32 = 0;

    for i in 0..s1_vec.len() {
        if s1_vec[i] != s2_vec[i] {
            dist += 1;
        }
    }
    dist
}

//final answer in part two is the common characters between two strings
fn in_common(s1: &String, s2: &String) -> String {

    let mut out_vec = vec![];
    let s1_vec:Vec<char> = s1.chars().collect();
    let s2_vec:Vec<char> = s2.chars().collect();

    for i in 0..s1_vec.len() {
        if s1_vec[i] == s2_vec[i] {
            out_vec.push(s1_vec[i]);
        }
    }

    let s: String = out_vec.into_iter().collect();
    s

}



fn main() {

    // itter all the inputs
    // indicate if they have double or triple letters and keep track of the counts
    // multiple those counts
    
    //first I need to find out how to get file input
        //copy from day01
            //forgot semicolon
    //itter each string
        //how to itter characters in rust
        //function syntax in rust

        //fight the borrow checker, is it really a fight? what did I mean with foo(x)
        //I didn't mean to give x away, I was letting foo see x, borrow it

    //worked

    //part2 find two strings with only a single difference
        //this is hamming distance of 1, return characters that are same
        //how to index string char
            //gonna try with an iterator and do a walk through the input line by line n^2
            //how to get input saved into array
                //arrays in rust are static, you mean vec
                    //yes
                //found method to try on stack overflow, for got semicolon twice and missed an import
                    //lots and lots of syntax errors, lots of issues with borrowing the vec in two places

        //didn't work, no value returned
            //printline the hamming shows everything is either 25 or 26 distance, obviously incorrect
            //assume I itterated the characters wrong
                // try a different method to first change the strings into vectors, iter that
                //worked

    //hmm, whats the difference between str and String in Rust?


    let filename = "input";

    let lines = lines_from_file(filename);

    //part 1
    let mut cnt_twos: i32 = 0;
    let mut cnt_threes: i32 = 0;

    for box_id in &lines {
        
        if is_twos(&box_id) {
            cnt_twos += 1
        }
        if is_threes(&box_id) {
            cnt_threes += 1
        }

    }

    let chk_sum: i32 = cnt_twos * cnt_threes;
    println!("Part 1 result is {}", chk_sum);

    //part 2
    let mut ind: usize = 0;

    'outer: for box_id1 in &lines {
        ind += 1;
        for box_id2 in &lines[ind..] {
            if hamming(&box_id1, &box_id2) == 1 {
                let result: str = in_common(&box_id1, &box_id2);
                println!("Part 2 result is {}", result);
                break 'outer
            }
        }

    }

    //println!("Part 2 result is {}", chk_sum);

}
