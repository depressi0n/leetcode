package question

import "strings"

func simplifyPath(path string) string {
	return simplifyPathCore1(path)
}

// 基本思想：使用'/'分割字符串，类似于栈的思想处理每个分隔好的串
func simplifyPathCore1(path string) string {
	res := make([]string, 0, len(path))
	fragments := strings.Split(path, "/")
	// '/'都被去掉了
	for _, fragment := range fragments {
		switch fragment {
		case ".":
			// 不用管
		case "..":
			//删除前一节，如果删除完了，则加上根
			if len(res)!=0{
				res=res[:len(res)-1]
			}
		case "":
			// 跳过去
			// pass
		default:
			res=append(res,fragment)
		}
	}
	return "/"+strings.Join(res,"/")
}
