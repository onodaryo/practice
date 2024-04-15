import sys

def is_leap_year(year):
    return year % 4 == 0 and (year % 100 != 0 or year % 400 == 0)

if __name__ == "__main__":
    if len(sys.argv) > 1:
        is_leap_year(sys.argv[1])
