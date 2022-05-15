import unittest
from chapter1 import circumference

class TestChapter1(unittest.TestCase):

    def test_circumference(self):
        self.assertAlmostEqual(circumference(10), 62.8318, places=3)

if __name__ == '__main__':
    unittest.main()
