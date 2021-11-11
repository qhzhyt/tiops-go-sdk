package utils

import (
	"bytes"
	"compress/gzip"
	"container/list"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/qhzhyt/tiops-go-sdk/common/models"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
	"io/ioutil"
	"sync"
)

func ProtoDes() *descriptorpb.FileDescriptorProto {

	user := &models.User{Name: "hyt"}
	descriptor, _ := user.Descriptor()

	reader, _ := gzip.NewReader(bytes.NewReader(descriptor))

	descriptor, _ = ioutil.ReadAll(reader)
	//d := descriptor


	descriptorpb0 := &descriptorpb.FileDescriptorProto{}




	_ = proto.Unmarshal(descriptor, descriptorpb0)

	//fmt.Println(err)

	fd, err := protodesc.NewFile(descriptorpb0, nil)


	//descriptorpb.FileDescriptorSet{}

	l := list.New()

	m := sync.Map{}

	fmt.Println(err)

	messageDescriptor := fd.Messages().ByName("User")

	msg := dynamicpb.NewMessage(messageDescriptor)

	//message := dynamicpb.NewMessage(User)

	//fmt.Println(message.Descriptor().Name())

	ud, _ := proto.Marshal(user)

	err = proto.Unmarshal(ud, msg)

	fmt.Println(err)

	//protoreflect.

	return descriptorpb
}