import unittest
from two_largest import *

class TestTwoLargest(unittest.TestCase):

    def test_two_largest(self):
        self.assertEqual([3,2], two_largest(3,2,1))

if __name__ == '__main__':
    unittest.main()
