package internal

import (
	"fmt"
	capsolver_go "github.com/capsolver/capsolver-go"
	"os"
)

func SolveCap() (string, string) {
	capSolver := capsolver_go.CapSolver{ApiKey: os.Getenv("CAP_SOLVER_API_KEY")}

	proxy := os.Getenv("PROXY")
	task := map[string]any{
		"websiteURL": "https://erb.minjust.gov.ua",
		"websiteKey": "6LevzOUUAAAAAGjAekCNws95tBDm5m69m5LT4L7X",
		"pageAction": "search_person",
	}

	if len(proxy) != 0 {
		task["type"] = "ReCaptchaV3Task"
		task["proxy"] = proxy
	} else {
		task["type"] = "ReCaptchaV3TaskProxyLess"
	}

	s, err := capSolver.Solve(task)

	if err != nil {
		fmt.Println("Error during solving the captcha!")
		os.Exit(7)
	}

	return s.Solution.GRecaptchaResponse, s.Solution.UserAgent
}
