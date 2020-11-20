package main

import "fmt"

func main(){
	//接口和结构体之间的联系和使用规范
	//接口：是一套标准,适合于定义共性，抽离并定义出一套标准
	//结构体：实体的描述和定义
	//接口:有廷议的共性进行判断时,往往使用的是接口
	anime1 :=NewJapenAnime()
	if anime1.Animenum > 13 {
		fmt.Println("集数太多没空看")
	}else {
		fmt.Println("下次一定看,集数为",anime1.Animenum,"集")
	}

	if anime1.Com != "京阿尼" {
		fmt.Println("我只看京阿尼的动漫")
	}else {
		fmt.Println("京阿尼赛高")
	}

	if anime1.IsSequel == true {
		fmt.Println("下一部在做了")
	}else {
		fmt.Println("没有了")
	}
}


//接口动漫都要满足这个标准
type Anime interface {
	Company() string //公司
	Num() int // 集数
	Animetype() string //动漫类型
}

//日本动漫
type JapenAnime struct {
	Com string //公司
	Atype string //动漫类型
	Animenum int //动漫集数
	Name string //动漫名字
	IsSequel  bool //是否有下一部
}

func NewJapenAnime() *JapenAnime{
	j :=&JapenAnime{
		Com:      "京阿尼",
		Atype:    "催泪",
		Animenum: 12,
		Name:     "玉子市场",
		IsSequel: true,
	}
	return j
}

func (j *JapenAnime) Company() string{
	return j.Company()
}
func (j *JapenAnime) Num() int{
	return j.Num()
}
func (j *JapenAnime) Animetype() string{
	return j.Animetype()
}