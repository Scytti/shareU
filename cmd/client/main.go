package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"
)

// API URLs
const (
	allocateURL = "http://localhost:8082/api/v1/tasks/allocate"
	submitURL   = "http://localhost:8082/api/v1/tasks/submit"
	agentIP     = "10.163.8.144"
)

// TaskAllocateResponse - структура ответа от сервера на allocate
type TaskAllocateResponse struct {
	TaskID    int    `json:"task-id"`
	Command   string `json:"command"`
	Condition string `json:"condition,omitempty"`
	After     string `json:"after,omitempty"`
}

// TaskSubmitRequest - структура запроса на submit
type TaskSubmitRequest struct {
	ID     int    `json:"id"`
	IP     string `json:"ip"`
	Result string `json:"result"`
}

// sendAllocateRequest отправляет запрос на выделение задачи
func sendAllocateRequest() (*TaskAllocateResponse, error) {
	data := map[string]string{"ip": agentIP}
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("ошибка кодирования JSON: %v", err)
	}

	resp, err := http.Post(allocateURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("ошибка HTTP-запроса: %v", err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Ответ сервера allocate:", string(body)) // Логируем ответ

	var task TaskAllocateResponse
	if err := json.Unmarshal(body, &task); err != nil {
		return nil, fmt.Errorf("ошибка парсинга JSON: %v", err)
	}

	return &task, nil
}

// sendSubmitRequest отправляет результат выполнения задачи на сервер
func sendSubmitRequest(taskID int, result string) error {
	data := TaskSubmitRequest{
		ID:     taskID,
		IP:     agentIP,
		Result: result,
	}
	jsonData, _ := json.Marshal(data)

	resp, err := http.Post(submitURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("server returned status: %d", resp.StatusCode)
	}

	return nil
}

// executeCommand выполняет команду и ждет появления condition в выводе
func executeCommand(task *TaskAllocateResponse) (string, error) {
	cmd := exec.Command("sh", "-c", task.Command)

	cmd.Env = os.Environ()

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return "", err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return "", err
	}

	if err := cmd.Start(); err != nil {
		return "", err
	}

	var outputBuffer bytes.Buffer
	scanner := bufio.NewScanner(io.MultiReader(stdout, stderr))

	for scanner.Scan() {
		line := scanner.Text()
		outputBuffer.WriteString(line + "\n")

		// Если найдено условие, прерываемся
		if strings.Contains(line, task.Condition) {
			cmd.Process.Kill()
			break
		}
	}

	cmd.Wait()
	return outputBuffer.String(), nil
}

func main() {
	cmd := exec.Command("sh", "-c", "cd $UTOPIA_HOME && ./utopia.sh && exit")

	cmd.Env = os.Environ()
	for _, env := range os.Environ() {
		fmt.Println(env)
	}
	for {
		task, err := sendAllocateRequest()
		if err != nil {
			fmt.Println("Ошибка запроса allocate:", err)
			time.Sleep(20 * time.Second)
			continue
		}

		if task.TaskID == 0 {
			fmt.Println("Нет доступных задач. Ожидание 20 секунд...")
			time.Sleep(20 * time.Second)
			continue
		}

		fmt.Println("Получена задача:", task.TaskID)
		result, err := executeCommand(task)
		if err != nil {
			fmt.Println("Ошибка выполнения команды:", err)
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("Отправка результата на сервер...")
		if err := sendSubmitRequest(task.TaskID, result); err != nil {
			fmt.Println("Ошибка отправки результата:", err)
		}
	}
}
