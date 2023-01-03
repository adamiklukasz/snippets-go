package main

import (
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	os "os"
	defs2 "snippets-go/formats/proto-buffer/defs"
)

func main() {
	save()
	load()
}

func save() {
	p := defs2.Person{
		Name:  "Lukasz",
		Id:    21,
		Email: "lukasz.1234@oa.qw",
		Phones: []*defs2.Person_PhoneNumber{
			{
				Number: "123-456-789",
				Type:   defs2.Person_HOME,
			},
		},
	}

	out, _ := proto.Marshal(&p)
	_ = os.WriteFile("person.sav", out, 0644)
}

func load() {
	in, _ := os.ReadFile("person.sav")
	p := defs2.Person{}
	_ = proto.Unmarshal(in, &p)

	protoToJSON := jsonpb.Marshaler{}
	personJSON, _ := protoToJSON.MarshalToString(&p)

	fmt.Printf("person=[%v]\n", p)
	fmt.Printf("personJSON=%s\n", personJSON)
}
