package oneforall

import (
	"bytes"
	"fmt"
	"os/exec"
)

//接口鉴权等

//调度oneforall

func Oneforall(domain string) (string, error) {
	command := exec.Command("python3", "oneforall.py", "--target", domain, "run")
	command.Stderr = &bytes.Buffer{}
	//执行命令，直到命令结束
	err := command.Run()
	if err != nil {
		//打印程序中的错误以及命令行标准错误中的输出
		fmt.Println(err)
		fmt.Println(command.Stderr.(*bytes.Buffer).String())

	}
	//打印命令行的标准输出
	resultPath := domain + ".cvs"

	return resultPath, nil

}
