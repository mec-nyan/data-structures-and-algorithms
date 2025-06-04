#!/usr/bin/env python3
"""Check if a series of parenthesis is balanced."""

from data_structures import Stack


def is_balanced(expr: str) -> bool:
    '''Check if parenthesis are balanced.'''

    s = Stack()

    for i in range(len(expr)):
        next = expr[i]
        if next == "(":
            s.push(next)
        elif next == ")":
            if s.is_empty():
                return False
            s.pop()

    return s.is_empty()


if __name__ == "__main__":
    s1 = "((A + B) * (C - D) / X) * A"
    s2 = "((()(()())"

    print(f"{s1}\n\t - balanced? {is_balanced(s1)}")
    print(f"{s2}\n\t - balanced? {is_balanced(s2)}")
