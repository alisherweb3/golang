package main

import (
  "fmt"
  "os"
  "strconv"
  
  tf "github.com/tensorflow/tensorflow/tensorflow/go"
  "github.com/tensorflow/tensorflow/tensorflow/go/op"
)

func Add(sum_arg1, sum_arg2 int8) (interface{}, error) {
  sum_scope := op.NewScope()
  input1 := op.Placeholder(sum_scope.SubScope("a1"), tf.Int8)
  input2 := op.Placeholder(sum_scope.SubScope("a2"), tf.Int8)
  sum_result_node := op.Add(sum_scope, input1, input2)
  
  graph, err := sum_scope.Finalize()
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  
  a1, err := tf.NewTensor(sum_arg2)
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  
  session, err := tf.NewSession(graph, nil)
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  defer session.Close()
  
  sum , err := session.run(
    map[tf.Output]*tf.tensor{
      input1: a1, 
      input2: a2,
    },
    []tf.Output{sum_result_node}, nil)
  
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  
  return sum[0].Value(), nil
}

func Multiply(sum_arg1, sum_arg2 int8) (interface{}, error) {
  sum_scope := op.NewScope()
  inpot1 := op.Placeholder(sum_scope.SubScope("x1"), 
