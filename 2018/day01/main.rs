//use std::env; //env variable functions, split paths, args, etc
use std::fs::File;
use std::io::prelude::*;
use std::io::BufReader;

use std::collections::HashSet;

fn main() {
    // itter all the inputs and sum
    
    //first I need to find out how to get file input
        //needed to look up how to make string literal
        //forgot the semi colon
        //forgot to import file module
        //worked, but now need line by line
        //what's the ? operator? equilivant to a match unpacking a result
        //import BufReader
        //no parse() in line?? is line string
        //cannot use ? in a function that returns '()' // 
    //second, how to convert string to int
        //unwrap the line, get the string
        //parse the string, unwrap that result
        //.unwrap() will panic if it incounters an error in main, fair enough

    //third, sum them together
        //missed uwrap again
        //ah, got it

    //part two, sum goes up and down in value. Break on the first duplicate
        //consider a hash table and b-tree
        //how do I hash in rust -> find HashSet
        //expected reference, found i32 -> ah, the first borrow check
        //help: consider borrowing here: `&sum` -> the famous rust borrowchecker, lets see what this does
        //cannot infer type for `T`
        //help: remove this 'mut' -> refering to mut on HashSet, which is probably telling me the reference is static
        //run, no output -> we're checking for contains but not actually adding anything to the set
        //forgot semicolon -> I can see this will be a common trend coming from Go and Python
        //expected i32, found &i32 help: consider removing the borrow: `sum'
            //now this perplexed me, is this because we borrow to check, but on insert we don't?
        //^^^^ cannot borrow mutably -> refering to the above removal of mut on HashSet, which now makes sense
        //no result, check that we didn't mess up anyting with the input by printline the sums, works
        //check that .contains works with literal inputs
            //how to input int literals in rust, 9999i32
        //above worked, so lets try a test input
        //test inpout counting down from 5 and back up again from 0, duplicates on 1, worked
            //check understanding of the problem via samples
            //Note that your device might need to repeat its list of frequency changes many times before a duplicate frequency is found
                //ah, while loop then
                //loop instead of while true
                //break outter loop -> loop labels
                    //value moved here in previous iteration of loop
                        //If I understand whats going on so far. I can move the variable declaration
                        // for the file handle inside the loop this should solve the scope issue
                            //runs -> and is the correct answer

                            //sweet baby jane, this feels like someone took the "explicit is better than implicit"
                            //and constructed a club from it to hit people. It's also kinda fun. How do I clean up this mess...


/*    let filename = "input";

    let f = File::open(filename).expect("file not found");

    let mut sum: i32 = 0;

    for line in BufReader::new(f).lines() {
        let change_str = line.unwrap();
        let delta = change_str.parse::<i32>();
        sum += delta.unwrap();
        //println!("{}", delta.unwrap());
    }
*/
    //println!("Part 1 result is {}", sum);

    //println!("With text:\n{}", contents);

    let mut sum: i32 = 0;
    let mut sums: HashSet<i32> = HashSet::new();

    'repeat: loop {

        let filename = "input".to_string();
        let f = File::open(filename).expect("file not found");

        for line in BufReader::new(f).lines() {

            if sums.contains(&sum) {
                println!("Part 2 result is {}", sum);
                break 'repeat;
            }
            sums.insert(sum);

            let change_str = line.unwrap();
            let delta = change_str.parse::<i32>();
            sum += delta.unwrap();
        }
    }
}
