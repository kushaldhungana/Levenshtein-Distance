package main
import "fmt"
import "unicode/utf8"

// This is the implementation of the levenstein algorithm to calculate the distance between two strings
func levenshtein(token1 string, token2 string) int{

// Checks if the two strings are equal, if it is, return with 0.
  if token1==token2 {
    fmt.Println("They are equal")
    return 0
  }
  tk1 := [] rune(token1)  //Conversion to rune slices.
  tk2 := [] rune(token2)
  lentk1 := utf8.RuneCountInString(token1) // Rune count of the string which gives the exact number of characters present.
  lentk2 := utf8.RuneCountInString(token2)

  //Initiating the column
  column := make([] int,lentk1+1)
  for b:=1; b<=lentk1 ; b++{
    column[b]=b
  }

  for a:=1; a<=lentk2; a++{
    column[0]= a
    last:= a-1
    for b:=1; b<=lentk1; b++{
      prev := column[b]
      var incr int
      if tk1[b-1] != tk2[a-1]{
        incr = 1
      }
      column[b] = min(column[b]+1,column[b-1]+1, last+incr)
      last = prev
    }
  }
  return column[lentk1]
}

// Implementation of the levenstein extended version
func levenshteinExtend(token1 string, token2 string, maxdist int) int{

  if token1==token2 {
    fmt.Println("They are equal")
    return 0
  }
  tk1 := [] rune(token1)
  tk2 := [] rune(token2)
  lentk1 := utf8.RuneCountInString(token1)
  lentk2 := utf8.RuneCountInString(token2)
  //Checks if the difference of length between two strings is greater than max distance, if it is return maxdis+1
  // This condition is not enought to check this because we might have difference of length
  // between two string to be less than 2 but might have different characters so we should again check it in last.
  if lentk1-lentk2 > maxdist || lentk2-lentk1 > maxdist{
    return maxdist+1
  }
  //Init column
  column := make([] int,lentk1+1)
  for b:=1; b<=lentk1 ; b++{
    column[b]=b
  }

  for a:=1; a<=lentk2; a++{
    column[0]= a
    last:= a-1
    for b:=1; b<=lentk1; b++{
      prev := column[b]
      var incr int
      if tk1[b-1] != tk2[a-1]{
        incr = 1
      }
      column[b] = min(column[b]+1,column[b-1]+1, last+incr)
      last = prev
    }
  }
  //Here again checking if the distance is greater than maxdist.
  if column[lentk1] > maxdist{
    return maxdist+1
  }
  return column[lentk1]
}

// Function to return the minimum value between x, y and z because we need the minimum distance.
func min(x,y,z int) int{
  if x<y{
    if x<z{
    return x
    }
  }else{
    if y<z{
      return y
    }
  }
  return z
}

  func main(){
    test1 := levenshtein("Haus","Maus")
    test2 := levenshtein("Haus","Mausi")
    test3 := levenshtein("Haus","Häuser")
    test4 := levenshtein("Kartoffelsalat","Runkelrüben")

    test5 := levenshteinExtend("Haus","Maus",2)
    test6 := levenshteinExtend("Haus","Mausi",2)
    test7 := levenshteinExtend("Haus","Häuser",2)
    test8 := levenshteinExtend("Kartoffelsalat","Runkelrüben",2)
    fmt.Printf("The Levenshtein distance for test1 is %v \n",test1)
    fmt.Printf("The Levenshtein distance for test2 is %v \n",test2)
    fmt.Printf("The Levenshtein distance for test3 is %v \n",test3)
    fmt.Printf("The Levenshtein distance for test4 is %v \n \n",test4)
    fmt.Printf("The Levenshtein extended distance for test1 is %v \n",test5)
    fmt.Printf("The Levenstein extended distance for test2 is %v \n",test6)
    fmt.Printf("The Levenshtein extended distance for test3 is %v \n",test7)
    fmt.Printf("The Levenshtein extended distance for test4 is %v \n",test8)
  }
