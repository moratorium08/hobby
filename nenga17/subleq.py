import sys
mem = []
pc = 0

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
    mem[addr] = val

def mainloop(src):
    while pc + 2 < len(mem):
        r1, r2, addr = mem[pc], mem[pc + 1], mem[pc + 2]
        val =  get_val(r2) - get_val(r1)
        set_val(r1, val)
        if val == 0:
            pc = addr
        else:
            pc += 3

def main():
    while True:
        tmp = ""
        while True:
            x = sys.stdin.read(1)
            if x in ' \n\t':
                tmp += x
                mem.append(int(tmp, 16))
                tmp = ""
    mainloop(mem)
