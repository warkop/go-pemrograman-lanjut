package main

import (
	"bytes"
	"fmt"
	"C29/model"
	"os"
	"strings"

	"github.com/golang/protobuf/jsonpb"
)

func main() {
	var user1 = &model.User{
		Id:       "u001",
		Name:     "John Cena",
		Password: "curangdosa",
		Gender:   model.UserGender_MALE,
	}

	// var userList = &model.userList{
	// 	List: []*model.User{
	// 		user1,
	// 	},
	// }

	var garage1 = &model.Garage{
		Id:   "g001",
		Name: "Kalimdor",
		Coordinate: &model.GarageCoordinate{
			Latitude:  23.2212847,
			Longitude: 53.22033123,
		},
	}

	var garageList = &model.GarageList{
		List: []*model.Garage{
			garage1,
		},
	}

	// var garageListByUser = &model.GarageListByUser{
	// 	List: map[string]*model.GarageList{
	// 		user1.Id: garageList,
	// 	},
	// }

	fmt.Printf("#==== Original \n    %#v \n", user1)
	fmt.Printf("#==== as String\n    %v \n", user1.String())

	var buf bytes.Buffer
	err1 := (&jsonpb.Marshaler{}).Marshal(&buf, garageList)

	if err1 != nil {
		fmt.Println(err1.Error())
		os.Exit(0)
	}

	jsonString := buf.String()
	fmt.Printf("# ==== As JSON String\n       %v \n", jsonString)

	buf2 := strings.NewReader(jsonString)
	protoObject := new(model.GarageList)

	err2 := (&jsonpb.Unmarshaler{}).Unmarshal(buf2, protoObject)
	if err2 != nil {
		fmt.Println(err2.Error())
		os.Exit(0)
	}

	fmt.Printf("# ==== As String\n       %v \n", protoObject.String())
}
