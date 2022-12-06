#### Keybord

   程序运行时，每当安按下键盘时，播放一个MP3(可以进行随机播放 math.rand   rand.rand)  修改MP3的名字

```
s :="1234567890zxcvbnmasdfghjkl;qtyuiop"
for i:=0;i <len(s);i++{
	if event.Rune == s[i]{
	
	Mp3.Open()
	Mp3.Play()
	}
	if event.Rune =="F1"{
	 
	}
}

func Open(){
	a:=rand.rand(63)
  os.Open(./....a)
}


```

