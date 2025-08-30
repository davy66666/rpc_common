package errorx

import (
	"fmt"
	"os"
	"strconv"
	"testing"

	"github.com/bregydoc/gtranslate"
	json "github.com/bytedance/sonic"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/text/language"
)

// 读取本地的errors.xlsx文件，生成error.json文件
func TestBuildExcelErrCode(t *testing.T) {
	f, err := excelize.OpenFile("errors.xlsx")
	if err != nil {
		logx.Error("打开 Excel 失败:", err)
	}
	defer f.Close()

	sheet := f.GetSheetName(0)
	rows, err := f.GetRows(sheet)
	if err != nil {
		logx.Error("读取 Sheet 内容失败:", err)
	}
	// 使用 slice 保持顺序，然后拼接 JSON 手动输出
	type entry struct {
		Code string
		Zh   string
		En   string
		Pt   string
	}
	var data []entry
	for i, row := range rows {
		if i == 0 || len(row) < 4 {
			continue
		}
		data = append(data, entry{
			Code: row[0],
			Zh:   row[1],
			En:   row[2],
			Pt:   row[3],
		})
	}
	// 构建 JSON 字符串
	result := "{"
	for i, item := range data {
		entryStr := fmt.Sprintf(`"%s":{"zh":"%s","pt":"%s","en":"%s"}`, item.Code, item.Zh, item.Pt, item.En)
		result += entryStr
		if i < len(data)-1 {
			result += ","
		}
	}
	result += "}"
	// 写入文件（单行）
	err = os.WriteFile("error.json", []byte(result), 0644)
	if err != nil {
		logx.Error("写入文件失败:", err)
	}
	fmt.Println("✅ 已写入单行格式的 error.json")
}

func TestErrCode(t *testing.T) {
	var m = map[string]map[string]string{}
	for code, msg := range errorMsg {
		// 中文
		v, ok := m[strconv.Itoa(code)]
		if !ok {
			v = make(map[string]string)
		}
		v["zh"] = msg
		fmt.Println("Code: ", code, " Msg: ", msg)
		// 巴西葡萄牙语
		pt, err := translate(msg, language.SimplifiedChinese.String(), language.BrazilianPortuguese.String())
		if err != nil {
			t.Error(err)
			return
		}
		v["pt"] = pt

		// 英语
		en, err := translate(msg, language.SimplifiedChinese.String(), language.English.String())
		if err != nil {
			t.Error(err)
			return
		}
		v["en"] = en

		m[strconv.Itoa(code)] = v
	}

	marshal, _ := json.Marshal(m)
	file, err := os.OpenFile("error.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		t.Error(err)
		return
	}

	defer func(file *os.File) { _ = file.Close() }(file)
	_, err = file.Write(marshal)
	if err != nil {
		t.Error(err)
		return
	}
}

func translate(text string, src, dst string) (string, error) {
	defer func() {
		if err := recover(); err != nil {
			logx.Error(err)
			logx.Error("请求翻译的文字", text)
		}
	}()

	translated, err := gtranslate.TranslateWithParams(
		text,
		gtranslate.TranslationParams{
			From: src,
			To:   dst,
		},
	)
	return translated, err
}
