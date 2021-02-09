package ioc

import (
	"errors"
	"fmt"
	"reflect"
)

// Bean is a representation of Golang Object.
type Bean struct {
	name           string
	dep            map[string]string
	val            interface{}
	reflectedValue reflect.Value
	fin            bool
	visit          bool
}

// AutowireContainer represents Ioc Container
type AutowireContainer struct {
	beanMap map[string]*Bean
}

// New will create an Ioc container
func New() *AutowireContainer {
	return &AutowireContainer{beanMap: make(map[string]*Bean)}
}

// Add a bean definition into the Ioc Container
func (ctx *AutowireContainer) Add(beanName string, val interface{}) {
	ctx.beanMap[beanName] = &Bean{name: beanName, reflectedValue: reflect.ValueOf(val)}
}

// AddDep will add a bean dependency into the Ioc Container
func (ctx *AutowireContainer) AddDep(beanName string, fieldName string, toInject string) error {
	var bean *Bean
	if bean = ctx.beanMap[beanName]; bean == nil {
		return errors.New("bean name not exists")
	}
	if bean.dep == nil {
		bean.dep = make(map[string]string)
	}
	bean.dep[fieldName] = toInject
	return nil
}

// Refresh tries to create all beans and do the dep injection.
func (ctx *AutowireContainer) Refresh() {
	// reset
	for _, v := range ctx.beanMap {
		v.visit = false
	}
	// populate
	for k := range ctx.beanMap {
		ctx.populate(k)
	}
}

func (ctx *AutowireContainer) populate(beanName string) string {
	bean := ctx.beanMap[beanName]
	if bean == nil {
		panic(fmt.Sprintf("Bean name %s not exists!", beanName))
	}
	if bean.visit == false && bean.fin == false && bean.dep != nil {
		bean.visit = true
		fieldNum := bean.reflectedValue.Elem().NumField()
		beanType := bean.reflectedValue.Elem().Type()
		for i := 0; i < fieldNum; i++ {
			fieldValue := bean.reflectedValue.Elem().Field(i)
			fieldName := beanType.Field(i).Name
			if bean.dep[fieldName] != "" {
				injectBean := ctx.beanMap[bean.dep[fieldName]]
				fieldValue.Set(injectBean.reflectedValue)
			}
		}
	}
	bean.visit = true
	bean.fin = true
	bean.val = bean.reflectedValue.Interface()
	return bean.name
}

// Get a bean from the Ioc Container
func (ctx *AutowireContainer) Get(beanName string) interface{} {
	bean := ctx.beanMap[beanName]
	if bean.fin == false {
		return nil
	}
	return bean.val
}
