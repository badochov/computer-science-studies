#[cfg(test)]
mod tests {
    use crate::solution::Fibonacci;
    use ntest::timeout;

    #[test]
    #[timeout(100)]
    fn fibonacci_smoke_test() {
        assert_eq!(Fibonacci::fibonacci(3), 2);
    }

    #[test]
    #[timeout(100)]
    fn iterator_smoke_test() {
        let mut fib = Fibonacci::new();
        assert_eq!(fib.nth(3), Some(2));
    }

    #[test]
    #[timeout(1000)]
    fn iterator_and_fibonacci_equal() {
        let mut fib = Fibonacci::new();
        for i in 0..10 {
            match fib.next() {
                None => assert!(false),
                Some(it_val) => assert_eq!(it_val as u8, Fibonacci::fibonacci(i)),
            }
        }
    }

    #[test]
    #[timeout(1000)]
    fn check_iterator_ends() {
        let fib = Fibonacci::new();
        let mut prev: u128 = 0;
        for i in fib {
            assert!(i >= prev);
            prev = i;
        }

        assert!(prev > 1000000000000000);
    }
}
