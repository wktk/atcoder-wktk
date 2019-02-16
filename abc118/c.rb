gets
m = gets.split.map(&:to_i)

m.uniq!

while m.size > 1
  if m.size > 3
    ms = m.min([m.size / 2, 100].min)
m.reject! { |l| l != m.min && ms.any? { |s| (l % s).zero? } }
  end
  break if m.size < 2
m[m.index(m.max)] = m.max % m.min
  m.reject!(&:zero?)
  m.uniq!
end

puts m.first.to_i
