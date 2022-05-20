use float_eq::assert_float_eq;

fn square(x: f64) -> f64 {
    x*x
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn it_squares() {
        assert_float_eq!(4.0, square(2.0), abs <= 0.000_1);
    }
}
