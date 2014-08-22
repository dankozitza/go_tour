package main

import (
   "io"
   "os"
   "strings"
   "fmt"
)

type rot13Reader struct {
   r io.Reader
   letters []byte
}

func (r13 *rot13Reader) Read(p []byte) (n int, e error) {

   if len(p) == 0 {
      return 0, nil
   }

   n, e = r13.r.Read(p)

   //fmt.Printf("p before work: [%b]", p)


   l_map := make(map[byte]int, len(p)+1)

   for i, l := range r13.letters {
      fmt.Println("letter: [", string(l), "], index: [", i, "]")
      l_map[l] = i
   }

   fmt.Printf("letters length: [%i]\n", len(r13.letters))

   
   for i, l := range p {
      if (l_map[l] != 0) {

         offset := (l_map[l] + 13)

         if offset >= (len(r13.letters) - 1) {
            offset -= len(r13.letters) - 1
         }

         fmt.Println("offset: [", offset, "], i: [", i, "], l: [", string(l), "]")

         p[i] = r13.letters[offset]
      }
   }

   //n, e = r13.r.Read(p)

   //p[0] = 'a'

//   for l, i := range r.buff {
//      p[i] = l
//   }
   return
}

func main() {
   s := strings.NewReader(
      "Lbh penxrq gur pbqr!")
   r := rot13Reader{s, []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}}

   io.Copy(os.Stdout, &r)
}
