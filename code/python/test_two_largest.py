import unittest
from two_largest import *

class TestTwoLargest(unittest.TestCase):

    def test_first_two(self):
        self.assertEqual([3,2], two_largest(3,2,1))

    def test_first_two_out_of_order(self):
        self.assertEqual([3,2], two_largest(2,3,1))

    def test_last_two_largest_in_order(self):
        self.assertEqual([6,5], two_largest(4,6,5))

    def test_last_two_out_of_order(self):
        self.assertEqual([9,8], two_largest(7,8,9))

    def test_first_and_last_in_order(self):
        self.assertEqual([102, 101], two_largest(102, 100, 101))

    def test_first_and_last_out_of_order(self):
        self.assertEqual([502, 501], two_largest(501, 500, 502))

    def test_all_equal(self):
        self.assertEqual([123, 123], two_largest(123, 123, 123))

    def test_negative_numbers(self):
        self.assertEqual([-10, -20], two_largest(-10, -20, -30))

    def test_decimals(self):
        self.assertEqual([2.001, 2.0003], two_largest(2.001, 2.0002, 2.0003))

if __name__ == '__main__':
    unittest.main()
