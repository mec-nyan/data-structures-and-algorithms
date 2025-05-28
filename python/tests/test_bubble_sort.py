#!/usr/bin/env python3


import unittest
from algorithms.sort.bubble import bubble_sort


class TestBubbleSort(unittest.TestCase):
    def test_bubble_sort(self) -> None:
        nums = [7, 4, 9, 2, 1, 2]
        bubble_sort(nums)
        self.assertEqual([1, 2, 2, 4, 7, 9], nums)


if __name__ == "__main__":
    unittest.main()
