//g++ task1.cpp -o task1.exe
//./task1.exe
#include <iostream>
#include <string>
#include <vector>
#include <algorithm>
#include <fstream>

using namespace std;

// Функция проверки и построения палиндрома
// Возвращает true, если палиндром можно составить, и записывает результат в palindrome
bool canFormPalindrome(string s, string &palindrome) {
    // Массив для подсчёта символов (работает для ASCII символов)
    // 256 - чтобы хватило на все символы ASCII
    int count[256] = {0};
    
    // Считаем, сколько раз встречается каждый символ
    for (int i = 0; i < s.length(); i++) {
        count[s[i]]++;
    }
    
    // Проверяем, сколько символов встречается нечётное количество раз
    int oddCount = 0;
    char oddChar = ' ';
    
    for (int i = 0; i < 256; i++) {
        if (count[i] % 2 == 1) {
            oddCount++;
            oddChar = (char)i;
        }
    }
    
    // Если нечётных символов больше одного - палиндром невозможен
    if (oddCount > 1) {
        return false;
    }
    
    // Строим половину палиндрома
    string half = "";
    
    for (int i = 0; i < 256; i++) {
        if (count[i] >= 2) {
            // Добавляем половину символов
            for (int j = 0; j < count[i] / 2; j++) {
                half += (char)i;
            }
        }
    }
    
    // Сортируем для красивого вывода (опционально)
    sort(half.begin(), half.end());
    
    // Собираем палиндром: половина + центр + перевёрнутая половина
    palindrome = half;
    
    if (oddCount == 1) {
        palindrome += oddChar;
    }
    
    // Переворачиваем половину и добавляем
    reverse(half.begin(), half.end());
    palindrome += half;
    
    return true;
}

int main() {
    cout << "Select input method:" << endl;
    cout << "1 - Input from console" << endl;
    cout << "2 - Read from file" << endl;
    cout << "Your choice: ";
    
    int choice;
    cin >> choice;
    cin.ignore(); // Очищаем буфер после ввода числа
    
    if (choice == 1) {
        // ========== ВВОД С КОНСОЛИ ==========
        cout << "\n=== Console Input ===" << endl;
        
        string input;
        cout << "Enter a string: ";
        getline(cin, input);
        
        string result;
        if (canFormPalindrome(input, result)) {
            cout << "Result: Yes" << endl;
            cout << "Palindrome: " << result << endl;
        } else {
            cout << "Result: No" << endl;
        }
    } 
    else if (choice == 2) {
        // ========== ЧТЕНИЕ ИЗ ФАЙЛА ==========
        cout << "\n=== File Input ===" << endl;
        
        string filename;
        cout << "Enter filename (e.g., input.txt): ";
        cin >> filename;
        
        ifstream file(filename);
        
        if (!file.is_open()) {
            cout << "Error: Could not open file " << filename << endl;
            return 1;
        }
        
        string line;
        int lineNumber = 1;
        
        while (getline(file, line)) {
            // Пропускаем пустые строки
            if (line.empty()) {
                continue;
            }
            
            cout << "\nLine " << lineNumber << ": \"" << line << "\"" << endl;
            
            string result;
            if (canFormPalindrome(line, result)) {
                cout << "  Result: Yes" << endl;
                cout << "  Palindrome: " << result << endl;
            } else {
                cout << "  Result: No" << endl;
            }
            
            lineNumber++;
        }
        
        file.close();
        cout << "\nTotal lines processed: " << lineNumber - 1 << endl;
    }
    else {
        cout << "Invalid choice!" << endl;
    }
    
    return 0;
}