import bcrypt

loop = 100

target = b"$2b$10$/LG4sfSIyNmjOB2cijyD7u20sWgy4zQ0idjbz1kmfZNaeroPEMAla"
password = b"uaoefaiweoafjaoweifjaoweifjaoweuhgapewguahwep"


def main():
    print(loop)
    for i in range(loop):
        if not bcrypt.checkpw(password, target):
            print("fail")
            return


main()
