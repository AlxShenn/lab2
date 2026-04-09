import sys

def is_valid_local_part(local: str) -> bool:
    # Length check
    if not (6 <= len(local) <= 30):
        return False
    # Cannot start or end with a dot
    if local[0] == '.' or local[-1] == '.':
        return False
    # No consecutive dots
    if '..' in local:
        return False
    # Allowed characters: a-z, 0-9, ., +
    return all(c.isalnum() or c in '.+' for c in local) and local.islower()

def normalize(email: str) -> str:
    try:
        at_pos = email.index('@')
    except ValueError:
        return ""
    if at_pos == 0 or at_pos == len(email) - 1:
        return ""

    local = email[:at_pos]
    domain = email[at_pos:]

    if not is_valid_local_part(local):
        return ""

    new_local = []
    for c in local:
        if c == '+':
            break
        if c != '.':
            new_local.append(c)
    return ''.join(new_local) + domain

def main():
    print("1 - input from console\n2 - input from file")
    try:
        choice = int(input())
    except ValueError:
        print("Invalid choice")
        return

    emails = set()

    if choice == 1:
        n = int(input("Enter number of emails: "))
        for _ in range(n):
            email = input().strip()
            norm = normalize(email)
            if norm:
                emails.add(norm)
    elif choice == 2:
        filename = input("Enter file name: ")
        try:
            with open(filename, 'r') as f:
                for line in f:
                    for email in line.split():
                        norm = normalize(email)
                        if norm:
                            emails.add(norm)
        except IOError:
            print("File open error")
            return
    else:
        print("Invalid choice")
        return

    print("Number of unique emails:", len(emails))

if __name__ == "__main__":
    main()