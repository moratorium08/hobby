package main

import (
	"fmt"
	"os/exec"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

const (
	loop   = 100
	target = "$2b$10$/LG4sfSIyNmjOB2cijyD7u20sWgy4zQ0idjbz1kmfZNaeroPEMAla"
	pass   = "uaoefaiweoafjaoweifjaoweifjaoweuhgapewguahwep"
)

func BenchmarkGolang(b *testing.B) {
	passwd := []byte(target)
	b.ResetTimer()
	for i := 0; i < loop; i++ {
		if err := bcrypt.CompareHashAndPassword(passwd, []byte(pass)); err != nil {
			fmt.Printf("golang: fail: %v\n", err)
			break
		}
	}
}

func BenchmarkCallPython(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < loop; i++ {
		cmd := exec.Command("python", "run.py", pass, target)
		if err := cmd.Run(); err != nil {
			fmt.Println("python: fail")
			break
		}
	}
}
