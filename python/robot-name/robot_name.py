import string
import random


class Robot(object):
    used_names = {}
    max_name_gen_attempts = 1000

    def __init__(self):
        self.name = self.generate_name()

    def reset(self):
        self.name = self.generate_name()

    @classmethod
    def generate_name(cls):
        for x in range(cls.max_name_gen_attempts):
            name = "{}{}".format(
                ''.join(random.choice(string.ascii_uppercase)
                        for x in range(2)),
                ''.join(random.choice(string.digits) for x in range(3))
            )

            if name in cls.used_names:
                continue

            cls.used_names[name] = True
            return name

        raise Exception(f"Cannot find an unused name after {cls.max_name_gen_attempts} attempts")
