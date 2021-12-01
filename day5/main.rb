require "digest"
require "pry"

def main
  input = "abc"
  #input = "abbhdwsy"

  final_one = ""
  tmp_one = ""
  final_two = []
  int = 0

  until final_one.size == 8 && final_two.size == 8
    hash = Digest::MD5.hexdigest(input + String(int))
    binding.pry if hash == "00000155f8105dff7f56ee10fa9b9abd"
    puts "hash: #{hash}" if hash.start_with?("00000")

    final_one = tmp_one if tmp_one.size == 8 && final_one.size != 8

    if hash.start_with?("00000")
      tmp_one += hash[5]

      index = hash[5]

      next unless index.is_a?(Integer)

      if index < hash.size
        final_two[index+1] = hash[index]
      end

    end
    puts "int: #{int}"

    int += 1
  end

  puts "part 1: #{final_one}"
  puts "part 2: #{final_two}"
end

main
