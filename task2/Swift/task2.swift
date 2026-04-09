import Foundation

func isValidLocalPart(_ local: String) -> Bool {
    // Length check
    guard local.count >= 6 && local.count <= 30 else { return false }
    // Cannot start or end with a dot
    guard !local.hasPrefix(".") && !local.hasSuffix(".") else { return false }
    // No consecutive dots
    guard !local.contains("..") else { return false }
    // Allowed characters: a-z, 0-9, ., +
    let allowedSet = CharacterSet.lowercaseLetters
        .union(.decimalDigits)
        .union(CharacterSet(charactersIn: ".+"))
    return local.unicodeScalars.allSatisfy { allowedSet.contains($0) }
}

func normalize(_ email: String) -> String? {
    guard let atPos = email.firstIndex(of: "@") else { return nil }
    let local = String(email[..<atPos])
    let domain = String(email[atPos...])

    guard isValidLocalPart(local) else { return nil }

    var newLocal = ""
    for ch in local {
        if ch == "+" { break }
        if ch != "." { newLocal.append(ch) }
    }
    return newLocal + domain
}

func main() {
    print("1 - input from console\n2 - input from file")
    guard let choiceStr = readLine(), let choice = Int(choiceStr) else {
        print("Invalid input")
        return
    }

    var uniqueEmails = Set<String>()

    if choice == 1 {
        print("Enter number of emails:")
        guard let nStr = readLine(), let n = Int(nStr) else { return }

        for _ in 0..<n {
            guard let email = readLine() else { continue }
            if let norm = normalize(email) {
                uniqueEmails.insert(norm)
            }
        }
    } else if choice == 2 {
        print("Enter file name:")
        guard let filename = readLine() else { return }

        do {
            let content = try String(contentsOfFile: filename, encoding: .utf8)
            let words = content.split { $0.isWhitespace || $0.isNewline }.map(String.init)
            for email in words {
                if let norm = normalize(email) {
                    uniqueEmails.insert(norm)
                }
            }
        } catch {
            print("File open error")
            return
        }
    } else {
        print("Invalid choice")
        return
    }

    print("Number of unique emails:", uniqueEmails.count)
}

main()