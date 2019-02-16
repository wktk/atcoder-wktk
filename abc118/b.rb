n, m = gets.split

o = gets.split.tap(&:shift)

while l = gets
  o = o & l.split.tap(&:shift)
end

puts o.size
