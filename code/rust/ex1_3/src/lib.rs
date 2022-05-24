fn two_largest(a: i32, b: i32, _c: i32) -> [i32; 2] {
    if a > b { [a, b] } else { [b, a] }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn should_return_first_two_if_in_descending_order() {
        assert_eq!([3, 2], two_largest(3, 2, 1));
    }

    #[test]
    fn should_return_first_two_if_in_ascending_order() {
        assert_eq!([3, 2], two_largest(2, 3, 1));
    }
}
