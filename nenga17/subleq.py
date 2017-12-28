import sys
mem = []
pc = 0
src = []

def get_val(addr):
    # input
    if addr == -1:
        return ord(sys.stdin.read(1))
    if len(mem) <= addr:
        return 0
    return mem[addr]

def set_val(addr, val):
    if addr == -1:
        return sys.stdout.write(chr(val))
    while len(mem) <= addr:
        mem.append(0)
    mem[addr] = val & 0xff

def mainloop(src):
    while pc < len(src):
        r1, r2, addr = src[pc]
        val =  get_val(r2) - get_val(r1)
        set_val(r1, val)
        if val == 0:
            pc = addr
        else:
            pc += 1

def main():
    ls = input().split(" ")
    while len(ls) == 3:
        src.append([int(x, 16) for x in ls])
        ls = input().split(" ")
    mainloop(src)
