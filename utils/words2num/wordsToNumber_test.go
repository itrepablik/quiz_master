package words2num

import "testing"

func TestConvertWordsToNumber(t *testing.T) {
	wordNum := "ten"
	num, err := ConvertWordsToNumber(wordNum)
	if err != nil {
		t.Errorf("word is invalid or not in a collections yet, got %v", err)
		return
	}
	if num != 10 {
		t.Errorf("Expected %d, got %d", 10, num)
	}
	t.Logf("%s is %d", wordNum, num)
}
