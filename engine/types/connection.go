// @Title  connection.go
// @Description  流程节点连接
// @Create  heyitong  2022/6/23 17:16
// @Update  heyitong  2022/6/23 17:16

package types

import (
	"reflect"
	"sync"
	"time"
	"tiops/common/services"
)

type ConnectionInfo struct {
	OutputID   string
	OutputName string
	InputID    string
	InputName  string
}

type Connection struct {
	ID         string
	Variable   Variable
	SourceNode *Node
	TargetNode *Node
	Info       *ConnectionInfo
	DataChan   chan *services.ActionData
	hasDone    bool
	rwLock     sync.RWMutex
}

type InputConnections []*Connection
type OutputConnections []*Connection

type InputConnectionsMap map[string]InputConnections
type OutputConnectionsMap map[string]OutputConnections

func (ics InputConnections) VarCount() int {
	count := 0
	for _, connection := range ics {
		if connection.Variable != nil {
			count++
		}
	}
	return count
}

func (ics InputConnections) HasInputQueue() bool {

	for _, connection := range ics {
		if connection.DataChan != nil && !connection.HasDone() {
			return true
		}
	}
	return false
}

func (ics InputConnections) HasVarOnly() bool {

	return len(ics) > 0 && len(ics)-ics.VarCount() < 1
}

func (ics InputConnections) SelectInputOnce() (*services.ActionData, bool) {

	var selectCase []reflect.SelectCase
	var variableData *services.ActionData

	//fmt.Println(ics)

	for _, connection := range ics {
		if connection.DataChan != nil {
			if connection.HasDone() && len(connection.DataChan) < 1 {
				continue
			}
			selectCase = append(selectCase, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(connection.DataChan)})
		} else {
			variableData = connection.Variable.ToActionArguments(1)
		}
	}
	if len(selectCase) > 0 {
		_, recv, recvOk := reflect.Select(selectCase)
		//fmt.Println(i, recv, recvOk)
		if recvOk {
			//fmt.Println(recv)
			return recv.Interface().(*services.ActionData), true
		}
	}

	//fmt.Println(variableData)

	return variableData, variableData != nil
}

// SelectInput 从多个输入队列中选择一条输入，输入全部关闭后返回nil
func (ics InputConnections) SelectInput() (*services.ActionData, bool) {

	var selectCase []reflect.SelectCase
	var variableData *services.ActionData

	//fmt.Println(ics)

	//logger.Error(len(ics))

	for _, connection := range ics {
		if connection.DataChan != nil {
			if connection.HasDone() && len(connection.DataChan) < 1 {
				continue
			}
			selectCase = append(selectCase, reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(connection.DataChan)})
		} else {
			variableData = connection.Variable.ToActionArguments(1)
		}
	}

	//logger.Warning(selectCase)

	if len(selectCase) > 0 {
		_, recv, recvOk := reflect.Select(selectCase)
		//fmt.Println(i, recv, recvOk)
		//logger.Warning(recv.Type(), recvOk)
		if recvOk {
			//fmt.Println(recv)

			if recv.Interface() == nil {
				retryCount := 0
				for data, _ := ics.SelectInputOnce(); data == nil || retryCount < 10; retryCount++ {
					time.Sleep(100 * time.Millisecond)
					data, _ = ics.SelectInput()
				}
			}

			return recv.Interface().(*services.ActionData), true
		}
	}

	//fmt.Println(variableData)

	return variableData, variableData != nil
}

func (ics InputConnections) BacklogCount() int {
	count := 0
	for _, connection := range ics {
		count += len(connection.DataChan)
	}

	return count
}

// Connection method implement

func (c *Connection) HasDone() bool {
	c.rwLock.RLock()
	defer c.rwLock.RUnlock()
	return c.hasDone
}

// Done 结束数据队列
func (c *Connection) Done() {
	c.rwLock.Lock()
	defer c.rwLock.Unlock()
	if c.hasDone {
		return
	}
	c.hasDone = true

	if c.DataChan != nil {
		close(c.DataChan)
	}
}

// InputConnectionsMap method implement

func (icsm InputConnectionsMap) HasVarOnly() bool {
	if len(icsm) < 1 {
		return false
	}
	for _, connections := range icsm {
		if connections.HasVarOnly() {
			continue
		}
		for _, connection := range connections {
			if !connection.SourceNode.Inputs.HasVarOnly() {
				return false
			}
		}
	}
	return true
}
