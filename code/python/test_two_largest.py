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

if __name__ == '__main__':
    unittest.main()
