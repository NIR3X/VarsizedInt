package varsizedint

import "testing"

func TestVarsizedInt(t *testing.T) {
	var (
		buffer   = make([]uint8, MaxSize)
		testData = []uint64{
			127, 16383, 2097151, 268435455,
			34359738367, 4398046511103, 562949953421311,
			72057594037927935, 18446744073709551615,
		}
	)

	for i := 0; i < len(testData); i++ {
		var testDataI = testData[i]
		Encode(buffer, testDataI)

		var (
			size         = i + 1
			sizeRead     = ParseSize(buffer)
			testDataRead = Decode(buffer)
		)

		t.Logf("Test case %d: size=%d, sizeRead=%d, testDataI=%d, testDataRead=%d",
			i+1, size, sizeRead, testDataI, testDataRead)

		if size != sizeRead {
			t.Errorf("Test case %d failed: size != sizeRead", i+1)
		}

		if testDataI != testDataRead {
			t.Errorf("Test case %d failed: testDataI != testDataRead", i+1)
		}
	}
}
