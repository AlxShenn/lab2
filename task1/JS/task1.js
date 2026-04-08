//node task1.js
const readline = require('readline');
const fs = require('fs');

// Функция проверки и построения палиндрома
// Возвращает объект {possible: boolean, palindrome: string}
function canFormPalindrome(s) {
    // Массив для подсчёта символов (ASCII)
    const count = new Array(256).fill(0);
    
    // Считаем, сколько раз встречается каждый символ
    for (let i = 0; i < s.length; i++) {
        const code = s.charCodeAt(i);
        count[code]++;
    }
    
    // Проверяем, сколько символов встречается нечётное количество раз
    let oddCount = 0;
    let oddChar = ' ';
    
    for (let i = 0; i < 256; i++) {
        if (count[i] % 2 === 1) {
            oddCount++;
            oddChar = String.fromCharCode(i);
        }
    }
    
    // Если нечётных символов больше одного - палиндром невозможен
    if (oddCount > 1) {
        return { possible: false, palindrome: "" };
    }
    
    // Строим половину палиндрома
    let half = [];
    
    for (let i = 0; i < 256; i++) {
        if (count[i] >= 2) {
            // Добавляем половину символов
            for (let j = 0; j < Math.floor(count[i] / 2); j++) {
                half.push(String.fromCharCode(i));
            }
        }
    }
    
    // Сортируем для красивого вывода
    half.sort();
    
    // Собираем палиндром: половина + центр + перевёрнутая половина
    let palindrome = half.join('');
    
    if (oddCount === 1) {
        palindrome += oddChar;
    }
    
    // Добавляем перевёрнутую половину
    palindrome += half.slice().reverse().join('');
    
    return { possible: true, palindrome: palindrome };
}

// Настройка ввода с консоли
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout
});

function askQuestion(query) {
    return new Promise(resolve => rl.question(query, resolve));
}

async function main() {
    console.log("Select input method:");
    console.log("1 - Input from console");
    console.log("2 - Read from file");
    const choice = await askQuestion("Your choice: ");
    
    if (choice === "1") {
        // ========== ВВОД С КОНСОЛИ ==========
        console.log("\n=== Console Input ===");
        const input = await askQuestion("Enter a string: ");
        
        const { possible, palindrome } = canFormPalindrome(input);
        if (possible) {
            console.log("Result: Yes");
            console.log("Palindrome:", palindrome);
        } else {
            console.log("Result: No");
        }
    } 
    else if (choice === "2") {
        // ========== ЧТЕНИЕ ИЗ ФАЙЛА ==========
        console.log("\n=== File Input ===");
        const filename = await askQuestion("Enter filename (e.g., input.txt): ");
        
        if (!fs.existsSync(filename)) {
            console.log("Error: Could not open file", filename);
            rl.close();
            return;
        }
        
        const content = fs.readFileSync(filename, 'utf-8');
        const lines = content.split('\n');
        
        let lineNumber = 1;
        
        for (const line of lines) {
            const trimmedLine = line.trim();
            
            // Пропускаем пустые строки
            if (trimmedLine === "") {
                continue;
            }
            
            console.log(`\nLine ${lineNumber}: "${trimmedLine}"`);
            
            const { possible, palindrome } = canFormPalindrome(trimmedLine);
            if (possible) {
                console.log("  Result: Yes");
                console.log("  Palindrome:", palindrome);
            } else {
                console.log("  Result: No");
            }
            
            lineNumber++;
        }
        
        console.log(`\nTotal lines processed: ${lineNumber - 1}`);
    }
    else {
        console.log("Invalid choice!");
    }
    
    rl.close();
}

main();