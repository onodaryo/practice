import function14 as f14


def check_error(test_input, expected_result):
    if f14.is_leap_year(test_input) == expected_result:
        print("input:"+str(test_input)+" is correct.")
    else:
        print("input:"+str(test_input)+" is not correct.")


if __name__ == "__main__":
    check_error(404, True)
    check_error(400, True)
    check_error(17, False)
    check_error(100, False)
