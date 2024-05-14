package main

func contador(x int) {
    for i := 0; i < x; i++ {
        fnmt.Println(i)
        time.Sleep(time.Second)
    }
}
func main () { //thread1
  go contador (10) //thread2
  go contador (10) //thread3
  contador (10) 
}

//Processo -> Alocar um bloco de memoria
//Thread ->Acessar o bloco de memoria
// thread1 e thread2 -> acessam o mesmo bloco de memoria
// Race Condition - > Condição de corrida
// Evitar a thread com condição de corrida e criar a comunicação entres as threas com CHANNELS

// GO ROUTINE -> Canal -> GO ROUTINE 2

}
func main () { //thread1
  canal := make(chan string)

  go func () {
    canal <- "opa" //encia no canal o qual comunica via com a func main
    }()

  msg := <- canal
  fmt.Println(msg)
  go contador (10) //thread2
  go contador (10) //thread3
  contador (10) 
}


//#######################################
package main
import "fmt"
func worker (workerID int, data chan int) {
    for x := range data {
      fmt.Println("Worker %d received %d\n, workerID, x)
                  time.Sleep(timeSecoond)
    }
}
func main () { //thread1
  canal := make(chan int)
  go worker(1, canal) //entender direito/ criado a Thread2(GoRouTINE)
  for i := 0; i < 10; i++ {
    canal <- i
  }
}
//o Canal é zerado quando é lido o dado do canal
