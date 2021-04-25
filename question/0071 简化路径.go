package question

import "strings"

func simplifyPath1(path string) string {
	ret := []string{}
	for _, v := range strings.Split(path, "/") {
		switch v {
		case "": //这种情况出现在连续两个//
			break
		case ".": //出现/./
			break
		case "..": //出现/../
			if l := len(ret); l > 0 {
				ret = ret[:l-1]
			}
		default: //其他情况
			ret = append(ret, v)
		}
	}
	return "/" + strings.Join(ret, "/")
}

//最后只能是分割/做，因为没考虑到...的情况
func simplifyPath(path string) string {
	var res []byte
	for i := 0; i < len(path); i++ { //肯定以/开始，不作处理
		if path[i] == '/' {
			//找到下一个不是/的地方
			j := i
			for j < len(path) && path[j] == '/' {
				j++
			}
			i = j - 1 //因为等一下还要经历i++
			if len(res) != 0 && res[len(res)-1] == '/' {
				continue
			}
			res = append(res, '/')
		} else if path[i] == '.' {
			//看下一个是不是.，如果是的话就要回退上一级目录，需要查看上级目录是否存在
			if i+1 < len(path) && path[i+1] == '.' { //退到上一级目录
				j := len(res) - 2
				if j < 0 { //顶级目录，不用退
					continue
				}
				for j >= 0 && path[j] != '/' {
					j--
				}
				res = res[:j] //截取到/之前，因为下次遇到/是可以加上的
			}
		} else { //找到下一个不是字母的地方
			j := i
			for j < len(path) && path[j] != '.' && path[j] != '/' {
				res = append(res, path[j])
				j++
			}
			i = j - 1
		}
	}
	if len(res) == 0 {
		return "/"
	}
	if len(res) != 1 && res[len(res)-1] == '/' {
		res = res[:len(res)-1]
	}
	return string(res)
}
