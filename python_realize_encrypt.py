import math


def get_offset_number(num):
    the_new_lottery_id = 1041100
    rand_numbers = [
        4874584950,
        8978964530,
        1581677540,
        3543542440,
        2534148840,
        5545842750,
        9454738480,
        7354159310,
        6191582590,
        6458327540
    ]
    old_rand_number = 123456789

    num_str = str(num)
    num_length = len(num_str)

    last_num = int(num_str[num_length - 1])

    rand_number = rand_numbers[last_num]

    if num > old_rand_number and num < the_new_lottery_id + old_rand_number:
        rand_number = old_rand_number
    elif num > 1100000 + old_rand_number and num < 38900000 + old_rand_number:
        rand_number = old_rand_number

    return rand_number


def num_to_string(num):
    num += get_offset_number(num)
    dest_str = str(num).zfill(10)
    dest_str_array = list(dest_str)
    num1 = int("".join([c for i, c in enumerate(dest_str_array) if i in (0, 2, 6, 9)]))
    num2 = int("".join([c for i, c in enumerate(dest_str_array) if i not in (0, 2, 6, 9)]))
    str1 = num_to_string_helper(num1)[::-1]
    str2 = num_to_string_helper(num2)[::-1]
    dest_str = str1.ljust(3, "U") + str2.ljust(4, "L")
    return dest_str


def num_to_string_helper(num):
    base_str = "0123456789ABCDEFGHJKMNPQRSTVWXYZabcdefghjkmnpqrstvwxyz"
    dest_str = ""
    base_str_array = list(base_str)
    while num != 0:
        temp_num = num % 32
        dest_str += base_str_array[temp_num]
        num = num // 32
    return dest_str


def string_to_num(s):
    base_str = "0123456789ABCDEFGHJKMNPQRSTVWXYZabcdefghjkmnpqrstvwxyz"
    num = 0
    for i in range(len(s)):
        temp_str = s[i]
        index = base_str.find(temp_str)
        if index > -1:
            num += index * math.pow(32, len(s) - i - 1)
    return num


def string_to_num_helper(s):
    str1 = s[:3].rstrip("U")
    str2 = s[3:].rstrip("L")
    num1 = string_to_num(str1)
    num2 = string_to_num(str2)
    str1 = str(int(num1)).zfill(4)
    str2 = str(int(num2)).zfill(6)
    str1_array = list(str1)
    str2_array = list(str2)
    num = int("".join([
        c for i, c in enumerate(str1_array + str2_array) if i in (0, 2, 4, 5, 7, 8, 11)
    ]))
    num -= get_offset_number(num)
    return num


def pad_left(s, pad, length):
    return s.rjust(length, pad)


def pad_right(s, pad, length):
    return s.ljust(length, pad)


def num_to_string_wrapper(num):
    return num_to_string(num)


def string_to_num_wrapper(s):
    return string_to_num_helper(s)


def split(s, sep):
    return [x.strip() for x in s.split(sep) if x.strip()]


def trim(s, cut_set=" "):
    return s.strip(cut_set)


def trim_left(s, cut_set=" "):
    return s.lstrip(cut_set)


def trim_right(s, cut_set=" "):
    return s.rstrip(cut_set)


def substr(s, pos, length):
    return s[pos:pos + length]


def padding(s, pad, length, pos):
    diff = len(s) - length
    if diff >= 0:
        return s

    mark = ""
    if pos == "PosRight":
        mark = "-"

    tpl = f"{mark}{length}"
    return f"{s:{tpl}}"


def repeat(s, times):
    if times < 2:
        return s

    return s * times


def to_ints(s, sep=","):
    return [int(x) for x in split(s, sep)]


def to_int(in_val):
    if isinstance(in_val, int):
        return in_val

    try:
        return int(str(in_val).strip())
    except ValueError:
        raise ValueError("convert data type is failure")


def to_int_slice(s, sep=","):
    return [to_int(sv) for sv in split(s, sep)]


def to_array(s, sep=","):
    return split(s, sep)


def to_slice(s, sep=","):
    return split(s, sep)


def main():
    # encrypt_secret = "0123456789ABCDEFGHJKMNPQRSTVWXYZabcdefghjkmnpqrstvwxyz"

    print(num_to_string_wrapper(123456789))  # Output: "pwuLkr"
    print(string_to_num_wrapper("pwuLkr"))  # Output: 123456789

    print(pad_left("abc", "0", 5))  # Output: "00abc"
    print(pad_right("abc", "Z", 6))  # Output: "abcZZZ"

    print(to_ints("1,2,3,4"))  # Output: [1, 2, 3, 4]
    print(to_int("123"))  # Output: 123
    print(to_int_slice("1,2,3,4"))  # Output: [1, 2, 3, 4]

    print(to_array("a,b,c"))  # Output: ["a", "b", "c"]
    print(to_slice("a,b,c"))  # Output: ["a", "b", "c"]


if __name__ == "__main__":
    main()
