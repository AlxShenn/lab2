//g++ task3.cpp -o task3.exe
//./task3.exe
#include <iostream>
#include <fstream>
#include <cmath>

using namespace std;

// проверка на полный квадрат
bool isPerfectSquare(int x) {
    int root = (int)sqrt(x);
    return root * root == x;
}

int main() {
    int choice;
    cout << "1 - input from console\n2 - input from file\n";
    cin >> choice;

    int count = 0;

    if (choice == 1) {
        int n;
        cout << "Enter N: ";
        cin >> n;

        int x;
        for (int i = 0; i < n; i++) {
            cin >> x;
            if (isPerfectSquare(x)) {
                count++;
            }
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

        int n, x;
        file >> n;

        for (int i = 0; i < n; i++) {
            file >> x;
            if (isPerfectSquare(x)) {
                count++;
            }
        }

        file.close();
    }
    else {
        cout << "Invalid choice\n";
        return 1;
    }

    cout << "Result: " << count << endl;

    return 0;
}