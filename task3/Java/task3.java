import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

public class task3 {

    // проверка на полный квадрат
    public static boolean isPerfectSquare(int x) {
        int root = (int)Math.sqrt(x);
        return root * root == x;
    }

    public static void main(String[] args) {
        Scanner sc = new Scanner(System.in);

        System.out.println("1 - input from console\n2 - input from file");
        int choice = sc.nextInt();

        int count = 0;

        if (choice == 1) {
            System.out.print("Enter N: ");
            int n = sc.nextInt();

            for (int i = 0; i < n; i++) {
                int x = sc.nextInt();
                if (isPerfectSquare(x)) {
                    count++;
                }
            }

        } else if (choice == 2) {
            System.out.print("Enter file name: ");
            String filename = sc.next();

            try {
                Scanner file = new Scanner(new File(filename));

                int n = file.nextInt();
                for (int i = 0; i < n; i++) {
                    int x = file.nextInt();
                    if (isPerfectSquare(x)) {
                        count++;
                    }
                }

                file.close();

            } catch (FileNotFoundException e) {
                System.out.println("File open error");
                return;
            }

        } else {
            System.out.println("Invalid choice");
            return;
        }

        System.out.println("Result: " + count);
    }
}