package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	SessionID string = "xxx"
	userid    string = "xxx"
)

func main() {
	getvideolist()
}

func getvideolist() {
	videolisturl := "https://www.djbhjg.net:4430/ucc?AppID=YRWA_JGGLPT&Password=YRWA_JGGLPT2557013/&PageName=MHWZ&Controlname=videolist&return=true&SessionID=" + SessionID
	data, _ := json.Marshal(struct {
		Pagesize int `json:"pagesize"`
		Pagenum  int `json:"pagenum"`
		Lv       int `json:"lv"`
		IsMust   int `json:"is_must"`
	}{
		Pagenum:  8,
		Pagesize: 1,
		Lv:       0,
		IsMust:   2,
	})

	request, _ := http.NewRequest("POST", videolisturl, bytes.NewBuffer(data))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36")
	request.Header.Set("Origin", "https://www.djbhjg.net:4430")
	request.Header.Set("Referer", "https://www.djbhjg.net:4430/jgglweb/index.html")

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var videolistresult struct {
		List []struct {
			Title     string `json:"title"`
			IsMust    int    `json:"is_must"`
			CoursesId string `json:"courses_id"`
		} `json:"list"`
	}
	json.Unmarshal(body, &videolistresult)
	for x, _ := range videolistresult.List {
		fmt.Println("课程名：", videolistresult.List[x].Title)
		if videolistresult.List[x].IsMust == 1 {
			//fmt.Println("必修")
		} else {
			//fmt.Println("选修")
		}
		coursesid := videolistresult.List[x].CoursesId
		记录点击(coursesid)
		time.Sleep(10 * time.Second)
		getcourselist(coursesid)
	}
	fmt.Println("End.")
}

// 啊咧咧，中文函数名
func 记录点击(coursesid string) {
	记录点击url := "https://www.djbhjg.net:4430/ucc?ucc?AppID=YRWA_JGGLPT&Password=YRWA_JGGLPT2557013/&Password=KF_YRWA_JGGLPT2557013&PageName=MHWZ&Controlname=click&return=true&SessionID=" + SessionID
	data, _ := json.Marshal(
		map[string]interface{}{
			"list": []map[string]string{
				{"courses_id": coursesid},
			},
		},
	)

	request, _ := http.NewRequest("POST", 记录点击url, bytes.NewBuffer(data))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36")
	request.Header.Set("Origin", "https://www.djbhjg.net:4430")
	request.Header.Set("Referer", "https://www.djbhjg.net:4430/jgglweb/index.html")

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	response, _ := client.Do(request)
	defer response.Body.Close()

	body, _ := io.ReadAll(response.Body)
	var 没想好变量1 struct {
		ReturnValue int    `json:"ReturnValue"`
		Msg         string `json:"msg"`
	}
	json.Unmarshal(body, &没想好变量1)
	if 没想好变量1.ReturnValue != 1 || 没想好变量1.Msg != "已记录" {
		//fmt.Println("第二步:", 没想好变量1.ReturnValue, 没想好变量1.Msg)
	}
	time.Sleep(10 * time.Second)
	未知步骤3(coursesid)
}

func 未知步骤3(coursesid string) {
	// 只管发包 不要问作用
	未知步骤3url := "https://www.djbhjg.net:4430/ucc?ucc?AppID=YRWA_JGGLPT&Password=YRWA_JGGLPT2557013/&PageName=MHWZ&Controlname=KCXQ&return=true&SessionID=" + SessionID
	data, _ := json.Marshal(struct {
		CoursesId string `json:"courses_id"`
	}{
		CoursesId: coursesid,
	},
	)
	request, _ := http.NewRequest("POST", 未知步骤3url, bytes.NewBuffer(data))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36")
	request.Header.Set("Origin", "https://www.djbhjg.net:4430")
	request.Header.Set("Referer", "https://www.djbhjg.net:4430/jgglweb/index.html")

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	response, _ := client.Do(request)
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {

	} else {
		fmt.Println("未知步骤3状态码：", response.StatusCode)
	}
}

func getcourselist(coursesid string) {
	courselisturl := "https://www.djbhjg.net:4430/ucc?AppID=YRWA_JGGLPT&Password=YRWA_JGGLPT2557013/&PageName=MHWZ&Controlname=courselist&return=true&SessionID=" + SessionID
	data, _ := json.Marshal(struct {
		CoursesId string `json:"courses_id"`
		UserId    string `json:"user_id"`
		Pagesize  int    `json:"pagesize"`
		Pagenum   int    `json:"pagenum"`
	}{
		CoursesId: coursesid,
		UserId:    userid,
		Pagesize:  1,
		Pagenum:   50,
	},
	)
	request, _ := http.NewRequest("POST", courselisturl, bytes.NewBuffer(data))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36")
	request.Header.Set("Origin", "https://www.djbhjg.net:4430")
	request.Header.Set("Referer", "https://www.djbhjg.net:4430/jgglweb/index.html")

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	response, _ := client.Do(request)
	defer response.Body.Close()

	f, _ := io.ReadAll(response.Body)
	var sss struct {
		Num  int `json:"num"`
		List []struct {
			Bid           string `json:"Bid"`
			VideoId       string `json:"video_id"`
			Url           string `json:"url"`
			VideoSize     string `json:"video_size"`
			Title         string `json:"title"`
			VideoDuration int    `json:"video_duration"`
			SortNum       int    `json:"sort_num"`
			CreateTime    string `json:"create_time"`
			LearnRate     int    `json:"learn_rate"`
			FinishStatu   int    `json:"finish_statu"`
			LearnTime     int    `json:"learn_time"`
		} `json:"list"`
	}
	json.Unmarshal(f, &sss)
	//fmt.Println("共有", sss.Num, "个视频")
	for y, _ := range sss.List {
		if sss.List[y].FinishStatu == 1 {
			//fmt.Println("视频名：", sss.List[y].Title, "已看完")
		} else {
			fmt.Println("待学习视频名：", sss.List[y].Title)
			learntime := sss.List[y].VideoDuration
			videoid := sss.List[y].VideoId
			// 刷
			time.Sleep(10 * time.Second)
			刷视频(coursesid, userid, videoid, learntime)
		}
	}

}

func 刷视频(coursesid string, userid string, videoid string, learntime int) {
	刷视频地址 := "https://www.djbhjg.net:4430/ucc?AppID=YRWA_JGGLPT&Password=YRWA_JGGLPT2557013/&PageName=MHWZ&Controlname=MCCN5&return=true&SessionID=" + SessionID
	data, _ := json.Marshal(struct {
		CoursesId string `json:"courses_id"`
		UserId    string `json:"user_id"`
		VideoId   string `json:"video_id"`
		LearnTime int    `json:"learn_time"`
	}{
		CoursesId: coursesid,
		UserId:    userid,
		VideoId:   videoid,
		LearnTime: learntime,
	},
	)
	request, _ := http.NewRequest("POST", 刷视频地址, bytes.NewBuffer(data))
	request.Header.Set("Accept", "application/json, text/plain, */*")
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.125 Safari/537.36")
	request.Header.Set("Origin", "https://www.djbhjg.net:4430")
	request.Header.Set("Referer", "https://www.djbhjg.net:4430/jgglweb/index.html")

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}}
	response, _ := client.Do(request)
	defer response.Body.Close()

	e, _ := io.ReadAll(response.Body)
	var as struct {
		ReturnValue int    `json:"ReturnValue"`
		Msg         string `json:"msg"`
	}
	json.Unmarshal(e, &as)
	if response.StatusCode == http.StatusOK {
		if as.ReturnValue == 1 && as.Msg == "提交成功" {
			//fmt.Println("此条视频已刷")
		}
	} else {
		fmt.Println("最后提交数据返回状态码异常", response.StatusCode)
	}
}
