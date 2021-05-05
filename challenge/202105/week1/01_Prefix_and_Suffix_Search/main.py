# Prefix and Suffix Search
from typing import List


class WordFilter:

    def __init__(self, words: List[str]):
        self.db = dict()

        for i, word in enumerate(words):
            for start in range(len(word), -1, -1):
                for end in range(len(word) + 1):
                    substr = f'{word[start:]}#{word[:end]}'
                    self.db[substr] = i
        
    def f(self, prefix: str, suffix: str) -> int:
        target = f'{suffix}#{prefix}'

        if target in self.db:
            return self.db[target]

        return -1
        

if __name__ == '__main__':
    words = ["apple"]
    prefix, suffix = "a", "e"
    obj = WordFilter(words)
    result = obj.f(prefix, suffix)
    assert result == 0
