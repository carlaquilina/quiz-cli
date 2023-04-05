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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
)

var (
	questionID int
	answerID   int
)

// answerCmd represents the answer command
var answerCmd = &cobra.Command{
	Use:   "answer",
	Short: "Submit an answer to a question",
	Long:  `Submit an answer to a question specified by ID.`,
	Run: func(cmd *cobra.Command, args []string) {
		err := answerQuestion(cmd, args)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(answerCmd)
	answerCmd.Flags().IntVarP(&questionID, "question-id", "q", 0, "ID of the question to answer (required)")
	answerCmd.MarkFlagRequired("question-id")
	answerCmd.Flags().IntVarP(&answerID, "answer-id", "a", 0, "ID of the answer (required)")
	answerCmd.MarkFlagRequired("answer-id")
}

func answerQuestion(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("http://%s:%d/answers", host, port)
	answer := UserAnswer{
		QuestionID: questionID,
		AnswerID:   answerID,
	}
	fmt.Printf("questionID: %d, answerID: %d", questionID, answerID)
	jsonValue, _ := json.Marshal(answer)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s failed to submit answer: %s", resp.Status, string(respBody))
	}

	fmt.Println("Answer submitted successfully.")
	return nil
}
