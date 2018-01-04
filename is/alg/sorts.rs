use std::io::*;
use std::str::*;
use std::iter::*;

fn readline() -> String {
    let mut s = String::new();
    stdin().read_line(&mut s).unwrap();
    s
}

fn val<T: FromStr>() -> T {
    let buf = readline();
    buf.trim_right().parse::<T>().ok().unwrap()
}

fn get_vals<T: FromStr>() -> Vec<T> {
    let s = readline();
    s.trim_right()
        .split_whitespace()
        .map(|t| t.parse::<T>().ok().unwrap())
        .collect()
}

fn bubble_sort(vs: &mut Vec<i32>) {
    for i in 0..vs.len() {
        for j in (i + 1..vs.len()).rev() {
            if vs[j - 1] > vs[j] {
                let tmp = vs[j];
                vs[j] = vs[j - 1];
                vs[j - 1] = tmp;
            }
        }
    }
}


fn _quick_sort(vs: &mut Vec<i32>) {
    if vs.len() == 1 {
        return;
    }
}

fn quick_sort(vs: &mut Vec<i32>) {
    _quick_sort(vs);
}

fn main() {
    let N = val::<i32>();
    let mut vs = get_vals::<i32>();

    for (i, x) in vs.iter().enumerate() {
        //println!("{}, {}", i, x);
    }
    bubble_sort(&mut vs);
    for (i, x) in vs.iter().enumerate() {
        //println!("{}, {}", i, x);
    }
    println!("{}", N);
    println!("{:?}", vs);
}
