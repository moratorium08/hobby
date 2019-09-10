import sys
import bcrypt


def main():
    if len(sys.argv) < 3:
        return 2

    return 0 if bcrypt.checkpw(
            sys.argv[1].encode('utf-8'),
            sys.argv[2].encode('utf-8')) else 1


sys.exit(main())
