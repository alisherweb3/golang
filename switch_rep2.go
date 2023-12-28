/* Sound Frequencies
The human ear can hear sounds with frequencies in the range of 20 Hz to 20,000 Hz.
This is called the Audible frequency.
Anything below 20 Hz is called Infrasound, while anything above 20,000 Hz is called Ultrasound.

You need to make a program that takes a sound frequency as input and outputs the corresponding category.

Sample Input:
1800

Sample Output:
Audible

In case the given input is negative, your program should output "Wrong Input".
*/

package main

import "fmt"

func main() {
  var f int
  fmt.Scanln(&f)
  //your code goes here
  switch {
    case f >= 20 && f <= 20000:
      fmt.Println("Audible")
    case f < 20:
      fmt.Println("Infrasound")
    case f > 20000:
      fmt.Println("Ultrasound")
    default:
      fmt.Println("Wrong Input")
  }
}
