"""Stack implementation in Python."""


class Stack:
    """Stack implements the stack ADT."""

    def __init__(self):
        self.items = []

    def push(self, item):
        self.items.append(item)

    def pop(self):
        if self.size():
            return self.items.pop()
        else:
            return None

    def peek(self):
        if self.size():
            return self.items[self.size() - 1]
        else:
            return None

    def is_empty(self):
        return self.size() == 0

    def size(self):
        return len(self.items)

    def __str__(self):
        output = "-" * 80
        output += "\nStack:"
        if self.size():
            for i in range(self.size()):
                output += f"\n\t{i + 1:2}:{self.items[i]}"
            output += f"\nSize: {self.size()}"
        else:
            output += " [] (emtpy)"

        return output


if __name__ == "__main__":
    s = Stack()
    print("Is empty:", s.is_empty())
    print(s)

    s.push("Neko")
    s.push(7)
    s.push("Inu")

    print(s)
    print("Peek:", s.peek())

    last = s.pop()
    print("Poped:", last)
    print(s, f"(size: {s.size()})")
