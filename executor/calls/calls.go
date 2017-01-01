package calls

import (
	"github.com/mesos/mesos-go"
	"github.com/mesos/mesos-go/executor"
)

// Framework sets a executor.Call's FrameworkID
func Framework(id string) executor.CallOpt {
	return func(c *executor.Call) {
		c.FrameworkID = mesos.FrameworkID{Value: id}
	}
}

// Executor sets a executor.Call's ExecutorID
func Executor(id string) executor.CallOpt {
	return func(c *executor.Call) {
		c.ExecutorID = mesos.ExecutorID{Value: id}
	}
}

// Subscribe returns an executor call with the given parameters.
func Subscribe(unackdTasks []mesos.TaskInfo, unackdUpdates []executor.Call_Update) *executor.Call {
	return &executor.Call{
		Type: executor.Call_SUBSCRIBE.Enum(),
		Subscribe: &executor.Call_Subscribe{
			UnacknowledgedTasks:   unackdTasks,
			UnacknowledgedUpdates: unackdUpdates,
		},
	}
}

// Update returns an executor call with the given parameters.
func Update(status mesos.TaskStatus) *executor.Call {
	return &executor.Call{
		Type: executor.Call_UPDATE.Enum(),
		Update: &executor.Call_Update{
			Status: status,
		},
	}
}

// Message returns an executor call with the given parameters.
func Message(data []byte) *executor.Call {
	return &executor.Call{
		Type: executor.Call_MESSAGE.Enum(),
		Message: &executor.Call_Message{
			Data: data,
		},
	}
}
