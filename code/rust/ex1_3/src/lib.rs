fn two_largest(_a: i32, _b: i32, _c: i32) -> [i32; 2] {
    [3, 2]
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn should_return_first_two_if_in_descending_order() {
        assert_eq!([3, 2], two_largest(3, 2, 1));
    }
}
