/*
Copyright © 2023 Steve Francia <spf@spf13.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// examCmd represents the exam command
var (
	examTaken = false
	examCmd   = &cobra.Command{
		Use:   "exam",
		Short: "Will exam all the questions",
		Long:  `Will exam all the questions in the quiz and the options for each question.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			err := takeExam(cmd, args)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(0)
			}
			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(examCmd)
}

func takeExam(cmd *cobra.Command, args []string) error {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Please choose an option:")
		fmt.Println("1. Take Exam")
		fmt.Println("2. Show Results")
		fmt.Println("3. Quit")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			if examTaken {
				fmt.Println("You have already taken the exam.")
				continue
			}
			examTaken = true
			questions, err := getQuestions()
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, question := range questions {
				question.Print()
				answer, _ := reader.ReadString('\n')
				answerID, err = strconv.Atoi(strings.TrimSpace(answer))
				if err != nil {
					return err
				}

				//send the answer in the answerQuestion function
				questionID = question.ID
				err := answerQuestion(&cobra.Command{}, []string{})
				if err != nil {
					return err
				}
				fmt.Println("Your answer:", answer)
			}
		case "2":
			fmt.Println("Showing Results...")
			getStats(&cobra.Command{}, []string{})
		case "3":
			fmt.Println("Exiting program...")
			return nil
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func getQuestions() ([]Question, error) {
	url := fmt.Sprintf("http://%s:%d/questions", host, port)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch questions: %s", resp.Status)
	}

	var questions []Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		return nil, err
	}

	return questions, nil
}
