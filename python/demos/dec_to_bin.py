#!/usr/bin/env python3
"""Convert a decimal number to its binary representation."""

from data_structures import Stack


def to_binary(dec: int) -> str:
    # Assume dec is positive.
    if dec == 0:
        return "0"

    stack = Stack()

    while dec > 0:
        stack.push(dec % 2)
        dec //= 2

    bin_digits: list[str] = []

    while not stack.is_empty():
        bin_digits.append(str(stack.pop()))

    return "".join(bin_digits)


if __name__ == "__main__":
    print("decimal - binary")
    for i in range(16):
        print(f"{i:7} - {to_binary(i):>08}")
