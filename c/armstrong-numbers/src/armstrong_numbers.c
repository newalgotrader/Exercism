#include "armstrong_numbers.h"
#include <math.h>

int is_armstrong_number(int candidate) {
  int sum = 0;
  int result = candidate;
  int digits = floor(log10(candidate)) + 1;

  while (result > 0) {
    sum += round(pow(result % 10, digits));
    result /= 10;
  }

  return sum == candidate;
}
