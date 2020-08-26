# https://tinygo.org/usage/subcommands/
# https://tinygo.org/usage/important-options/

# tinygo gdb -target=bluepill main.go
# tinygo run -target=bluepill .
# tinygo run main.go
# tinygo run -target=bluepill main.go

# tinygo build -o main.bin -target=bluepill .
# st-flash --reset write main.bin 0x08000000


tinygo flash -target=bluepill .
