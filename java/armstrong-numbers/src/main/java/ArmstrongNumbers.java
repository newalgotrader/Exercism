class ArmstrongNumbers {

  boolean isArmstrongNumber(int numberToCheck) {
    int sum = 0;
    int result = numberToCheck;
    int digits = numberOfDigits(numberToCheck);

    while (result > 0) {
      sum += power(result % 10, digits);
      result /= 10;
    }

    return sum == numberToCheck;
  }

  private int numberOfDigits(int num) {
    int digits = 0;

    while(num > 0){
      num /= 10;
      ++digits;
    }

    return digits;
  }

  private int power(int base, int power) {
    int result = 1;

    while(power > 0) {
      result *= base;
      --power;
    }

    return result;
  }
}
