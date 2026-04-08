//kotlinc task1.kt -include-runtime -d task1.jar
//java -jar task1.jar
import java.io.File

// Функция проверки и построения палиндрома
// Возвращает Pair (возможно ли, сам палиндром)
fun canFormPalindrome(s: String): Pair<Boolean, String> {
    // Массив для подсчёта символов (ASCII)
    val count = IntArray(256)
    
    // Считаем, сколько раз встречается каждый символ
    for (c in s) {
        count[c.code]++
    }
    
    // Проверяем, сколько символов встречается нечётное количество раз
    var oddCount = 0
    var oddChar = ' '
    
    for (i in 0 until 256) {
        if (count[i] % 2 == 1) {
            oddCount++
            oddChar = i.toChar()
        }
    }
    
    // Если нечётных символов больше одного - палиндром невозможен
    if (oddCount > 1) {
        return Pair(false, "")
    }
    
    // Строим половину палиндрома
    val half = StringBuilder()
    
    for (i in 0 until 256) {
        if (count[i] >= 2) {
            // Добавляем половину символов
            repeat(count[i] / 2) {
                half.append(i.toChar())
            }
        }
    }
    
    // Сортируем для красивого вывода
    val sortedHalf = half.toString().toCharArray().sorted().joinToString("")
    
    // Собираем палиндром: половина + центр + перевёрнутая половина
    val palindrome = StringBuilder()
    palindrome.append(sortedHalf)
    
    if (oddCount == 1) {
        palindrome.append(oddChar)
    }
    
    // Добавляем перевёрнутую половину
    palindrome.append(sortedHalf.reversed())
    
    return Pair(true, palindrome.toString())
}

fun main() {
    println("Select input method:")
    println("1 - Input from console")
    println("2 - Read from file")
    print("Your choice: ")
    
    val choice = readLine()?.toIntOrNull() ?: 0
    
    if (choice == 1) {
        // ========== ВВОД С КОНСОЛИ ==========
        println("\n=== Console Input ===")
        print("Enter a string: ")
        val input = readLine() ?: ""
        
        val (possible, result) = canFormPalindrome(input)
        if (possible) {
            println("Result: Yes")
            println("Palindrome: $result")
        } else {
            println("Result: No")
        }
    } else if (choice == 2) {
        // ========== ЧТЕНИЕ ИЗ ФАЙЛА ==========
        println("\n=== File Input ===")
        print("Enter filename (e.g., input.txt): ")
        val filename = readLine() ?: ""
        
        val file = File(filename)
        if (!file.exists()) {
            println("Error: Could not open file $filename")
            return
        }
        
        var lineNumber = 1
        
        file.forEachLine { line ->
            // Пропускаем пустые строки
            if (line.isNotBlank()) {
                println("\nLine $lineNumber: \"$line\"")
                
                val (possible, result) = canFormPalindrome(line)
                if (possible) {
                    println("  Result: Yes")
                    println("  Palindrome: $result")
                } else {
                    println("  Result: No")
                }
                
                lineNumber++
            }
        }
        
        println("\nTotal lines processed: ${lineNumber - 1}")
    } else {
        println("Invalid choice!")
    }
}