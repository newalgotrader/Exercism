pub fn is_armstrong_number(num: u32) -> bool {
    let mut sum = 0;
    let mut result = num;

    let digits = (num as f32).log10().floor() as u32 + 1;

    while result > 0 {
        sum += (result % 10).pow(digits);
        result /= 10;
    }

    sum == num
}
