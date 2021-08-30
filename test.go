package main

import (
	mruby "webimizer.dev/go_mruby"
)

func main() {
	mruby.Mruby_load_from_string(`# The Greeter class
	class Greeter
	  def initialize(name)
		@name = name.capitalize
	  end
	
	  def salute
		puts "Hello #{@name}!"
	  end
	end
	
	# Create a new object
	g = Greeter.new("world")
	
	# Output "Hello World!"
	g.salute`)
}
