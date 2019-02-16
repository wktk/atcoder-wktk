gets
m = gets.split.map(&:to_i)

m.uniq!
ms = m.min(m.size / 2)
m.reject! { |l| l != m.min && ms.any? { |s| (l % s).zero? } }

while m.size > 1
  max0, max1 = m.max(2)
  times = (max0 - max1) / m.min
  times = 1 if times.zero?
  m[m.index(max0)] = max0 % m.min
  m.uniq!
  m.reject!(&:zero?)
end

puts m.first.to_i
