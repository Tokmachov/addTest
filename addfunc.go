package main

import (
	"os"
	"log"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	if (len(os.Args) != 3) {
		return
	}
	hpp, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	testFuncName, testsCount := getLastTestFuncNameAndNum(hpp)
	hpp.Close();
	cpp, err := os.OpenFile(os.Args[2], os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	insertTestToCpp(cpp, testFuncName, testsCount)
	cpp.Close()
}

func getLastTestFuncNameAndNum(f *os.File) (string, int) {
	scanner := bufio.NewScanner(f)
	var tName string
	var tNum int
	for scanner.Scan() {
        if strings.Contains(scanner.Text(), "static void test") {
			tNum += 1
		}
		if scanner.Text() == "};" {
			break
		}
		tName = scanner.Text()
    }
	tName = strings.TrimSpace(tName)
	tName = strings.TrimPrefix(tName, "static void ")
	tName = strings.TrimSuffix(tName, "();")
	return tName, tNum
}

func insertTestToCpp(f *os.File, tName string, tNum int) {
	scanner := bufio.NewScanner(f)
	var cppLines []string
	numOfTestPushBlocksPassed := 0
	numOfTestFuncDefinishionsPassed := 0
	testCaseClassName := ""
	for scanner.Scan() {
		cppLines = append(cppLines, scanner.Text())
		if strings.Contains(scanner.Text(), "::") {
			comps := strings.Split(scanner.Text(), "::")
			testCaseClassName = comps[0]
			break
		}
	}
	for scanner.Scan() {
		cppLines = append(cppLines, scanner.Text())
		if strings.Contains(scanner.Text(), "allTests.push_back") {
			numOfTestPushBlocksPassed += 1;
		}
		if strings.Contains(scanner.Text(), "void " + testCaseClassName + "::test") {
			numOfTestFuncDefinishionsPassed += 1
		}
		if numOfTestPushBlocksPassed == (tNum - 1) {
			pushStr := makeTestPushStr(tName, tNum, testCaseClassName)
			cppLines = append(cppLines, pushStr)
			numOfTestPushBlocksPassed = 0
		}
		if (numOfTestFuncDefinishionsPassed == (tNum - 1)) && (scanner.Text() == "}") {
			testDefinition := makeTestFuncDefinition(tName, testCaseClassName)
			cppLines = append(cppLines, testDefinition)
		}
	}
	_, _ = f.Seek(0, 0)
	datawriter := bufio.NewWriter(f)
	for _, line := range cppLines {
		_, _ = datawriter.WriteString(line + "\n")
	}
	datawriter.Flush()
}

func makeTestPushStr(tName string, tNum int, tClassName string) string {
	testPushStr := ""
	testPushStr += "\tTestNameAndFunc t" + strconv.Itoa(tNum) + " =\n"
	testPushStr += "\t{\n"
	testPushStr += "\t\t\"" + tName + "\", \n"
	testPushStr += "\t\t" + tName + "\n"
	testPushStr += "\t};" + "\n"
	testPushStr += "\t" + tClassName + "::allTests.push_back(t" + strconv.Itoa(tNum) +");"
	return testPushStr
}
func makeTestFuncDefinition(tName string, tClassName string) string {
	testPushStr := "void " + tClassName + "::" + tName + "()\n"
	testPushStr += "{\n"
	testPushStr += "\t//Arrange\n"
	testPushStr += "\n"
	testPushStr += "\t//Act\n"
	testPushStr += "\n"
	testPushStr += "\t//Assert\n"
	testPushStr += "\t//assertTrue();\n"
	testPushStr += "\t//assertFalse();\n"
	testPushStr += "\t//assertEqual(,);\n"
	testPushStr += "}\n"
	return testPushStr
}
