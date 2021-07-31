# Calculate length of string by Iterative and Recursive

        package main

        import "fmt"

        // Iterative length calculation O(n)
        // Как бы смысл в наличии этой функции не сильно понятен, 
        // так как в ней используем встроенную функцию len(),
        // но для сравнения с рекурсивной функцией решил ее написать
        func IterativeStrLen(inputString string) int {
            inputStrLen := 0
            for i := 0; i < len(inputString); i++ {
                inputStrLen += 1
            }
            return inputStrLen
        }

        // Recursive length calculation O(n)
        func RecursiveStrLen(inputString string) int {
            if inputString == "" {
                return 0
            }
            return 1 + RecursiveStrLen(inputString[1:])
        }

        func main() {
            inputStr := "LekanProgramming"

            fmt.Println(IterativeStrLen(inputStr))
            fmt.Println(RecursiveStrLen(inputStr))
        }
        
