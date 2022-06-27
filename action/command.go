// @Title  command.go
// @Description  执行命令组件封装
// @Create  heyitong  2022/6/27 9:21
// @Update  heyitong  2022/6/27 9:21

package action

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"time"
	"tiops/action/types"
)

type BatchMode uint8 // 批量数据处理模式

const (
	ProcessedSeparately BatchMode = iota // 分多次处理
	PackagePacket                        // 封装为数据包
	DelimiterSplit                       // 用分隔符分割
)

const (
	DefaultInputName  = "in"
	DefaultOutputName = "out"
	DefaultArgsName   = "args"
	DefaultDelimiter  = byte(0)
)

type batchFunc func(cmd string, args []string, options CommandActionOptions, input [][]byte) ([][]byte, error)

type CommandActionOptions struct {
	BatchMode  BatchMode
	Delimiter  byte
	InputName  string
	OutputName string
	ArgsName   string
	UseKVArgs  bool
}

type CommandAction struct {
	CommandActionOptions
	cmd  string
	args []string
}

func (a *CommandAction) processData(input [][]byte) ([][]byte, error) {
	return batchFuncMap[a.BatchMode](a.cmd, a.args, a.CommandActionOptions, input)
}

func (a *CommandAction) processItem(item types.ActionDataItem) (types.ActionDataItem, error) {
	out, err := a.processData([][]byte{item[a.InputName].([]byte)})
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{a.OutputName: out[0]}, nil
}

func (a *CommandAction) ProcessItems(itemBatch types.ActionDataBatch) (types.ActionDataBatch, error) {
	inputs := make([][]byte, len(itemBatch))

	for i, i2 := range itemBatch {
		inputs[i] = i2[a.InputName].([]byte)
	}

	outputs, err := a.processData(inputs)

	if err != nil {
		return nil, err
	}

	outputData := make(types.ActionDataBatch, len(outputs))

	for i, output := range outputs {
		outputData[i] = types.ActionDataItem{a.OutputName: output}
	}

	return outputData, err
}

func (a *CommandAction) RegisterNode(ctx *types.NodeRegisterContext) error {

	argString := ctx.ActionOptions.GetString(a.ArgsName)

	cmd := ctx.ActionOptions.GetString("cmd")

	if cmd != "" {
		a.cmd = cmd
	}

	a.args = append(a.args, strings.Split(argString, " ")...)

	return nil
}

func (a *CommandAction) CallBatch(ctx *types.BatchRequestContext) types.ActionDataBatch {

	//items, err := a.processItem(ctx.Inputs.Map(func(item types.ActionDataItem) types.ActionDataItem {
	//	return item
	//}))

	inputs := make([][]byte, ctx.Inputs[a.InputName].Count())

	for i, i2 := range ctx.Inputs[a.InputName].List() {
		inputs[i] = i2.([]byte)
	}

	outputs, err := a.processData(inputs)

	if err != nil {
		ctx.Logger.Error(err.Error())
		return nil
	}

	outputData := make(types.ActionDataBatch, len(outputs))

	for i, output := range outputs {
		outputData[i] = types.ActionDataItem{a.OutputName: output}
	}

	return outputData
}

func (a *CommandAction) Call(ctx *types.PieceRequestContext) types.ActionDataItem {

	//a.CallBatch(&types.BatchRequestContext{
	//	ActionNodeContext: ctx.ActionNodeContext,
	//	Inputs:,
	//})
	item, err := a.processItem(ctx.Input)
	//fmt.Println(err)
	if err != nil {
		ctx.Logger.Error(err.Error())
		return nil
	}
	return item
}

func init() {
	//exec.Command().StdinPipe()
}

func NewCommandAction(cmd string, options CommandActionOptions) *CommandAction {

	ca := strings.Split(cmd, " ")

	var remainder []string

	if len(ca) > 1 {
		remainder = ca[1:]
	}

	if options.InputName == "" {
		options.InputName = DefaultInputName
	}

	if options.OutputName == "" {
		options.OutputName = DefaultOutputName
	}

	if options.ArgsName == "options" {
		options.ArgsName = DefaultArgsName
	}

	return &CommandAction{
		CommandActionOptions: options,
		cmd:                  ca[0],
		args:                 remainder,
	}
}

var batchFuncMap = map[BatchMode]batchFunc{
	ProcessedSeparately: func(cmd string, args []string, options CommandActionOptions, input [][]byte) ([][]byte, error) {

		output := make([][]byte, len(input))

		// fmt.Println(cmd, args, string(input[0]))

		for i, bytes := range input {
			execCmd := exec.Command(cmd, args...)

			stdinPipe, err := execCmd.StdinPipe()

			if err != nil {
				return nil, err
			}

			_, err = stdinPipe.Write(bytes)
			if err != nil {
				return nil, err
			}

			stdinPipe.Close()

			out, err := execCmd.Output()

			if err != nil {
				return nil, err
			}

			output[i] = out

			//execCmd.Process.Kill()
		}

		return output, nil

	},
	PackagePacket: func(cmd string, args []string, options CommandActionOptions, input [][]byte) ([][]byte, error) {
		execCmd := exec.Command(cmd, args...)

		stdinPipe, err := execCmd.StdinPipe()

		if err != nil {
			return nil, err
		}

		stdoutPipe, err := execCmd.StdoutPipe()

		if err != nil {
			return nil, err
		}

		err = execCmd.Start()

		if err != nil {
			return nil, err
		}

		go func() {
			stdinPipe.Write([]byte(fmt.Sprintf("%d\n", len(input))))
			for _, bytes := range input {
				stdinPipe.Write([]byte(fmt.Sprintf("%d\n", len(bytes))))
				stdinPipe.Write(bytes)
			}
		}()

		reader := bufio.NewReader(stdoutPipe)

		var outputCount int

		_, err = fmt.Fscanf(reader, "%d\n", &outputCount)

		if err != nil {
			return nil, err
		}

		var output [][]byte

		for i := 0; i < outputCount; i++ {

			var byteCount int

			_, err = fmt.Fscanf(reader, "%d\n", &byteCount)

			if err != nil {
				return nil, err
			}

			out := make([]byte, byteCount)

			_, err = io.ReadFull(reader, out)

			if err != nil {
				return nil, err
			}

			output = append(output, out)
		}

		return output, nil
	},
	DelimiterSplit: func(cmd string, args []string, options CommandActionOptions, input [][]byte) ([][]byte, error) {
		execCmd := exec.Command(cmd, args...)

		stdinPipe, err := execCmd.StdinPipe()

		if err != nil {
			return nil, err
		}

		stdoutPipe, err := execCmd.StdoutPipe()

		if err != nil {
			return nil, err
		}

		var output [][]byte

		go func() {

			reader := bufio.NewReader(stdoutPipe)

			for true {
				readBytes, err := reader.ReadBytes(options.Delimiter)

				if err == io.EOF {
					break
				}

				if err != nil {
					output = append(output, []byte(err.Error()))
				}

				if len(readBytes) < 1 {
					break
				}

				output = append(output, readBytes)
			}

		}()

		go func() {
			//wg.Done()
			for _, bytes := range input {
				stdinPipe.Write(bytes)
				stdinPipe.Write([]byte{options.Delimiter})
			}
			stdinPipe.Write([]byte{options.Delimiter})
			time.Sleep(time.Second)
			stdinPipe.Close()
		}()

		//wg.Wait()

		err = execCmd.Start()

		if err != nil {
			return nil, err
		}

		err = execCmd.Wait()

		return output, err
	},
}
