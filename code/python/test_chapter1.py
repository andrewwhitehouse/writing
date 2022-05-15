import unittest
from chapter1 import circumference

class TestChapter1(unittest.TestCase):

    def test_circumference(self):
        self.assertEqual(circumference(10), 62.8318)

if __name__ == '__main__':
    unittest.main()
