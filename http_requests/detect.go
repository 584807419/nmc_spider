package http_requests

// import (
// 	"bufio"

// 	"golang.org/x/net/html/charset"
// 	"golang.org/x/text/encoding"
// 	"golang.org/x/text/encoding/unicode"
// )

// // 处理获取的数据
// func determiEncoding(r *bufio.Reader) encoding.Encoding { //Encoding编码是一种字符集编码，可以在 UTF-8 和 UTF-8 之间进行转换
// 	//获取数据,Peek返回输入流的下n个字节
// 	bytes, err := r.Peek(1024)
// 	if err != nil {
// 		return unicode.UTF8
// 	}
// 	//调用DEtermineEncoding函数，确定编码通过检查最多前 1024 个字节的内容和声明的内容类型来确定 HTML 文档的编码。
// 	e, _, _ := charset.DetermineEncoding(bytes, "")
// 	return e
// }
