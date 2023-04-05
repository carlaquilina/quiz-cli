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
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Will list all the questions",
	Long:  `Will list all the questions in the quiz and the options for each question.`,
	Run: func(cmd *cobra.Command, args []string) {
		listQuestions(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listQuestions(cmd *cobra.Command, args []string) error {
	url := fmt.Sprintf("http://%s:%d/questions", host, port)
	fmt.Println("URL:>", url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to fetch questions: %s", resp.Status)
	}

	var questions []Question
	err = json.NewDecoder(resp.Body).Decode(&questions)
	if err != nil {
		return err
	}

	fmt.Println("Quiz Questions:")
	for _, q := range questions {
		fmt.Printf("ID: %d\nQuestion: %s\nOptions: %v\n\n", q.ID, q.Text, q.Answers)
	}

	return nil
}
