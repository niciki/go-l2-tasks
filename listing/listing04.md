# 4. Что выведет программа? Объяснить вывод программы.
```go
package main
 
func main() {
    ch := make(chan int)
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
        }
    }()
 
    for n := range ch {
        println(n)
    }
}
```
Будут выведены цифры от 0 до 9, а затем сообщение о `deadlock`. Это произойдет потому, что в горутине мы последовательно пишем в канал цифры,
а в цикле читаем их. После записи всех чисел в цикле будет ожидаться очередная запись числа в канал. Канал необходимо было закрыть в горутине, чтобы
цикл `for` считывания автоматически завершился. 
