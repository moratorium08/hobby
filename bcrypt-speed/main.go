package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	/*
		cmd := exec.Command("python", "run.py", "hogehoge", target)
		if err := cmd.Run(); err != nil {
			fmt.Printf("fail: %v\n", err)
		} else {
			fmt.Println("ok")
		}

		passwd := []byte(target)
		if err := bcrypt.CompareHashAndPassword(passwd, []byte("hogehoge")); err != nil {
			fmt.Println("fail")
		} else {
			fmt.Println("ok")
		}
	*/
	loop := 100
	target := "$2b$10$/LG4sfSIyNmjOB2cijyD7u20sWgy4zQ0idjbz1kmfZNaeroPEMAla"
	pass := "uaoefaiweoafjaoweifjaoweifjaoweuhgapewguahwep"
	passwd := []byte(target)
	uo := []byte(pass)
	fmt.Println(loop)
	for i := 0; i < loop; i++ {
		if err := bcrypt.CompareHashAndPassword(passwd, uo); err != nil {
			fmt.Println("fail")
		}
	}

}
