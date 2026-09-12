package task

import (
	"code.gitea.io/gitea/entity"
	"code.gitea.io/gitea/modules/log"
	"code.gitea.io/gitea/routers/response"
	"code.gitea.io/gitea/services/ai_task_service/context"
	"errors"
	"fmt"
)

type CreateFunc func(ctx *context.CreationContext) *response.BizError

type CreationFuncNode struct {
	Funcs   []CreateFunc
	IsAsync bool
	ErrFunc CreateFunc
}

func (node CreationFuncNode) Run(ctx *context.CreationContext) *response.BizError {
	var err *response.BizError
	if node.IsAsync {
		go runFuncNode(node, ctx)
	} else {
		err = runFuncNode(node, ctx)
	}
	return err
}

type CreateOperator struct {
	FuncArray []CreationFuncNode
}

//添加同步节点
func (o *CreateOperator) Next(f ...CreateFunc) *CreateOperator {
	if o.FuncArray == nil {
		o.FuncArray = make([]CreationFuncNode, 0)
	}
	o.FuncArray = append(o.FuncArray, CreationFuncNode{Funcs: f, IsAsync: false})
	return o
}

//添加异步节点
func (o *CreateOperator) AsyncNext(f ...CreateFunc) *CreateOperator {
	if o.FuncArray == nil {
		o.FuncArray = make([]CreationFuncNode, 0)
	}
	o.FuncArray = append(o.FuncArray, CreationFuncNode{Funcs: f, IsAsync: true})
	return o
}

//添加同步节点，参数的最后一个Fun是异常处理节点，其他的Fun是正常节点
//只有当正常节点返回error时，异常处理节点才会执行
func (o *CreateOperator) NextWithErrFun(f ...CreateFunc) *CreateOperator {
	if o.FuncArray == nil {
		o.FuncArray = make([]CreationFuncNode, 0)
	}
	if f == nil || len(f) == 0 {
		return o
	}
	if len(f) < 2 {
		log.Error("AsyncNextWithErrFun err.funcs are less than 2")
		return o
	}

	errFun := f[len(f)-1]
	normalFun := f[0 : len(f)-1]
	o.FuncArray = append(o.FuncArray, CreationFuncNode{Funcs: normalFun, IsAsync: false, ErrFunc: errFun})
	return o
}

//添加异步节点，参数的最后一个Fun是异常处理节点，其他的Fun是正常节点
//只有当正常节点返回error时，异常处理节点才会执行
func (o *CreateOperator) AsyncNextWithErrFun(f ...CreateFunc) *CreateOperator {
	if o.FuncArray == nil {
		o.FuncArray = make([]CreationFuncNode, 0)
	}
	if f == nil || len(f) == 0 {
		return o
	}
	if len(f) < 2 {
		log.Error("AsyncNextWithErrFun err.funcs are less than 2")
		return o
	}

	errFun := f[len(f)-1]
	normalFun := f[0 : len(f)-1]
	o.FuncArray = append(o.FuncArray, CreationFuncNode{Funcs: normalFun, IsAsync: true, ErrFunc: errFun})
	return o
}

func (o *CreateOperator) Operate(ctx *context.CreationContext) *response.BizError {
	var err *response.BizError
	for i := 0; i < len(o.FuncArray); i++ {
		node := o.FuncArray[i]
		err = node.Run(ctx)
		if err != nil {
			log.Error("Operate err.%v", err)
			break
		}
	}
	return err
}

func runFuncNode(node CreationFuncNode, ctx *context.CreationContext) *response.BizError {
	var err *response.BizError
	defer func() {
		if tmpErr := recover(); tmpErr != nil {
			combinedErr := fmt.Errorf("%s\n%s", err, log.Stack(2))
			log.Error("PANIC:%v", combinedErr)
		}
		if err != nil && node.ErrFunc != nil {
			ctx.Response = &entity.CreationResponse{
				Error: errors.New(err.DefaultMsg),
			}
			newErr := node.ErrFunc(ctx)
			if newErr != nil {
				log.Error("runFuncNode ErrFunc error.%v", err)
				return
			}
		}
	}()
	for _, f := range node.Funcs {
		err = f(ctx)
		if err != nil {
			log.Error("runFuncNode err.%v", err)
			break
		}
	}
	return err
}
