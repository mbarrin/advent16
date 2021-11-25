ALPHABET = "abcdefghijklmnopqrstuvwxyz"

def main

  lines = File.readlines("input.txt", chomp: true)
  count = 0

  lines.each do |line|
    groups = /(.*)-(\d+)\[(.*)\]/.match(line)

    name, id, checksum = groups[1], groups[2], groups[3].chars.sort.join

    if checksum == make_checksum(name: name)
      count += Integer(id)
    end

    puts "part 2: #{id}" if decrypt(name: name, shift: id) == "northpole object storage"
  end

  puts "part 1: #{count}"
end

def decrypt(name:, shift:)
  decrypted = ""

  name.chars.each do |c|
    if c == "-"
      decrypted += " "
    else
      decrypted += ALPHABET[(ALPHABET.index(c) + (Integer(shift) % 26)) % 26]
    end
  end

  return decrypted
end

def make_checksum(name:)
  string = name.gsub("-","").split("").tally

  if string.size == 5
    return string.keys.sort.join
  end

  if string.size > 5
    tmp = []
    foo = string.group_by { |i| i[1] }.sort.reverse
    foo.each { |x| tmp << x[1].sort }
    bar = tmp.flatten.select{ |i| i.is_a? String}
    return bar[0..4].sort.join
  end
end

main
