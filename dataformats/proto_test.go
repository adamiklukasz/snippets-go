package dataformats

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"snippets-go/dataformats/defs"
	"testing"
)

func TestProtobufSave(t *testing.T) {
	p := defs.Person{
		Name:  "Lukasz",
		Id:    21,
		Email: "lukasz.1234@oa.qw",
		Phones: []*defs.Person_PhoneNumber{
			{
				Number: "123-456-789",
				Type:   defs.Person_HOME,
			},
		},
	}

	out, _ := proto.Marshal(&p)
	_ = ioutil.WriteFile("person.sav", out, 0644)
}

func TestProtobufLoad(t *testing.T) {
	in, _ := ioutil.ReadFile("person.sav")
	p := defs.Person{}
	_ = proto.Unmarshal(in, &p)

	protoToJSON := jsonpb.Marshaler{}
	personJSON, _ := protoToJSON.MarshalToString(&p)

	fmt.Printf("person=[%v]\n", p)
	fmt.Printf("personJSON=%s\n", personJSON)
}
