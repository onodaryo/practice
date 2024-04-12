import sys

def is_leap_year(year):
    if year % 4 == 0 and year % 100 != 0:
        return True
    else:
        return False

if __name__ == "__main__":
    if len(sys.argv) > 1:
        is_leap_year(sys.argv[1])