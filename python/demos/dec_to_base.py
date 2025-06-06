#!/usr/bin/env python3
"""Convert a decimal number to another base."""

from data_structures import Stack


class BaseError(Exception):
    def __init__(self, msg):
        self.msg = msg


def dec_to_base(dec: int, base: int) -> str:
    # Max base is 16 in this particular function.
    if base > 16:
        raise BaseError(f"base is too big: {base}")

    # Assume dec is positive.
    # Zero is 0 in every base.
    if dec == 0:
        return "0"

    digits = "0123456789ABCDEF"

    stack = Stack()

    while dec > 0:
        stack.push(dec % base)
        dec //= base

    bin_digits: list[str] = []

    while not stack.is_empty():
        bin_digits.append(digits[stack.pop()])

    return "".join(bin_digits)


if __name__ == "__main__":
    print("{:6}  {:6}  {:6}  {:6}".format("dec", "bin", "oct", "hex"))
    for i in range(32):
        dec = i
        bin = dec_to_base(dec, 2)
        oct = dec_to_base(dec, 8)
        hex = dec_to_base(dec, 16)
        print(f"{dec:<6}  {bin:<6}  0o{oct:<4}  0x{hex:<4}")
