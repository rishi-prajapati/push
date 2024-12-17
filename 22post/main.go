package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int 
	Platform string `json:"website"`
	Password string `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to json video")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJS Boodtcamp", 299,"learncodeonline.in","abc123",[]string{"Web-dev","js"}},
		{"MERN Boodtcamp", 199,"learncodeonline.in","gbc123",[]string{"full-stack","js"}},
		{"Angular Boodtcamp", 399,"learncodeonline.in","avc123",nil},
	}
	//packagae this data as JSON data

	finalJson, err := json.MarshalIndent(lcoCourses, "","\t")
	if err != nil{
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

func DecodeJson(){
	jsonDataFromWeb := []byte(`
        {
                "coursename": "ReactJS Boodtcamp",
                "Price": 299,
                "website": "learncodeonline.in",  
                "tags": ["Web-dev","js"]
        }
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid{
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n",lcoCourse)
	}else{
		fmt.Println("JSON WAS NOT VALID")
	}

	//SOME CASES WHERE YOU JUST WANT TO ADD DATA TO KEY VALUE
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb,&myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k, v := range myOnlineData{
		fmt.Printf("Key is %v and value is %v and Type is %T\n",k,v,v)
	}


}