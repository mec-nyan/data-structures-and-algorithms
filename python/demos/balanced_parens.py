#!/usr/bin/env python3
"""Check if a series of parenthesis is balanced."""

from data_structures import Stack


def is_balanced(s: str) -> bool:
    '''Check if parenthesis are balanced.'''

    stk = Stack()
    balanced = True

    i = 0
    while i < len(s) and balanced:
        if s[i] == "(":
            stk.push(s[i])
        elif s[i] == ")":
            if not stk.is_empty():
                stk.pop()
            else:
                balanced = False
        i += 1

    return stk.size() == 0


if __name__ == "__main__":
    s1 = "((()())())"
    s2 = "((()(()())"

    print(f"{s1} - balanced? {is_balanced(s1)}")
    print(f"{s2} - balanced? {is_balanced(s2)}")
