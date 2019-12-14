package main
import "testing"

//This is the function to check the levenstein algorithm in all the test cases provided.
func TestTableLevenshtein(t *testing.T){
  var tests= []struct{
    input1 string
    input2 string
    expected int
  }{
    {"Haus","Maus",1},
    {"Haus","Mausi",2},
    {"Haus","Häuser",3},
    {"Kartoffelsalat","Runkelrüben",12},
  }

  for _, test := range tests{
      if output:= levenshtein(test.input1,test.input2); output!=test.expected{
        t.Error("Test Failed: {} expected, received: {}",test.expected, output)
      }
    }
}

// This is the function to check the levenstein Extended version in all the test cases provided.
func TestTableLevenshteinExtend(t *testing.T){
  var tests= []struct{
    input1 string
    input2 string
    input3 int
    expected int
  }{
    {"Haus","Maus",2,1},
    {"Haus","Mausi",2,2},
    {"Haus","Häuser",2,3},
    {"Kartoffelsalat","Runkelrüben",2,3},
  }

  for _, test := range tests{
      if output:= levenshteinExtend(test.input1,test.input2,test.input3); output!=test.expected{
        t.Error("Test Failed: {} expected, received: {}",test.expected, output)
      }
    }
}

// This is the function to check the benchmark of the Levenshtein function in all the test cases.
func BenchmarkLevenshtein(b *testing.B){

  var tests= [] struct{
    input1 string
    input2 string
    name string
  }{
    {"Haus","Maus","Firstcase"},
    {"Haus","Mausi","Secondcase"},
    {"Haus","Häuser","Thirdcase"},
    {"Kartoffelsalat","Runkelruben","Fourthcase"},
  }

  for _,test := range tests{
    b.Run(test.name, func(b *testing.B){
        for i:=0; i<b.N; i++{
          levenshtein(test.input1,test.input2)
        }
    })
  }
}

// This is the function to check the benchmark of the Levenshtein extended version in all the test cases.
func BenchmarkLevenshteinExtend(b *testing.B){

  var tests= [] struct{
    input1 string
    input2 string
    input3 int
    name string
  }{
    {"Haus","Maus",2,"FirstcaseExtended"},
    {"Haus","Mausi",2,"SecondcaseExtended"},
    {"Haus","Häuser",2,"ThirdcaseExtended"},
    {"Kartoffelsalat","Runkelruben",2,"FourthcaseExtended"},
  }

  for _,test := range tests{
    b.Run(test.name, func(b *testing.B){
        for i:=0; i<b.N; i++{
          levenshteinExtend(test.input1,test.input2,test.input3)
        }
    })
  }
}
