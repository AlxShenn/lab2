#python task2.py
# функция для "нормализации" email
def normalize(email):
    at_pos = email.find('@')
    name = email[:at_pos]
    domain = email[at_pos:]

    new_name = ""
    for ch in name:
        if ch == '+':
            break  # игнорируем всё после +
        if ch != '.':
            new_name += ch

    return new_name + domain

emails = set()

print("1 - input from console\n2 - input from file")
choice = int(input())

if choice == 1:
    n = int(input("Enter number of emails: "))
    for _ in range(n):
        email = input().strip()
        emails.add(normalize(email))

elif choice == 2:
    filename = input("Enter file name: ").strip()

    try:
        # открываем файл с указанием кодировки
        with open(filename, "r", encoding="utf-8") as f:
            for line in f:
                email = line.strip()

                if email == "":
                    continue  # пропускаем пустые строки

                emails.add(normalize(email))

    except Exception as e:
        print("File open error:", e)
        exit()

else:
    print("Invalid choice")
    exit()

print("Number of unique emails:", len(emails))