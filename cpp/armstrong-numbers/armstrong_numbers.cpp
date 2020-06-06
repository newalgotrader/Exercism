#include "armstrong_numbers.h"

namespace armstrong_numbers {

static int number_of_digits(int);
static int power(int, int);

bool is_armstrong_number(int candidate) {
  int sum = 0;
  int result = candidate;
  int digits = number_of_digits(candidate);

  while (result > 0) {
    sum += power(result % 10, digits);
    result /= 10;
  }

  return sum == candidate;
}

static int number_of_digits(int number) {
  int digits = 0;

  while(number > 0) {
    number /= 10;
    ++digits;
  }

  return digits;
}

static int power(int base, int power) {
  int result = 1;

  while(power) {
    result *= base;
    --power;
  }

  return result;
}
}  // namespace armstrong_numbers
