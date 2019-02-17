    a, b = gets.split.map(&:to_i)
    puts (b % a).zero? ? a + b : b - a