#!/usr/bin/env python3
"""Are these braces balanced?"""

from data_structures import Stack


def is_balanced(expr: str) -> bool:
    """Check if the braces in the expression are balanced."""

    s = Stack()

    open = "{[("
    close = "}])"

    for i in range(len(expr)):
        next = expr[i]
        if next in open:
            s.push(next)
        elif next in close:
            if s.is_empty():
                return False
            if matches(s.peek(), next):
                s.pop()
            else:
                return False

    return s.is_empty()


def matches(open: str, close: str) -> bool:
    open_braces = "{[("
    close_braces = "}])"

    return open_braces.index(open) == close_braces.index(close)


if __name__ == "__main__":
    s1 = "[ A + {(B + C) * (D - A) + Z} * 2]"

    print(f"{s1}\n\t - balanced? {is_balanced(s1)}")
