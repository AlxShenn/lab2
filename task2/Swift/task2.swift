//swiftc task2.swift
//./task2.exe
import Foundation

// функция для "нормализации" email
func normalize(_ email: String) -> String? {
    // проверка на наличие @
    guard email.contains("@") else {
        return nil
    }

    let parts = email.split(separator: "@")
    
    // должно быть ровно 2 части
    if parts.count != 2 {
        return nil
    }

    let name = String(parts[0])
    let domain = String(parts[1])

    if name.isEmpty || domain.isEmpty {
        return nil
    }

    var newName = ""

    for ch in name {
        if ch == "+" {
            break // игнорируем всё после +
        }
        if ch != "." {
            newName.append(ch)
        }
    }

    return newName + "@" + domain
}

var emails = Set<String>()

print("1 - input from console\n2 - input from file")
if let choice = Int(readLine() ?? "") {

    if choice == 1 {
        print("Enter number of emails: ")
        if let n = Int(readLine() ?? "") {
            for _ in 0..<n {
                if let email = readLine() {
                    if let norm = normalize(email.trimmingCharacters(in: .whitespacesAndNewlines)) {
                        emails.insert(norm)
                    }
                }
            }
        }

    } else if choice == 2 {
        print("Enter file name: ")
        if let filename = readLine() {
            do {
                let content = try String(contentsOfFile: filename, encoding: .utf8)
                let lines = content.components(separatedBy: .newlines)

                for line in lines {
                    let email = line.trimmingCharacters(in: .whitespacesAndNewlines)
                    if email.isEmpty { continue }

                    if let norm = normalize(email) {
                        emails.insert(norm)
                    }
                }
            } catch {
                print("File open error")
                exit(1)
            }
        }

    } else {
        print("Invalid choice")
        exit(1)
    }

    print("Number of unique emails: \(emails.count)")
}