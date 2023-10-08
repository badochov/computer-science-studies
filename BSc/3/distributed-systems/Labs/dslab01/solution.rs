use std::convert::TryInto;

pub struct Fibonacci {
    // Add here any fields you need.
    curr: usize,
    prev: u128,
    prevprev: u128,
}

impl Fibonacci {
    /// Create new `Fibonacci`.
    pub fn new() -> Fibonacci {
        return Fibonacci {
            curr: 0,
            prev: 1,
            prevprev: 0
        }
    }

    /// Calculate the n-th Fibonacci number.
    ///
    /// This shall not change the state of the iterator.
    /// The calculations shall wrap around at the boundary of u8.
    /// The calculations might be slow (recursive calculations are acceptable).
    pub fn fibonacci(n: usize) -> u8 {
        if n < 2 {
            return n as u8;
        }
        else {
            return Fibonacci::fibonacci(n - 1).wrapping_add(Fibonacci::fibonacci(n - 2));
        }
        
    }
}

impl Iterator for Fibonacci {
    type Item = u128;

    /// Calculate the next Fibonacci number.
    ///
    /// The first call to `next()` shall return the 0th Fibonacci number (i.e., `0`).
    /// The calculations shall not overflow and shall not wrap around. If the result
    /// doesn't fit u128, the sequence shall end (the iterator shall return `None`).
    /// The calculations shall be fast (recursive calculations are **un**acceptable).
    fn next(&mut self) -> Option<Self::Item> {
        self.curr += 1;
        if self.curr < 3 {
            return Some((self.curr - 1).try_into().unwrap());
        }
        else{
            let res = self.prev.checked_add(self.prevprev);
            self.prevprev = self.prev;
            match res {
                None => {},
                Some(el) => self.prev = el,
            }
            return res;
        }
            
    }
}
