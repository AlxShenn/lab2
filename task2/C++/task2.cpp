//g++ task2.cpp -o task2.exe
//./task2.exe
#include <iostream>
#include <fstream>
#include <string>
#include <set>

using namespace std;

// функция для "нормализации" email
string normalize(string email) {
    int atPos = email.find('@');
    string name = email.substr(0, atPos);
    string domain = email.substr(atPos);

    string newName = "";
    for (int i = 0; i < name.size(); i++) {
        if (name[i] == '+') {
            break; // игнорируем всё после +
        }
        if (name[i] != '.') {
            newName += name[i];
        }
    }

    return newName + domain;
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
        for (int i = 0; i < n; i++) {
            cin >> email;
            emails.insert(normalize(email));
        }
    }
    else if (choice == 2) {
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
            emails.insert(normalize(email));
        }

        file.close();
    }
    else {
        cout << "Invalid choice\n";
        return 1;
    }

    cout << "Number of unique emails: " << emails.size() << endl;

    return 0;
}