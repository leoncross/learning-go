package iteration

func Repeat(charater string, count int) string{
  var repeated string
  for i := 0; i < count; i++ {
    repeated += charater
  }

  return repeated
}
