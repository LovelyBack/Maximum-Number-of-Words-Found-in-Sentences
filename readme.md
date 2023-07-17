# Solving the problem on leet code

## Maximum Number of Words Found in Sentences with golang

A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

You are given an array of strings sentences, where each ```sentences[i]``` represents a single sentence.

Return the maximum number of words that appear in a single sentence.

### Example 

```Input : sentences = ["alice and bob love leetcode", "i think so too", "this is great thanks very much"]```
```Output : 6 ```

### how it works

You can change the array passed to ```mostWordsFound``` function in this code

- First, we called the function and passed the array to it,

```golang
func main(){
    sentence := []string{"x","y","z"}
    mostWordsFound(sentence)
}
```
- then we called the second function (```mostWordsFound```)
- now we should write the second function :
```golang
func mostWordsFound(sentence []string){
    //We define this variable so that the maximum value is saved in it
    max := 0

    // now we define a for loop on the array"sentence"
    for _, word := range sentence{
        //Here we use the split method to convert each sentence into several words and each word into an array member.
        // and then we  get the len() of the array
        //Here we saved the number of words of each sentence in the length variable
        lenght := len(strings.Split(word," "))
        
        //Then we check if length is greater than len and save it in len
        if lenght > max {
            max = lenght
        }
    }
    //then we print the max
    fmt.Println(max)
}
```