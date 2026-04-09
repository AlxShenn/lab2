//g++ task2.cpp -o task2.exe
//./task2.exe
#include <iostream>
#include <fstream>
#include <string>
#include <set>

using namespace std;

bool isValidLocalPart(const string& local) {
    // Длина от 6 до 30 символов
    if (local.length() < 6 || local.length() > 30)
        return false;

    // Не может начинаться или заканчиваться точкой
    if (local.front() == '.' || local.back() == '.')
        return false;

    // Запрещены несколько точек подряд
    if (local.find("..") != string::npos)
        return false;

    // Проверка каждого символа
    for (char c : local) {
        bool isAllowed = (c >= 'a' && c <= 'z') ||
                         (c >= '0' && c <= '9') ||
                         c == '.' || c == '+';
        if (!isAllowed)
            return false;
    }
    return true;
}

string normalize(const string& email) {
    size_t atPos = email.find('@');
    if (atPos == string::npos || atPos == 0 || atPos == email.length() - 1)
        return "";

    string local = email.substr(0, atPos);
    string domain = email.substr(atPos);

    if (!isValidLocalPart(local))
        return "";

    string newLocal = "";
    for (char c : local) {
        if (c == '+')
            break;
        if (c != '.')
            newLocal += c;
    }

    return newLocal + domain;
}

int main() {
    set<string> emails;
    int choice;

    cout << "1 - input from console\n2 - input from file\n";
    cin >> choice;

    if (choice == 1) {
        int n;
        cout << "Enter number of emails: ";
        cin >> n;

        string email;
        for (int i = 0; i < n; ++i) {
            cin >> email;
            string norm = normalize(email);
            if (!norm.empty())
                emails.insert(norm);
        }
    } else if (choice == 2) {
        string filename;
        cout << "Enter file name: ";
        cin >> filename;

        ifstream file(filename);
        if (!file.is_open()) {
            cout << "File open error\n";
            return 1;
        }

        string email;
        while (file >> email) {
            string norm = normalize(email);
            if (!norm.empty())
                emails.insert(norm);
        }
        file.close();
    } else {
        cout << "Invalid choice\n";
        return 1;
    }

    cout << "Number of unique emails: " << emails.size() << endl;
    return 0;
}