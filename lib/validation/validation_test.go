package validation

import (
	"testing"
)

func TestRequired(t *testing.T) {
	valid := Valid{}
	if err := valid.Required(""); err == nil {
		t.Error("Require方法失败")
	}
	if err := valid.Required("test"); err != nil {
		t.Error("Require方法失败")
	}
}

func TestMin(t *testing.T) {
	valid := Valid{}
	if err := valid.Min("1", 2); err == nil {
		t.Error("Min方法失败")
	}
	if err := valid.Min("3", 2); err != nil {
		t.Error("Min方法失败")
	}
}

func TestMax(t *testing.T) {
	valid := Valid{}
	if err := valid.Max("1", 2); err != nil {
		t.Error("Max方法失败")
	}
	if err := valid.Max("3", 2); err == nil {
		t.Error("Max方法失败")
	}
}

func TestRange(t *testing.T) {
	valid := Valid{}
	if err := valid.Range("1", 1, 6); err != nil {
		t.Error("Range方法失败")
	}
	if err := valid.Range("0", 1, 6); err == nil {
		t.Error("Range方法失败")
	}
	if err := valid.Range("7", 1, 6); err == nil {
		t.Error("Range方法失败")
	}
}

func TestMinLength(t *testing.T) {
	valid := Valid{}
	if err := valid.MinLength("test", 5); err == nil {
		t.Error("MinLength方法失败")
	}
	if err := valid.MinLength("test", 4); err != nil {
		t.Error("MinLength方法失败")
	}
	if err := valid.MinLength("中文汉字", 5); err == nil {
		t.Error("MinLength方法失败")
	}
	if err := valid.MinLength("中文汉字", 4); err != nil {
		t.Error("MinLength方法失败")
	}
	if err := valid.MinLength("中w,，", 5); err == nil {
		t.Error("MinLength方法失败")
	}
	if err := valid.MinLength("中w,，", 4); err != nil {
		t.Error("MinLength方法失败")
	}
}

func TestMaxLength(t *testing.T) {
	valid := Valid{}
	if err := valid.MaxLength("test", 5); err != nil {
		t.Error("MaxLength方法失败")
	}
	if err := valid.MaxLength("test", 4); err != nil {
		t.Error("MaxLength方法失败")
	}
	if err := valid.MaxLength("中文汉字", 5); err != nil {
		t.Error("MaxLength方法失败")
	}
	if err := valid.MaxLength("中文汉字", 4); err != nil {
		t.Error("MaxLength方法失败")
	}
	if err := valid.MaxLength("中w,，", 5); err != nil {
		t.Error("MaxLength方法失败")
	}
	if err := valid.MaxLength("中w,，", 4); err != nil {
		t.Error("MaxLength方法失败")
	}
}

func TestLength(t *testing.T) {
	valid := Valid{}
	if err := valid.Length("test", 4); err != nil {
		t.Error("Length方法失败")
	}
	if err := valid.Length("bob", 4); err == nil {
		t.Error("Length方法失败")
	}
	if err := valid.Length("中文汉字", 4); err != nil {
		t.Error("Length方法失败")
	}
	if err := valid.Length("中w,，", 4); err != nil {
		t.Error("Length方法失败")
	}
}

func TestAlphaNumeric(t *testing.T) {
	valid := Valid{}
	if err := valid.AlphaNumeric("abc123"); err != nil {
		t.Error("AlphaNumeric方法失败")
	}
	if err := valid.AlphaNumeric("abc123啊"); err == nil {
		t.Error("AlphaNumeric方法失败")
	}
	if err := valid.AlphaNumeric("abcD123啊"); err == nil {
		t.Error("AlphaNumeric方法失败")
	}
	if err := valid.AlphaNumeric("abc123·"); err == nil {
		t.Error("AlphaNumeric方法失败")
	}
	if err := valid.AlphaNumeric("abc123`"); err == nil {
		t.Error("AlphaNumeric方法失败")
	}
	if err := valid.AlphaNumeric("*"); err == nil {
		t.Error("AlphaNumeric方法失败")
	}
	if err := valid.AlphaNumeric("-"); err == nil {
		t.Error("AlphaNumeric方法失败")
	}
}

func TestEmail(t *testing.T) {
	valid := Valid{}
	if err := valid.Email("test@qq.com"); err != nil {
		t.Error("Mail方法失败")
	}
	if err := valid.Email("test@@qq.com"); err == nil {
		t.Error("Mail方法失败")
	}
	if err := valid.Email("testqq.com"); err == nil {
		t.Error("Mail方法失败")
	}
	if err := valid.Email("test@qq"); err == nil {
		t.Error("Mail方法失败")
	}
}

func TestIP(t *testing.T) {
	valid := Valid{}
	if err := valid.IP("127.0.0.1"); err != nil {
		t.Error("IP方法失败")
	}
	if err := valid.IP("255.255.255.255"); err != nil {
		t.Error("IP方法失败")
	}
	if err := valid.IP("127.0.0.1.1"); err == nil {
		t.Error("IP方法失败")
	}
	if err := valid.IP("127.0.0.256"); err == nil {
		t.Error("IP方法失败")
	}
}

func TestBase64(t *testing.T) {
	valid := Valid{}
	if err := valid.Base64("YXNkYXNk"); err != nil {
		t.Error("Base64方法失败")
	}
	if err := valid.Base64("5pKS5aSn5Y+U5aSn5Y+U5aSn5Y+R"); err != nil {
		t.Error("Base64方法失败")
	}
	if err := valid.Base64("sefserh3u24df"); err == nil {
		t.Error("Base64方法失败")
	}
}

func TestMobile(t *testing.T) {
	valid := Valid{}
	if err := valid.Mobile("15626548844"); err != nil {
		t.Error("Mobile方法失败")
	}
	if err := valid.Mobile("156265488448"); err == nil {
		t.Error("Mobile方法失败")
	}
	if err := valid.Mobile("12345681234"); err == nil {
		t.Error("Mobile方法失败")
	}
}

func TestTel(t *testing.T) {
	valid := Valid{}
	if err := valid.Tel("2645835"); err != nil {
		t.Error("Tel方法失败")
	}
	if err := valid.Tel("65864253"); err != nil {
		t.Error("Tel方法失败")
	}
	if err := valid.Tel("546"); err == nil {
		t.Error("Tel方法失败")
	}
}

func TestMobileOrTel(t *testing.T) {
	valid := Valid{}
	if err := valid.MobileOrTel("15626548844"); err != nil {
		t.Error("MobileOrTel方法失败")
	}
	if err := valid.MobileOrTel("156265488448"); err == nil {
		t.Error("MobileOrTel方法失败")
	}
	if err := valid.MobileOrTel("12345681234"); err == nil {
		t.Error("MobileOrTel方法失败")
	}
	if err := valid.MobileOrTel("2645835"); err != nil {
		t.Error("MobileOrTel方法失败")
	}
	if err := valid.MobileOrTel("65864253"); err != nil {
		t.Error("MobileOrTel方法失败")
	}
	if err := valid.MobileOrTel("546"); err == nil {
		t.Error("MobileOrTel方法失败")
	}
}

func TestZipCode(t *testing.T) {
	valid := Valid{}
	if err := valid.ZipCode("100000"); err != nil {
		t.Error("ZipCode方法失败")
	}
	if err := valid.ZipCode("000000"); err == nil {
		t.Error("ZipCode方法失败")
	}
	if err := valid.ZipCode("1000000"); err == nil {
		t.Error("ZipCode方法失败")
	}
}
